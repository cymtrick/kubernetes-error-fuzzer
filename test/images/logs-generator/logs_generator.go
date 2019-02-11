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

package main

import (
	"flag"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/klog"
)

var (
	httpMethods = []string{
		"GET",
		"POST",
		"PUT",
	}
	namespaces = []string{
		"kube-system",
		"default",
		"ns",
	}
)

var (
	linesTotal = flag.Int("log-lines-total", 0, "Total lines that should be generated by the end of the run")
	duration   = flag.Duration("run-duration", 0, "Total duration of the run")
)

func main() {
	flag.Parse()

	if *linesTotal <= 0 {
		klog.Fatalf("Invalid total number of lines: %d", *linesTotal)
	}

	if *duration <= 0 {
		klog.Fatalf("Invalid duration: %v", *duration)
	}

	generateLogs(*linesTotal, *duration)
}

// Outputs linesTotal lines of logs to stdout uniformly for duration
func generateLogs(linesTotal int, duration time.Duration) {
	delay := duration / time.Duration(linesTotal)

	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for id := 0; id < linesTotal; id++ {
		klog.Info(generateLogLine(id))
		<-ticker.C
	}
}

// Generates apiserver-like line with average length of 100 symbols
func generateLogLine(id int) string {
	method := httpMethods[rand.Intn(len(httpMethods))]
	namespace := namespaces[rand.Intn(len(namespaces))]

	podName := rand.String(rand.IntnRange(3, 5))
	url := fmt.Sprintf("/api/v1/namespaces/%s/pods/%s", namespace, podName)
	status := rand.IntnRange(200, 600)

	return fmt.Sprintf("%d %s %s %d", id, method, url, status)
}
