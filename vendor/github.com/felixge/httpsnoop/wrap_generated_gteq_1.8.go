// +build go1.8
// Code generated by "httpsnoop/codegen"; DO NOT EDIT

package httpsnoop

import (
	"bufio"
	"io"
	"net"
	"net/http"
)

// HeaderFunc is part of the http.ResponseWriter interface.
type HeaderFunc func() http.Header

// WriteHeaderFunc is part of the http.ResponseWriter interface.
type WriteHeaderFunc func(code int)

// WriteFunc is part of the http.ResponseWriter interface.
type WriteFunc func(b []byte) (int, error)

// FlushFunc is part of the http.Flusher interface.
type FlushFunc func()

// CloseNotifyFunc is part of the http.CloseNotifier interface.
type CloseNotifyFunc func() <-chan bool

// HijackFunc is part of the http.Hijacker interface.
type HijackFunc func() (net.Conn, *bufio.ReadWriter, error)

// ReadFromFunc is part of the io.ReaderFrom interface.
type ReadFromFunc func(src io.Reader) (int64, error)

// PushFunc is part of the http.Pusher interface.
type PushFunc func(target string, opts *http.PushOptions) error

// Hooks defines a set of method interceptors for methods included in
// http.ResponseWriter as well as some others. You can think of them as
// middleware for the function calls they target. See Wrap for more details.
type Hooks struct {
	Header      func(HeaderFunc) HeaderFunc
	WriteHeader func(WriteHeaderFunc) WriteHeaderFunc
	Write       func(WriteFunc) WriteFunc
	Flush       func(FlushFunc) FlushFunc
	CloseNotify func(CloseNotifyFunc) CloseNotifyFunc
	Hijack      func(HijackFunc) HijackFunc
	ReadFrom    func(ReadFromFunc) ReadFromFunc
	Push        func(PushFunc) PushFunc
}

// Wrap returns a wrapped version of w that provides the exact same interface
// as w. Specifically if w implements any combination of:
//
// - http.Flusher
// - http.CloseNotifier
// - http.Hijacker
// - io.ReaderFrom
// - http.Pusher
//
// The wrapped version will implement the exact same combination. If no hooks
// are set, the wrapped version also behaves exactly as w. Hooks targeting
// methods not supported by w are ignored. Any other hooks will intercept the
// method they target and may modify the call's arguments and/or return values.
// The CaptureMetrics implementation serves as a working example for how the
// hooks can be used.
func Wrap(w http.ResponseWriter, hooks Hooks) http.ResponseWriter {
	rw := &rw{w: w, h: hooks}
	_, i0 := w.(http.Flusher)
	_, i1 := w.(http.CloseNotifier)
	_, i2 := w.(http.Hijacker)
	_, i3 := w.(io.ReaderFrom)
	_, i4 := w.(http.Pusher)
	switch {
	// combination 1/32
	case !i0 && !i1 && !i2 && !i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
		}{rw, rw}
	// combination 2/32
	case !i0 && !i1 && !i2 && !i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Pusher
		}{rw, rw, rw}
	// combination 3/32
	case !i0 && !i1 && !i2 && i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
		}{rw, rw, rw}
	// combination 4/32
	case !i0 && !i1 && !i2 && i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw}
	// combination 5/32
	case !i0 && !i1 && i2 && !i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
		}{rw, rw, rw}
	// combination 6/32
	case !i0 && !i1 && i2 && !i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			http.Pusher
		}{rw, rw, rw, rw}
	// combination 7/32
	case !i0 && !i1 && i2 && i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw}
	// combination 8/32
	case !i0 && !i1 && i2 && i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 9/32
	case !i0 && i1 && !i2 && !i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
		}{rw, rw, rw}
	// combination 10/32
	case !i0 && i1 && !i2 && !i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
		}{rw, rw, rw, rw}
	// combination 11/32
	case !i0 && i1 && !i2 && i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
		}{rw, rw, rw, rw}
	// combination 12/32
	case !i0 && i1 && !i2 && i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 13/32
	case !i0 && i1 && i2 && !i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
		}{rw, rw, rw, rw}
	// combination 14/32
	case !i0 && i1 && i2 && !i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 15/32
	case !i0 && i1 && i2 && i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw, rw}
	// combination 16/32
	case !i0 && i1 && i2 && i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 17/32
	case i0 && !i1 && !i2 && !i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
		}{rw, rw, rw}
	// combination 18/32
	case i0 && !i1 && !i2 && !i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Pusher
		}{rw, rw, rw, rw}
	// combination 19/32
	case i0 && !i1 && !i2 && i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
		}{rw, rw, rw, rw}
	// combination 20/32
	case i0 && !i1 && !i2 && i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 21/32
	case i0 && !i1 && i2 && !i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
		}{rw, rw, rw, rw}
	// combination 22/32
	case i0 && !i1 && i2 && !i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 23/32
	case i0 && !i1 && i2 && i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw, rw}
	// combination 24/32
	case i0 && !i1 && i2 && i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 25/32
	case i0 && i1 && !i2 && !i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
		}{rw, rw, rw, rw}
	// combination 26/32
	case i0 && i1 && !i2 && !i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Pusher
		}{rw, rw, rw, rw, rw}
	// combination 27/32
	case i0 && i1 && !i2 && i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
		}{rw, rw, rw, rw, rw}
	// combination 28/32
	case i0 && i1 && !i2 && i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 29/32
	case i0 && i1 && i2 && !i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
		}{rw, rw, rw, rw, rw}
	// combination 30/32
	case i0 && i1 && i2 && !i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			http.Pusher
		}{rw, rw, rw, rw, rw, rw}
	// combination 31/32
	case i0 && i1 && i2 && i3 && !i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{rw, rw, rw, rw, rw, rw}
	// combination 32/32
	case i0 && i1 && i2 && i3 && i4:
		return struct {
			Unwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{rw, rw, rw, rw, rw, rw, rw}
	}
	panic("unreachable")
}

