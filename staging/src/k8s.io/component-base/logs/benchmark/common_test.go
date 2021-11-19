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

package benchmark

import (
	"flag"
	"io"

	"github.com/go-logr/logr"
	"go.uber.org/zap/zapcore"

	logsjson "k8s.io/component-base/logs/json"
	"k8s.io/klog/v2"
)

func init() {
	// Cause all klog output to be discarded with minimal overhead.
	// We don't include time stamps and caller information.
	klog.InitFlags(nil)
	flag.Set("alsologtostderr", "false")
	flag.Set("logtostderr", "false")
	flag.Set("skip_headers", "true")
	flag.Set("one_output", "true")
	flag.Set("stderrthreshold", "FATAL")
	klog.SetOutput(&output)
}

type bytesWritten int64

func (b *bytesWritten) Write(data []byte) (int, error) {
	l := len(data)
	*b += bytesWritten(l)
	return l, nil
}

func (b *bytesWritten) Sync() error {
	return nil
}

var output bytesWritten
var jsonLogger = newJSONLogger(&output)

func newJSONLogger(out io.Writer) logr.Logger {
	encoderConfig := &zapcore.EncoderConfig{
		MessageKey: "msg",
	}
	logger, _ := logsjson.NewJSONLogger(zapcore.AddSync(out), nil, encoderConfig)
	return logger
}

func printf(item logMessage) {
	if item.isError {
		klog.Errorf("%s: %v %s", item.msg, item.err, item.kvs)
	} else {
		klog.Infof("%s: %v", item.msg, item.kvs)
	}
}

func prints(item logMessage) {
	if item.isError {
		klog.ErrorS(item.err, item.msg, item.kvs...)
	} else {
		klog.InfoS(item.msg, item.kvs...)
	}
}
