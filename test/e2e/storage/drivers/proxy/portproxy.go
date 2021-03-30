/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package proxy

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/httpstream"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"k8s.io/klog/v2"
)

// Maximum number of forwarded connections. In practice we don't
// need more than one per sidecar and kubelet. Keeping this reasonably
// small ensures that we don't establish connections through the apiserver
// and the remote kernel which then aren't needed.
const maxConcurrentConnections = 10

// Listen creates a listener which returns new connections whenever someone connects
// to a socat or mock driver proxy instance running inside the given pod.
//
// socat must by started with "<listen>,fork TCP-LISTEN:<port>,reuseport"
// for this to work. "<listen>" can be anything that accepts connections,
// for example "UNIX-LISTEN:/csi/csi.sock". In this mode, socat will
// accept exactly one connection on the given port for each connection
// that socat itself accepted.
//
// Listening stops when the context is done or Close() is called.
func Listen(ctx context.Context, clientset kubernetes.Interface, restConfig *rest.Config, addr Addr) (net.Listener, error) {
	// We connect through port forwarding. Strictly
	// speaking this is overkill because we don't need a local
	// port. But this way we can reuse client-go/tools/portforward
	// instead of having to replicate handleConnection
	// in our own code.
	restClient := clientset.CoreV1().RESTClient()
	if restConfig.GroupVersion == nil {
		restConfig.GroupVersion = &schema.GroupVersion{}
	}
	if restConfig.NegotiatedSerializer == nil {
		restConfig.NegotiatedSerializer = scheme.Codecs
	}

	// The setup code around the actual portforward is from
	// https://github.com/kubernetes/kubernetes/blob/c652ffbe4a29143623a1aaec39f745575f7e43ad/staging/src/k8s.io/kubectl/pkg/cmd/portforward/portforward.go
	req := restClient.Post().
		Resource("pods").
		Namespace(addr.Namespace).
		Name(addr.PodName).
		SubResource("portforward")
	transport, upgrader, err := spdy.RoundTripperFor(restConfig)
	if err != nil {
		return nil, fmt.Errorf("create round tripper: %v", err)
	}
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, "POST", req.URL())

	prefix := fmt.Sprintf("port forwarding for %s", addr)
	ctx, cancel := context.WithCancel(ctx)
	l := &listener{
		connections: make(chan *connection),
		ctx:         ctx,
		cancel:      cancel,
		addr:        addr,
	}

	var connectionsCreated, connectionsClosed int32

	runForwarding := func() {
		klog.V(2).Infof("%s: starting connection polling", prefix)
		defer klog.V(2).Infof("%s: connection polling ended", prefix)

		// This delay determines how quickly we notice when someone has
		// connected inside the cluster. With socat, we cannot make this too small
		// because otherwise we get many rejected connections. With the mock
		// driver as proxy that doesn't happen as long as we don't
		// ask for too many concurrent connections because the mock driver
		// keeps the listening port open at all times and the Linux
		// kernel automatically accepts our connection requests.
		tryConnect := time.NewTicker(100 * time.Millisecond)
		defer tryConnect.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-tryConnect.C:
				currentClosed := atomic.LoadInt32(&connectionsClosed)
				openConnections := connectionsCreated - currentClosed
				if openConnections >= maxConcurrentConnections {
					break
				}

				klog.V(5).Infof("%s: trying to create a new connection #%d, %d open", prefix, connectionsCreated, openConnections)
				stream, err := dial(ctx, fmt.Sprintf("%s #%d", prefix, connectionsCreated), dialer, addr.Port)
				if err != nil {
					klog.Errorf("%s: no connection: %v", prefix, err)
					break
				}
				// Make the connection available to Accept below.
				klog.V(5).Infof("%s: created a new connection #%d", prefix, connectionsCreated)
				l.connections <- &connection{
					stream:  stream,
					addr:    addr,
					counter: connectionsCreated,
					closed:  &connectionsClosed,
				}
				connectionsCreated++
			}
		}
	}

	// Portforwarding and polling for connections run in the background.
	go func() {
		for {
			running := false
			pod, err := clientset.CoreV1().Pods(addr.Namespace).Get(ctx, addr.PodName, metav1.GetOptions{})
			if err != nil {
				klog.V(5).Infof("checking for container %q in pod %s/%s: %v", addr.ContainerName, addr.Namespace, addr.PodName, err)
			}
			for i, status := range pod.Status.ContainerStatuses {
				if pod.Spec.Containers[i].Name == addr.ContainerName &&
					status.State.Running != nil {
					running = true
					break
				}
			}

			if running {
				klog.V(2).Infof("container %q in pod %s/%s is running", addr.ContainerName, addr.Namespace, addr.PodName)
				runForwarding()
			}

			select {
			case <-ctx.Done():
				return
			// Sleep a bit before restarting. This is
			// where we potentially wait for the pod to
			// start.
			case <-time.After(1 * time.Second):
			}
		}
	}()

	return l, nil
}

