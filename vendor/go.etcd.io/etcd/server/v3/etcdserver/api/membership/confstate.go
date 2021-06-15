// Copyright 2021 The etcd Authors
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

package membership

import (
	"encoding/json"
	"log"

	"go.etcd.io/etcd/raft/v3/raftpb"
	"go.etcd.io/etcd/server/v3/mvcc/backend"
	"go.etcd.io/etcd/server/v3/mvcc/buckets"
	"go.uber.org/zap"
)

var (
	confStateKey = []byte("confState")
)

// MustUnsafeSaveConfStateToBackend persists confState using given transaction (tx).
// confState in backend is persisted since etcd v3.5.
func MustUnsafeSaveConfStateToBackend(lg *zap.Logger, tx backend.BatchTx, confState *raftpb.ConfState) {
	confStateBytes, err := json.Marshal(confState)
	if err != nil {
		lg.Panic("Cannot marshal raftpb.ConfState", zap.Stringer("conf-state", confState), zap.Error(err))
	}

	tx.UnsafePut(buckets.Meta, confStateKey, confStateBytes)
}

// UnsafeConfStateFromBackend retrieves ConfState from the backend.
// Returns nil if confState in backend is not persisted (e.g. backend writen by <v3.5).
func UnsafeConfStateFromBackend(lg *zap.Logger, tx backend.ReadTx) *raftpb.ConfState {
	keys, vals := tx.UnsafeRange(buckets.Meta, confStateKey, nil, 0)
	if len(keys) == 0 {
		return nil
	}

	if len(keys) != 1 {
		lg.Panic(
			"unexpected number of key: "+string(confStateKey)+" when getting cluster version from backend",
			zap.Int("number-of-key", len(keys)),
		)
	}
	var confState raftpb.ConfState
	if err := json.Unmarshal(vals[0], &confState); err != nil {
		log.Panic("Cannot unmarshal confState json retrieved from the backend",
			zap.ByteString("conf-state-json", vals[0]),
			zap.Error(err))
	}
	return &confState
}
