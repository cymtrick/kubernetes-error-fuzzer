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

package remotecommand

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/httpstream"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/kubernetes/pkg/api"
)

type fakeReader struct {
	err error
}

func (r *fakeReader) Read([]byte) (int, error) { return 0, r.err }

type fakeWriter struct{}

func (*fakeWriter) Write([]byte) (int, error) { return 0, nil }

type fakeStreamCreator struct {
	created map[string]bool
	errors  map[string]error
}

var _ streamCreator = &fakeStreamCreator{}

func (f *fakeStreamCreator) CreateStream(headers http.Header) (httpstream.Stream, error) {
	streamType := headers.Get(api.StreamType)
	f.created[streamType] = true
	return nil, f.errors[streamType]
}

func TestV2CreateStreams(t *testing.T) {
	tests := []struct {
		name        string
		stdin       bool
		stdinError  error
		stdout      bool
		stdoutError error
		stderr      bool
		stderrError error
		errorError  error
		tty         bool
		expectError bool
	}{
		{
			name:        "stdin error",
			stdin:       true,
			stdinError:  errors.New("stdin error"),
			expectError: true,
		},
		{
			name:        "stdout error",
			stdout:      true,
			stdoutError: errors.New("stdout error"),
			expectError: true,
		},
		{
			name:        "stderr error",
			stderr:      true,
			stderrError: errors.New("stderr error"),
			expectError: true,
		},
		{
			name:        "error stream error",
			stdin:       true,
			stdout:      true,
			stderr:      true,
			errorError:  errors.New("error stream error"),
			expectError: true,
		},
		{
			name:        "no errors",
			stdin:       true,
			stdout:      true,
			stderr:      true,
			expectError: false,
		},
		{
			name:        "no errors, stderr & tty set, don't expect stderr",
			stdin:       true,
			stdout:      true,
			stderr:      true,
			tty:         true,
			expectError: false,
		},
	}
	for _, test := range tests {
		conn := &fakeStreamCreator{
			created: make(map[string]bool),
			errors: map[string]error{
				api.StreamTypeStdin:  test.stdinError,
				api.StreamTypeStdout: test.stdoutError,
				api.StreamTypeStderr: test.stderrError,
				api.StreamTypeError:  test.errorError,
			},
		}

		opts := StreamOptions{Tty: test.tty}
		if test.stdin {
			opts.Stdin = &fakeReader{}
		}
		if test.stdout {
			opts.Stdout = &fakeWriter{}
		}
		if test.stderr {
			opts.Stderr = &fakeWriter{}
		}

		h := newStreamProtocolV2(opts).(*streamProtocolV2)
		err := h.createStreams(conn)

		if test.expectError {
			if err == nil {
				t.Errorf("%s: expected error", test.name)
				continue
			}
			if e, a := test.stdinError, err; test.stdinError != nil && e != a {
				t.Errorf("%s: expected %v, got %v", test.name, e, a)
			}
			if e, a := test.stdoutError, err; test.stdoutError != nil && e != a {
				t.Errorf("%s: expected %v, got %v", test.name, e, a)
			}
			if e, a := test.stderrError, err; test.stderrError != nil && e != a {
				t.Errorf("%s: expected %v, got %v", test.name, e, a)
			}
			if e, a := test.errorError, err; test.errorError != nil && e != a {
				t.Errorf("%s: expected %v, got %v", test.name, e, a)
			}
			continue
		}

		if !test.expectError && err != nil {
			t.Errorf("%s: unexpected error: %v", test.name, err)
			continue
		}

		if test.stdin && !conn.created[api.StreamTypeStdin] {
			t.Errorf("%s: expected stdin stream", test.name)
		}
		if test.stdout && !conn.created[api.StreamTypeStdout] {
			t.Errorf("%s: expected stdout stream", test.name)
		}
		if test.stderr {
			if test.tty && conn.created[api.StreamTypeStderr] {
				t.Errorf("%s: unexpected stderr stream because tty is set", test.name)
			} else if !test.tty && !conn.created[api.StreamTypeStderr] {
				t.Errorf("%s: expected stderr stream", test.name)
			}
		}
		if !conn.created[api.StreamTypeError] {
			t.Errorf("%s: expected error stream", test.name)
		}

	}
}

func TestV2ErrorStreamReading(t *testing.T) {
	tests := []struct {
		name          string
		stream        io.Reader
		expectedError error
	}{
		{
			name:          "error reading from stream",
			stream:        &fakeReader{errors.New("foo")},
			expectedError: errors.New("error reading from error stream: foo"),
		},
		{
			name:          "stream returns an error",
			stream:        strings.NewReader("some error"),
			expectedError: errors.New("error executing remote command: some error"),
		},
	}

	for _, test := range tests {
		h := newStreamProtocolV2(StreamOptions{}).(*streamProtocolV2)
		h.errorStream = test.stream

		ch := watchErrorStream(h.errorStream, &errorDecoderV2{})
		if ch == nil {
			t.Fatalf("%s: unexpected nil channel", test.name)
		}

		var err error
		select {
		case err = <-ch:
		case <-time.After(wait.ForeverTestTimeout):
			t.Fatalf("%s: timed out", test.name)
		}

		if test.expectedError != nil {
			if err == nil {
				t.Errorf("%s: expected an error", test.name)
			} else if e, a := test.expectedError, err; e.Error() != a.Error() {
				t.Errorf("%s: expected %q, got %q", test.name, e, a)
			}
			continue
		}

		if test.expectedError == nil && err != nil {
			t.Errorf("%s: unexpected error: %v", test.name, err)
			continue
		}
	}
}
