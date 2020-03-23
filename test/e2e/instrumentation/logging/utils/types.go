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
	"regexp"
	"strconv"

	"sync"
)

var (
	// Regexp, matching the contents of log entries, parsed or not
	logEntryMessageRegex = regexp.MustCompile("(?:I\\d+ \\d+:\\d+:\\d+.\\d+ {7}\\d+ logs_generator.go:67] )?(\\d+) .*")
)

// LogEntry represents a log entry, received from the logging backend.
type LogEntry struct {
	LogName     string
	TextPayload string
	Location    string
	JSONPayload map[string]interface{}
}

// TryGetEntryNumber returns the number of the log entry in sequence, if it
// was generated by the load logging pod (requires special log format).
func (entry LogEntry) TryGetEntryNumber() (int, bool) {
	submatch := logEntryMessageRegex.FindStringSubmatch(entry.TextPayload)
	if submatch == nil || len(submatch) < 2 {
		return 0, false
	}

	lineNumber, err := strconv.Atoi(submatch[1])
	return lineNumber, err == nil
}

// LogsQueueCollection is a thread-safe set of named log queues.
type LogsQueueCollection interface {
	Push(name string, logs ...LogEntry)
	Pop(name string) []LogEntry
}

var _ LogsQueueCollection = &logsQueueCollection{}

type logsQueueCollection struct {
	mutex     *sync.Mutex
	queues    map[string]chan LogEntry
	queueSize int
}

// NewLogsQueueCollection returns a new LogsQueueCollection where each queue
// is created with a default size of queueSize.
func NewLogsQueueCollection(queueSize int) LogsQueueCollection {
	return &logsQueueCollection{
		mutex:     &sync.Mutex{},
		queues:    map[string]chan LogEntry{},
		queueSize: queueSize,
	}
}

func (c *logsQueueCollection) Push(name string, logs ...LogEntry) {
	q := c.getQueue(name)
	for _, log := range logs {
		q <- log
	}
}

func (c *logsQueueCollection) Pop(name string) []LogEntry {
	q := c.getQueue(name)
	var entries []LogEntry
polling_loop:
	for {
		select {
		case entry := <-q:
			entries = append(entries, entry)
		default:
			break polling_loop
		}
	}
	return entries
}

func (c *logsQueueCollection) getQueue(name string) chan LogEntry {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if q, ok := c.queues[name]; ok {
		return q
	}

	newQ := make(chan LogEntry, c.queueSize)
	c.queues[name] = newQ
	return newQ
}
