/*
Copyright 2014 The Kubernetes Authors.

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

package exec

import (
	"fmt"
	"io"
	"net/url"
	"time"

	dockerterm "github.com/docker/docker/pkg/term"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"

	"k8s.io/kubectl/pkg/scheme"
	"k8s.io/kubectl/pkg/util/interrupt"
	"k8s.io/kubectl/pkg/util/templates"
	"k8s.io/kubectl/pkg/util/term"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/polymorphichelpers"
	"k8s.io/kubernetes/pkg/kubectl/util/i18n"
)

var (
	execExample = templates.Examples(i18n.T(`
		# Get output from running 'date' command from pod mypod, using the first container by default
		kubectl exec mypod date

		# Get output from running 'date' command in ruby-container from pod mypod
		kubectl exec mypod -c ruby-container date

		# Switch to raw terminal mode, sends stdin to 'bash' in ruby-container from pod mypod
		# and sends stdout/stderr from 'bash' back to the client
		kubectl exec mypod -c ruby-container -i -t -- bash -il

		# List contents of /usr from the first container of pod mypod and sort by modification time.
		# If the command you want to execute in the pod has any flags in common (e.g. -i),
		# you must use two dashes (--) to separate your command's flags/arguments.
		# Also note, do not surround your command and its flags/arguments with quotes
		# unless that is how you would execute it normally (i.e., do ls -t /usr, not "ls -t /usr").
		kubectl exec mypod -i -t -- ls -t /usr

		# Get output from running 'date' command from the first pod of the deployment mydeployment, using the first container by default
		kubectl exec deploy/mydeployment date

		# Get output from running 'date' command from the first pod of the service myservice, using the first container by default
		kubectl exec svc/myservice date
		`))
)

const (
	execUsageStr          = "expected 'exec (POD | TYPE/NAME) COMMAND [ARG1] [ARG2] ... [ARGN]'.\nPOD or TYPE/NAME and COMMAND are required arguments for the exec command"
	defaultPodExecTimeout = 60 * time.Second
)

func NewCmdExec(f cmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	options := &ExecOptions{
		StreamOptions: StreamOptions{
			IOStreams: streams,
		},

		Executor: &DefaultRemoteExecutor{},
	}
	cmd := &cobra.Command{
		Use:                   "exec (POD | TYPE/NAME) [-c CONTAINER] [flags] -- COMMAND [args...]",
		DisableFlagsInUseLine: true,
		Short:                 i18n.T("Execute a command in a container"),
		Long:                  "Execute a command in a container.",
		Example:               execExample,
		Run: func(cmd *cobra.Command, args []string) {
			argsLenAtDash := cmd.ArgsLenAtDash()
			cmdutil.CheckErr(options.Complete(f, cmd, args, argsLenAtDash))
			cmdutil.CheckErr(options.Validate())
			cmdutil.CheckErr(options.Run())
		},
	}
	cmdutil.AddPodRunningTimeoutFlag(cmd, defaultPodExecTimeout)
	// TODO support UID
	cmd.Flags().StringVarP(&options.ContainerName, "container", "c", options.ContainerName, "Container name. If omitted, the first container in the pod will be chosen")
	cmd.Flags().BoolVarP(&options.Stdin, "stdin", "i", options.Stdin, "Pass stdin to the container")
	cmd.Flags().BoolVarP(&options.TTY, "tty", "t", options.TTY, "Stdin is a TTY")
	return cmd
}

// RemoteExecutor defines the interface accepted by the Exec command - provided for test stubbing
type RemoteExecutor interface {
	Execute(method string, url *url.URL, config *restclient.Config, stdin io.Reader, stdout, stderr io.Writer, tty bool, terminalSizeQueue remotecommand.TerminalSizeQueue) error
}

// DefaultRemoteExecutor is the standard implementation of remote command execution
type DefaultRemoteExecutor struct{}

func (*DefaultRemoteExecutor) Execute(method string, url *url.URL, config *restclient.Config, stdin io.Reader, stdout, stderr io.Writer, tty bool, terminalSizeQueue remotecommand.TerminalSizeQueue) error {
	exec, err := remotecommand.NewSPDYExecutor(config, method, url)
	if err != nil {
		return err
	}
	return exec.Stream(remotecommand.StreamOptions{
		Stdin:             stdin,
		Stdout:            stdout,
		Stderr:            stderr,
		Tty:               tty,
		TerminalSizeQueue: terminalSizeQueue,
	})
}

type StreamOptions struct {
	Namespace     string
	PodName       string
	ContainerName string
	Stdin         bool
	TTY           bool
	// minimize unnecessary output
	Quiet bool
	// InterruptParent, if set, is used to handle interrupts while attached
	InterruptParent *interrupt.Handler

	genericclioptions.IOStreams

	// for testing
	overrideStreams func() (io.ReadCloser, io.Writer, io.Writer)
	isTerminalIn    func(t term.TTY) bool
}

// ExecOptions declare the arguments accepted by the Exec command
type ExecOptions struct {
	StreamOptions

	ResourceName string
	Command      []string

	ParentCommandName       string
	EnableSuggestedCmdUsage bool

	Builder          func() *resource.Builder
	ExecutablePodFn  polymorphichelpers.AttachablePodForObjectFunc
	restClientGetter genericclioptions.RESTClientGetter

	Pod           *corev1.Pod
	Executor      RemoteExecutor
	PodClient     coreclient.PodsGetter
	GetPodTimeout time.Duration
	Config        *restclient.Config
}

// Complete verifies command line arguments and loads data from the command environment
func (p *ExecOptions) Complete(f cmdutil.Factory, cmd *cobra.Command, argsIn []string, argsLenAtDash int) error {
	// Let kubectl exec follow rules for `--`, see #13004 issue
	if len(argsIn) == 0 || argsLenAtDash == 0 {
		return cmdutil.UsageErrorf(cmd, execUsageStr)
	}

	p.ResourceName = argsIn[0]
	p.Command = argsIn[1:]

	var err error

	p.Namespace, _, err = f.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}

	p.ExecutablePodFn = polymorphichelpers.AttachablePodForObjectFn

	p.GetPodTimeout, err = cmdutil.GetPodRunningTimeoutFlag(cmd)
	if err != nil {
		return cmdutil.UsageErrorf(cmd, err.Error())
	}

	p.Builder = f.NewBuilder
	p.restClientGetter = f

	cmdParent := cmd.Parent()
	if cmdParent != nil {
		p.ParentCommandName = cmdParent.CommandPath()
	}
	if len(p.ParentCommandName) > 0 && cmdutil.IsSiblingCommandExists(cmd, "describe") {
		p.EnableSuggestedCmdUsage = true
	}

	p.Config, err = f.ToRESTConfig()
	if err != nil {
		return err
	}

	clientset, err := f.KubernetesClientSet()
	if err != nil {
		return err
	}
	p.PodClient = clientset.CoreV1()

	return nil
}

// Validate checks that the provided exec options are specified.
func (p *ExecOptions) Validate() error {
	if len(p.PodName) == 0 && len(p.ResourceName) == 0 {
		return fmt.Errorf("pod or type/name must be specified")
	}
	if len(p.Command) == 0 {
		return fmt.Errorf("you must specify at least one command for the container")
	}
	if p.Out == nil || p.ErrOut == nil {
		return fmt.Errorf("both output and error output must be provided")
	}
	return nil
}

func (o *StreamOptions) SetupTTY() term.TTY {
	t := term.TTY{
		Parent: o.InterruptParent,
		Out:    o.Out,
	}

	if !o.Stdin {
		// need to nil out o.In to make sure we don't create a stream for stdin
		o.In = nil
		o.TTY = false
		return t
	}

	t.In = o.In
	if !o.TTY {
		return t
	}

	if o.isTerminalIn == nil {
		o.isTerminalIn = func(tty term.TTY) bool {
			return tty.IsTerminalIn()
		}
	}
	if !o.isTerminalIn(t) {
		o.TTY = false

		if o.ErrOut != nil {
			fmt.Fprintln(o.ErrOut, "Unable to use a TTY - input is not a terminal or the right kind of file")
		}

		return t
	}

	// if we get to here, the user wants to attach stdin, wants a TTY, and o.In is a terminal, so we
	// can safely set t.Raw to true
	t.Raw = true

	if o.overrideStreams == nil {
		// use dockerterm.StdStreams() to get the right I/O handles on Windows
		o.overrideStreams = dockerterm.StdStreams
	}
	stdin, stdout, _ := o.overrideStreams()
	o.In = stdin
	t.In = stdin
	if o.Out != nil {
		o.Out = stdout
		t.Out = stdout
	}

	return t
}

// Run executes a validated remote execution against a pod.
func (p *ExecOptions) Run() error {
	var err error
	// we still need legacy pod getter when PodName in ExecOptions struct is provided,
	// since there are any other command run this function by providing Podname with PodsGetter
	// and without resource builder, eg: `kubectl cp`.
	if len(p.PodName) != 0 {
		p.Pod, err = p.PodClient.Pods(p.Namespace).Get(p.PodName, metav1.GetOptions{})
		if err != nil {
			return err
		}
	} else {
		builder := p.Builder().
			WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).
			NamespaceParam(p.Namespace).DefaultNamespace().ResourceNames("pods", p.ResourceName)

		obj, err := builder.Do().Object()
		if err != nil {
			return err
		}

		p.Pod, err = p.ExecutablePodFn(p.restClientGetter, obj, p.GetPodTimeout)
		if err != nil {
			return err
		}
	}

	pod := p.Pod

	if pod.Status.Phase == corev1.PodSucceeded || pod.Status.Phase == corev1.PodFailed {
		return fmt.Errorf("cannot exec into a container in a completed pod; current phase is %s", pod.Status.Phase)
	}

	containerName := p.ContainerName
	if len(containerName) == 0 {
		if len(pod.Spec.Containers) > 1 {
			fmt.Fprintf(p.ErrOut, "Defaulting container name to %s.\n", pod.Spec.Containers[0].Name)
			if p.EnableSuggestedCmdUsage {
				fmt.Fprintf(p.ErrOut, "Use '%s describe pod/%s -n %s' to see all of the containers in this pod.\n", p.ParentCommandName, pod.Name, p.Namespace)
			}
		}
		containerName = pod.Spec.Containers[0].Name
	}

	// ensure we can recover the terminal while attached
	t := p.SetupTTY()

	var sizeQueue remotecommand.TerminalSizeQueue
	if t.Raw {
		// this call spawns a goroutine to monitor/update the terminal size
		sizeQueue = t.MonitorSize(t.GetSize())

		// unset p.Err if it was previously set because both stdout and stderr go over p.Out when tty is
		// true
		p.ErrOut = nil
	}

	fn := func() error {
		restClient, err := restclient.RESTClientFor(p.Config)
		if err != nil {
			return err
		}

		// TODO: consider abstracting into a client invocation or client helper
		req := restClient.Post().
			Resource("pods").
			Name(pod.Name).
			Namespace(pod.Namespace).
			SubResource("exec")
		req.VersionedParams(&corev1.PodExecOptions{
			Container: containerName,
			Command:   p.Command,
			Stdin:     p.Stdin,
			Stdout:    p.Out != nil,
			Stderr:    p.ErrOut != nil,
			TTY:       t.Raw,
		}, scheme.ParameterCodec)

		return p.Executor.Execute("POST", req.URL(), p.Config, p.In, p.Out, p.ErrOut, t.Raw, sizeQueue)
	}

	if err := t.Safe(fn); err != nil {
		return err
	}

	return nil
}
