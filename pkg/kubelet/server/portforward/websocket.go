/*
Copyright 2016 The Kubernetes Authors.

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

package portforward

import (
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apiserver/pkg/server/httplog"
	"k8s.io/apiserver/pkg/util/wsstream"
	"k8s.io/kubernetes/pkg/api"
)

const (
	dataChannel = iota
	errorChannel

	v4BinaryWebsocketProtocol = "v4." + wsstream.ChannelWebSocketProtocol
	v4Base64WebsocketProtocol = "v4." + wsstream.Base64ChannelWebSocketProtocol
)

// options contains details about which streams are required for
// port forwarding.
type v4Options struct {
	ports []uint16
}

// newOptions creates a new options from the Request.
func newV4Options(req *http.Request) (*v4Options, error) {
	portStrings := req.URL.Query()[api.PortHeader]
	if len(portStrings) == 0 {
		return nil, fmt.Errorf("%q is required", api.PortHeader)
	}

	ports := make([]uint16, 0, len(portStrings))
	for _, portString := range portStrings {
		if len(portString) == 0 {
			return nil, fmt.Errorf("%q is cannot be empty", api.PortHeader)
		}
		for _, p := range strings.Split(portString, ",") {
			port, err := strconv.ParseUint(p, 10, 16)
			if err != nil {
				return nil, fmt.Errorf("unable to parse %q as a port: %v", portString, err)
			}
			if port < 1 {
				return nil, fmt.Errorf("port %q must be > 0", portString)
			}
			ports = append(ports, uint16(port))
		}
	}

	return &v4Options{
		ports: ports,
	}, nil
}

func handleWebSocketStreams(req *http.Request, w http.ResponseWriter, portForwarder PortForwarder, podName string, uid types.UID, supportedPortForwardProtocols []string, idleTimeout, streamCreationTimeout time.Duration) error {
	opts, err := newV4Options(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return err
	}

	channels := make([]wsstream.ChannelType, 0, len(opts.ports)*2)
	for i := 0; i < len(opts.ports); i++ {
		channels = append(channels, wsstream.ReadWriteChannel, wsstream.WriteChannel)
	}
	conn := wsstream.NewConn(map[string]wsstream.ChannelProtocolConfig{
		"": {
			Binary:   true,
			Channels: channels,
		},
		v4BinaryWebsocketProtocol: {
			Binary:   true,
			Channels: channels,
		},
		v4Base64WebsocketProtocol: {
			Binary:   false,
			Channels: channels,
		},
	})
	conn.SetIdleTimeout(idleTimeout)
	_, streams, err := conn.Open(httplog.Unlogged(w), req)
	if err != nil {
		err = fmt.Errorf("Unable to upgrade websocket connection: %v", err)
		return err
	}
	defer conn.Close()
	streamPairs := make([]*websocketStreamPair, len(opts.ports))
	for i := range streamPairs {
		streamPair := websocketStreamPair{
			port:        opts.ports[i],
			dataStream:  streams[i*2+dataChannel],
			errorStream: streams[i*2+errorChannel],
		}
		streamPairs[i] = &streamPair

		portBytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(portBytes, streamPair.port)
		streamPair.dataStream.Write(portBytes)
		streamPair.errorStream.Write(portBytes)
	}
	h := &websocketStreamHandler{
		conn:        conn,
		streamPairs: streamPairs,
		pod:         podName,
		uid:         uid,
		forwarder:   portForwarder,
	}
	h.run()

	return nil
}

// websocketStreamPair represents the error and data streams for a port
// forwarding request.
type websocketStreamPair struct {
	port        uint16
	dataStream  io.ReadWriteCloser
	errorStream io.WriteCloser
}

// websocketStreamHandler is capable of processing a single port forward
// request over a websocket connection
type websocketStreamHandler struct {
	conn        *wsstream.Conn
	ports       []uint16
	streamPairs []*websocketStreamPair
	pod         string
	uid         types.UID
	forwarder   PortForwarder
}

// run invokes the websocketStreamHandler's forwarder.PortForward
// function for the given stream pair.
func (h *websocketStreamHandler) run() {
	wg := sync.WaitGroup{}
	wg.Add(len(h.streamPairs))

	for _, pair := range h.streamPairs {
		p := pair
		go func() {
			defer wg.Done()
			h.portForward(p)
		}()
	}

	wg.Wait()
}

func (h *websocketStreamHandler) portForward(p *websocketStreamPair) {
	defer p.dataStream.Close()
	defer p.errorStream.Close()

	glog.V(5).Infof("(conn=%p) invoking forwarder.PortForward for port %d", h.conn, p.port)
	err := h.forwarder.PortForward(h.pod, h.uid, p.port, p.dataStream)
	glog.V(5).Infof("(conn=%p) done invoking forwarder.PortForward for port %d", h.conn, p.port)

	if err != nil {
		msg := fmt.Errorf("error forwarding port %d to pod %s, uid %v: %v", p.port, h.pod, h.uid, err)
		runtime.HandleError(msg)
		fmt.Fprint(p.errorStream, msg.Error())
	}
}