// Addr contains all relevant parameters for a certain port in a pod.
// The container must be running before connections are attempted.
type Addr struct {
	Namespace, PodName, ContainerName string
	Port                              int
}

var _ net.Addr = Addr{}

func (a Addr) Network() string {
	return "port-forwarding"
}

func (a Addr) String() string {
	return fmt.Sprintf("%s/%s:%d", a.Namespace, a.PodName, a.Port)
}

type stream struct {
	httpstream.Stream
	streamConn httpstream.Connection
}

func dial(ctx context.Context, prefix string, dialer httpstream.Dialer, port int) (s *stream, finalErr error) {
	streamConn, _, err := dialer.Dial(portforward.PortForwardProtocolV1Name)
	if err != nil {
		return nil, fmt.Errorf("dialer failed: %v", err)
	}
	requestID := "1"
	defer func() {
		if finalErr != nil {
			streamConn.Close()
		}
	}()

	// create error stream
	headers := http.Header{}
	headers.Set(v1.StreamType, v1.StreamTypeError)
	headers.Set(v1.PortHeader, fmt.Sprintf("%d", port))
	headers.Set(v1.PortForwardRequestIDHeader, requestID)

	// We're not writing to this stream, just reading an error message from it.
	// This happens asynchronously.
	errorStream, err := streamConn.CreateStream(headers)
	if err != nil {
		return nil, fmt.Errorf("error creating error stream: %v", err)
	}
	errorStream.Close()
	go func() {
		message, err := ioutil.ReadAll(errorStream)
		switch {
		case err != nil:
			klog.Errorf("%s: error reading from error stream: %v", prefix, err)
		case len(message) > 0:
			klog.Errorf("%s: an error occurred connecting to the remote port: %v", prefix, string(message))
		}
	}()

	// create data stream
	headers.Set(v1.StreamType, v1.StreamTypeData)
	dataStream, err := streamConn.CreateStream(headers)
	if err != nil {
		return nil, fmt.Errorf("error creating data stream: %v", err)
	}

	return &stream{
		Stream:     dataStream,
		streamConn: streamConn,
	}, nil
}

func (s *stream) Close() {
	s.Stream.Close()
	s.streamConn.Close()
}

type listener struct {
	addr        Addr
	connections chan *connection
	ctx         context.Context
	cancel      func()
}

var _ net.Listener = &listener{}

func (l *listener) Close() error {
	klog.V(5).Infof("forward listener for %s: closing", l.addr)
	l.cancel()
	return nil
}

func (l *listener) Accept() (net.Conn, error) {
	select {
	case <-l.ctx.Done():
		return nil, errors.New("listening was stopped")
	case c := <-l.connections:
		klog.V(5).Infof("forward listener for %s: got a new connection #%d", l.addr, c.counter)
		return c, nil
	}
}

type connection struct {
	stream  *stream
	addr    Addr
	counter int32
	closed  *int32
	mutex   sync.Mutex
}

var _ net.Conn = &connection{}

func (c *connection) LocalAddr() net.Addr {
	return c.addr
}

func (c *connection) RemoteAddr() net.Addr {
	return c.addr
}

func (c *connection) SetDeadline(t time.Time) error {
	return nil
}

func (c *connection) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *connection) SetWriteDeadline(t time.Time) error {
	return nil
}

func (c *connection) Read(b []byte) (int, error) {
	n, err := c.stream.Read(b)
	if errors.Is(err, io.EOF) {
		klog.V(5).Infof("forward connection #%d for %s: remote side closed the stream", c.counter, c.addr)
	}
	return n, err
}

func (c *connection) Write(b []byte) (int, error) {
	n, err := c.stream.Write(b)
	if errors.Is(err, io.EOF) {
		klog.V(5).Infof("forward connection #%d for %s: remote side closed the stream", c.counter, c.addr)
	}
	return n, err
}

func (c *connection) Close() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.closed != nil {
		// Do the logging and book-keeping only once. The function itself may be called more than once.
		klog.V(5).Infof("forward connection #%d for %s: closing our side", c.counter, c.addr)
		atomic.AddInt32(c.closed, 1)
		c.closed = nil
	}
	c.stream.Close()
	return nil
}

func (l *listener) Addr() net.Addr {
	return l.addr
}
