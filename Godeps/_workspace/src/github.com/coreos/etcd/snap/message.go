// Copyright 2015 CoreOS, Inc.
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

package snap

import (
	"io"

	"github.com/coreos/etcd/raft/raftpb"
)

// Message is a struct that contains a raft Message and a ReadCloser. The type
// of raft message MUST be MsgSnap, which contains the raft meta-data and an
// additional data []byte field that contains the snapshot of the actual state
// machine.
// Message contains the ReadCloser field for handling large snapshot. This avoid
// copying the entire snapshot into a byte array, which consumes a lot of memory.
//
// User of Message should close the Message after sending it.
type Message struct {
	raftpb.Message
	ReadCloser io.ReadCloser
	closeC     chan bool
}

func NewMessage(rs raftpb.Message, rc io.ReadCloser) *Message {
	return &Message{
		Message:    rs,
		ReadCloser: rc,
		closeC:     make(chan bool, 1),
	}
}

// CloseNotify returns a channel that receives a single value
// when the message sent is finished. true indicates the sent
// is successful.
func (m Message) CloseNotify() <-chan bool {
	return m.closeC
}

func (m Message) CloseWithError(err error) {
	m.ReadCloser.Close()
	if err == nil {
		m.closeC <- true
	} else {
		m.closeC <- false
	}
}
