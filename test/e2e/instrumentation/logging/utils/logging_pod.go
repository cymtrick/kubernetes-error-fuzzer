/*
Copyright 2017 The Kubernetes Authors.

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

package utils

import (
	"strconv"
	"time"

	"fmt"

	api_v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/test/e2e/framework"
	e2elog "k8s.io/kubernetes/test/e2e/framework/log"
	imageutils "k8s.io/kubernetes/test/utils/image"
)

const (
	// Amount of requested cores for logging container in millicores
	loggingContainerCPURequest = 10

	// Amount of requested memory for logging container in bytes
	loggingContainerMemoryRequest = 10 * 1024 * 1024

	// Name of the container used for logging tests
	loggingContainerName = "logging-container"
)

// LoggingPod is an interface of a pod that can be started and that logs
// something to its stdout, possibly indefinitely.
type LoggingPod interface {
	// Name equals to the Kubernetes pod name.
	Name() string

	// Start method controls when the logging pod is started in the cluster.
	Start(f *framework.Framework) error
}

// StartAndReturnSelf is a helper method to start a logging pod and
// immediately return it.
func StartAndReturnSelf(p LoggingPod, f *framework.Framework) (LoggingPod, error) {
	err := p.Start(f)
	return p, err
}

// FiniteLoggingPod is a logging pod that emits a known number of log lines.
type FiniteLoggingPod interface {
	LoggingPod

	// ExpectedLinesNumber returns the number of lines that are
	// expected to be ingested from this pod.
	ExpectedLineCount() int
}

var _ FiniteLoggingPod = &loadLoggingPod{}

type loadLoggingPod struct {
	name               string
	nodeName           string
	expectedLinesCount int
	runDuration        time.Duration
}

// NewLoadLoggingPod returns a logging pod that generates totalLines random
// lines over period of length loggingDuration. Lines generated by this
// pod are numbered and have well-defined structure.
func NewLoadLoggingPod(podName string, nodeName string, totalLines int,
	loggingDuration time.Duration) FiniteLoggingPod {
	return &loadLoggingPod{
		name:               podName,
		nodeName:           nodeName,
		expectedLinesCount: totalLines,
		runDuration:        loggingDuration,
	}
}

func (p *loadLoggingPod) Name() string {
	return p.name
}

func (p *loadLoggingPod) Start(f *framework.Framework) error {
	e2elog.Logf("Starting load logging pod %s", p.name)
	f.PodClient().Create(&api_v1.Pod{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: p.name,
		},
		Spec: api_v1.PodSpec{
			RestartPolicy: api_v1.RestartPolicyNever,
			Containers: []api_v1.Container{
				{
					Name:  loggingContainerName,
					Image: imageutils.GetE2EImage(imageutils.LogsGenerator),
					Env: []api_v1.EnvVar{
						{
							Name:  "LOGS_GENERATOR_LINES_TOTAL",
							Value: strconv.Itoa(p.expectedLinesCount),
						},
						{
							Name:  "LOGS_GENERATOR_DURATION",
							Value: p.runDuration.String(),
						},
					},
					Resources: api_v1.ResourceRequirements{
						Requests: api_v1.ResourceList{
							api_v1.ResourceCPU: *resource.NewMilliQuantity(
								loggingContainerCPURequest,
								resource.DecimalSI),
							api_v1.ResourceMemory: *resource.NewQuantity(
								loggingContainerMemoryRequest,
								resource.BinarySI),
						},
					},
				},
			},
			NodeName: p.nodeName,
		},
	})
	return framework.WaitForPodNameRunningInNamespace(f.ClientSet, p.name, f.Namespace.Name)
}

func (p *loadLoggingPod) ExpectedLineCount() int {
	return p.expectedLinesCount
}

// NewRepeatingLoggingPod returns a logging pod that each second prints
// line value to its stdout.
func NewRepeatingLoggingPod(podName string, line string) LoggingPod {
	cmd := []string{
		"/bin/sh",
		"-c",
		fmt.Sprintf("while :; do echo '%s'; sleep 1; done", line),
	}
	return NewExecLoggingPod(podName, cmd)
}

var _ LoggingPod = &execLoggingPod{}

type execLoggingPod struct {
	name string
	cmd  []string
}

// NewExecLoggingPod returns a logging pod that produces logs through
// executing a command, passed in cmd.
func NewExecLoggingPod(podName string, cmd []string) LoggingPod {
	return &execLoggingPod{
		name: podName,
		cmd:  cmd,
	}
}

func (p *execLoggingPod) Name() string {
	return p.name
}

func (p *execLoggingPod) Start(f *framework.Framework) error {
	e2elog.Logf("Starting repeating logging pod %s", p.name)
	f.PodClient().Create(&api_v1.Pod{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: p.name,
		},
		Spec: api_v1.PodSpec{
			Containers: []api_v1.Container{
				{
					Name:    loggingContainerName,
					Image:   imageutils.GetE2EImage(imageutils.BusyBox),
					Command: p.cmd,
					Resources: api_v1.ResourceRequirements{
						Requests: api_v1.ResourceList{
							api_v1.ResourceCPU: *resource.NewMilliQuantity(
								loggingContainerCPURequest,
								resource.DecimalSI),
							api_v1.ResourceMemory: *resource.NewQuantity(
								loggingContainerMemoryRequest,
								resource.BinarySI),
						},
					},
				},
			},
		},
	})
	return framework.WaitForPodNameRunningInNamespace(f.ClientSet, p.name, f.Namespace.Name)
}