type rw struct {
	w http.ResponseWriter
	h Hooks
}

func (w *rw) Unwrap() http.ResponseWriter {
	return w.w
}

func (w *rw) Header() http.Header {
	f := w.w.(http.ResponseWriter).Header
	if w.h.Header != nil {
		f = w.h.Header(f)
	}
	return f()
}

func (w *rw) WriteHeader(code int) {
	f := w.w.(http.ResponseWriter).WriteHeader
	if w.h.WriteHeader != nil {
		f = w.h.WriteHeader(f)
	}
	f(code)
}

func (w *rw) Write(b []byte) (int, error) {
	f := w.w.(http.ResponseWriter).Write
	if w.h.Write != nil {
		f = w.h.Write(f)
	}
	return f(b)
}

func (w *rw) Flush() {
	f := w.w.(http.Flusher).Flush
	if w.h.Flush != nil {
		f = w.h.Flush(f)
	}
	f()
}

func (w *rw) CloseNotify() <-chan bool {
	f := w.w.(http.CloseNotifier).CloseNotify
	if w.h.CloseNotify != nil {
		f = w.h.CloseNotify(f)
	}
	return f()
}

func (w *rw) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	f := w.w.(http.Hijacker).Hijack
	if w.h.Hijack != nil {
		f = w.h.Hijack(f)
	}
	return f()
}

func (w *rw) ReadFrom(src io.Reader) (int64, error) {
	f := w.w.(io.ReaderFrom).ReadFrom
	if w.h.ReadFrom != nil {
		f = w.h.ReadFrom(f)
	}
	return f(src)
}

func (w *rw) Push(target string, opts *http.PushOptions) error {
	f := w.w.(http.Pusher).Push
	if w.h.Push != nil {
		f = w.h.Push(f)
	}
	return f(target, opts)
}

type Unwrapper interface {
	Unwrap() http.ResponseWriter
}

// Unwrap returns the underlying http.ResponseWriter from within zero or more
// layers of httpsnoop wrappers.
func Unwrap(w http.ResponseWriter) http.ResponseWriter {
	if rw, ok := w.(Unwrapper); ok {
		// recurse until rw.Unwrap() returns a non-Unwrapper
		return Unwrap(rw.Unwrap())
	} else {
		return w
	}
}
