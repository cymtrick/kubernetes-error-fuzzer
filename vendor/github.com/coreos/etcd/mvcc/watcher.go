// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mvcc

import (
	"bytes"
	"errors"
	"sync"

	"github.com/coreos/etcd/mvcc/mvccpb"
)

var (
	ErrWatcherNotExist = errors.New("mvcc: watcher does not exist")
)

type WatchID int64

type WatchStream interface {
	// Watch creates a watcher. The watcher watches the events happening or
	// happened on the given key or range [key, end) from the given startRev.
	//
	// The whole event history can be watched unless compacted.
	// If `startRev` <=0, watch observes events after currentRev.
	//
	// The returned `id` is the ID of this watcher. It appears as WatchID
	// in events that are sent to the created watcher through stream channel.
	//
	Watch(key, end []byte, startRev int64) WatchID

	// Chan returns a chan. All watch response will be sent to the returned chan.
	Chan() <-chan WatchResponse

	// RequestProgress requests the progress of the watcher with given ID. The response
	// will only be sent if the watcher is currently synced.
	// The responses will be sent through the WatchRespone Chan attached
	// with this stream to ensure correct ordering.
	// The responses contains no events. The revision in the response is the progress
	// of the watchers since the watcher is currently synced.
	RequestProgress(id WatchID)

	// Cancel cancels a watcher by giving its ID. If watcher does not exist, an error will be
	// returned.
	Cancel(id WatchID) error

	// Close closes Chan and release all related resources.
	Close()

	// Rev returns the current revision of the KV the stream watches on.
	Rev() int64
}

type WatchResponse struct {
	// WatchID is the WatchID of the watcher this response sent to.
	WatchID WatchID

	// Events contains all the events that needs to send.
	Events []mvccpb.Event

	// Revision is the revision of the KV when the watchResponse is created.
	// For a normal response, the revision should be the same as the last
	// modified revision inside Events. For a delayed response to a unsynced
	// watcher, the revision is greater than the last modified revision
	// inside Events.
	Revision int64

	// CompactRevision is set when the watcher is cancelled due to compaction.
	CompactRevision int64
}

// watchStream contains a collection of watchers that share
// one streaming chan to send out watched events and other control events.
type watchStream struct {
	watchable watchable
	ch        chan WatchResponse

	mu sync.Mutex // guards fields below it
	// nextID is the ID pre-allocated for next new watcher in this stream
	nextID   WatchID
	closed   bool
	cancels  map[WatchID]cancelFunc
	watchers map[WatchID]*watcher
}

// Watch creates a new watcher in the stream and returns its WatchID.
// TODO: return error if ws is closed?
func (ws *watchStream) Watch(key, end []byte, startRev int64) WatchID {
	// prevent wrong range where key >= end lexicographically
	// watch request with 'WithFromKey' has empty-byte range end
	if len(end) != 0 && bytes.Compare(key, end) != -1 {
		return -1
	}

	ws.mu.Lock()
	defer ws.mu.Unlock()
	if ws.closed {
		return -1
	}

	id := ws.nextID
	ws.nextID++

	w, c := ws.watchable.watch(key, end, startRev, id, ws.ch)

	ws.cancels[id] = c
	ws.watchers[id] = w
	return id
}

func (ws *watchStream) Chan() <-chan WatchResponse {
	return ws.ch
}

func (ws *watchStream) Cancel(id WatchID) error {
	ws.mu.Lock()
	cancel, ok := ws.cancels[id]
	ok = ok && !ws.closed
	if ok {
		delete(ws.cancels, id)
		delete(ws.watchers, id)
	}
	ws.mu.Unlock()
	if !ok {
		return ErrWatcherNotExist
	}
	cancel()
	return nil
}

func (ws *watchStream) Close() {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	for _, cancel := range ws.cancels {
		cancel()
	}
	ws.closed = true
	close(ws.ch)
	watchStreamGauge.Dec()
}

func (ws *watchStream) Rev() int64 {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	return ws.watchable.rev()
}

func (ws *watchStream) RequestProgress(id WatchID) {
	ws.mu.Lock()
	w, ok := ws.watchers[id]
	ws.mu.Unlock()
	if !ok {
		return
	}
	ws.watchable.progress(w)
}
