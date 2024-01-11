//go:build go1.21
// +build go1.21

/*
Copyright 2021 The Kubernetes Authors.

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

package klog

import (
	"log/slog"

	"github.com/go-logr/logr"
)

// SetSlogLogger reconfigures klog to log through the slog logger. The logger must not be nil.
func SetSlogLogger(logger *slog.Logger) {
	SetLoggerWithOptions(logr.FromSlogHandler(logger.Handler()), ContextualLogger(true))
}
