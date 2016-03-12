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

package storage

import (
	"bytes"
	"errors"
	"fmt"
	"log"

	"github.com/google/btree"
)

var (
	ErrRevisionNotFound = errors.New("stroage: revision not found")
)

// keyIndex stores the revisions of a key in the backend.
// Each keyIndex has at least one key generation.
// Each generation might have several key versions.
// Tombstone on a key appends an tombstone version at the end
// of the current generation and creates a new empty generation.
// Each version of a key has an index pointing to the backend.
//
// For example: put(1.0);put(2.0);tombstone(3.0);put(4.0);tombstone(5.0) on key "foo"
// generate a keyIndex:
// key:     "foo"
// rev: 5
// generations:
//    {empty}
//    {4.0, 5.0(t)}
//    {1.0, 2.0, 3.0(t)}
//
// Compact a keyIndex removes the versions with smaller or equal to
// rev except the largest one. If the generation becomes empty
// during compaction, it will be removed. if all the generations get
// removed, the keyIndex should be removed.

// For example:
// compact(2) on the previous example
// generations:
//    {empty}
//    {4.0, 5.0(t)}
//    {2.0, 3.0(t)}
//
// compact(4)
// generations:
//    {empty}
//    {4.0, 5.0(t)}
//
// compact(5):
// generations:
//    {empty} -> key SHOULD be removed.
//
// compact(6):
// generations:
//    {empty} -> key SHOULD be removed.
type keyIndex struct {
	key         []byte
	modified    revision // the main rev of the last modification
	generations []generation
}

// put puts a revision to the keyIndex.
func (ki *keyIndex) put(main int64, sub int64) {
	rev := revision{main: main, sub: sub}

	if !rev.GreaterThan(ki.modified) {
		log.Panicf("store.keyindex: put with unexpected smaller revision [%v / %v]", rev, ki.modified)
	}
	if len(ki.generations) == 0 {
		ki.generations = append(ki.generations, generation{})
	}
	g := &ki.generations[len(ki.generations)-1]
	if len(g.revs) == 0 { // create a new key
		keysGauge.Inc()
		g.created = rev
	}
	g.revs = append(g.revs, rev)
	g.ver++
	ki.modified = rev
}

func (ki *keyIndex) restore(created, modified revision, ver int64) {
	if len(ki.generations) != 0 {
		log.Panicf("store.keyindex: cannot restore non-empty keyIndex")
	}

	ki.modified = modified
	g := generation{created: created, ver: ver, revs: []revision{modified}}
	ki.generations = append(ki.generations, g)
	keysGauge.Inc()
}

// tombstone puts a revision, pointing to a tombstone, to the keyIndex.
// It also creates a new empty generation in the keyIndex.
// It returns ErrRevisionNotFound when tombstone on an empty generation.
func (ki *keyIndex) tombstone(main int64, sub int64) error {
	if ki.isEmpty() {
		log.Panicf("store.keyindex: unexpected tombstone on empty keyIndex %s", string(ki.key))
	}
	if ki.generations[len(ki.generations)-1].isEmpty() {
		return ErrRevisionNotFound
	}
	ki.put(main, sub)
	ki.generations = append(ki.generations, generation{})
	keysGauge.Dec()
	return nil
}

// get gets the modified, created revision and version of the key that satisfies the given atRev.
// Rev must be higher than or equal to the given atRev.
func (ki *keyIndex) get(atRev int64) (modified, created revision, ver int64, err error) {
	if ki.isEmpty() {
		log.Panicf("store.keyindex: unexpected get on empty keyIndex %s", string(ki.key))
	}
	g := ki.findGeneration(atRev)
	if g.isEmpty() {
		return revision{}, revision{}, 0, ErrRevisionNotFound
	}

	f := func(rev revision) bool {
		if rev.main <= atRev {
			return false
		}
		return true
	}

	n := g.walk(f)
	if n != -1 {
		return g.revs[n], g.created, g.ver - int64(len(g.revs)-n-1), nil
	}

	return revision{}, revision{}, 0, ErrRevisionNotFound
}

// since returns revisions since the given rev. Only the revision with the
// largest sub revision will be returned if multiple revisions have the same
// main revision.
func (ki *keyIndex) since(rev int64) []revision {
	if ki.isEmpty() {
		log.Panicf("store.keyindex: unexpected get on empty keyIndex %s", string(ki.key))
	}
	since := revision{rev, 0}
	var gi int
	// find the generations to start checking
	for gi = len(ki.generations) - 1; gi > 0; gi-- {
		g := ki.generations[gi]
		if g.isEmpty() {
			continue
		}
		if since.GreaterThan(g.created) {
			break
		}
	}

	var revs []revision
	var last int64
	for ; gi < len(ki.generations); gi++ {
		for _, r := range ki.generations[gi].revs {
			if since.GreaterThan(r) {
				continue
			}
			if r.main == last {
				// replace the revision with a new one that has higher sub value,
				// because the original one should not be seen by external
				revs[len(revs)-1] = r
				continue
			}
			revs = append(revs, r)
			last = r.main
		}
	}
	return revs
}

// compact compacts a keyIndex by removing the versions with smaller or equal
// revision than the given atRev except the largest one (If the largest one is
// a tombstone, it will not be kept).
// If a generation becomes empty during compaction, it will be removed.
func (ki *keyIndex) compact(atRev int64, available map[revision]struct{}) {
	if ki.isEmpty() {
		log.Panicf("store.keyindex: unexpected compact on empty keyIndex %s", string(ki.key))
	}

	// walk until reaching the first revision that has an revision smaller or equal to
	// the atRev.
	// add it to the available map
	f := func(rev revision) bool {
		if rev.main <= atRev {
			available[rev] = struct{}{}
			return false
		}
		return true
	}

	i, g := 0, &ki.generations[0]
	// find first generation includes atRev or created after atRev
	for i < len(ki.generations)-1 {
		if tomb := g.revs[len(g.revs)-1].main; tomb > atRev {
			break
		}
		i++
		g = &ki.generations[i]
	}

	if !g.isEmpty() {
		n := g.walk(f)
		// remove the previous contents.
		if n != -1 {
			g.revs = g.revs[n:]
		}
		// remove any tombstone
		if len(g.revs) == 1 && i != len(ki.generations)-1 {
			delete(available, g.revs[0])
			i++
		}
	}
	// remove the previous generations.
	ki.generations = ki.generations[i:]
	return
}

func (ki *keyIndex) isEmpty() bool {
	return len(ki.generations) == 1 && ki.generations[0].isEmpty()
}

// findGeneration finds out the generation of the keyIndex that the
// given rev belongs to. If the given rev is at the gap of two generations,
// which means that the key does not exist at the given rev, it returns nil.
func (ki *keyIndex) findGeneration(rev int64) *generation {
	lastg := len(ki.generations) - 1
	cg := lastg

	for cg >= 0 {
		if len(ki.generations[cg].revs) == 0 {
			cg--
			continue
		}
		g := ki.generations[cg]
		if cg != lastg {
			if tomb := g.revs[len(g.revs)-1].main; tomb <= rev {
				return nil
			}
		}
		if g.revs[0].main <= rev {
			return &ki.generations[cg]
		}
		cg--
	}
	return nil
}

func (a *keyIndex) Less(b btree.Item) bool {
	return bytes.Compare(a.key, b.(*keyIndex).key) == -1
}

func (a *keyIndex) equal(b *keyIndex) bool {
	if !bytes.Equal(a.key, b.key) {
		return false
	}
	if a.modified != b.modified {
		return false
	}
	if len(a.generations) != len(b.generations) {
		return false
	}
	for i := range a.generations {
		ag, bg := a.generations[i], b.generations[i]
		if !ag.equal(bg) {
			return false
		}
	}
	return true
}

func (ki *keyIndex) String() string {
	var s string
	for _, g := range ki.generations {
		s += g.String()
	}
	return s
}

// generation contains multiple revisions of a key.
type generation struct {
	ver     int64
	created revision // when the generation is created (put in first revision).
	revs    []revision
}

func (g *generation) isEmpty() bool { return g == nil || len(g.revs) == 0 }

// walk walks through the revisions in the generation in descending order.
// It passes the revision to the given function.
// walk returns until: 1. it finishes walking all pairs 2. the function returns false.
// walk returns the position at where it stopped. If it stopped after
// finishing walking, -1 will be returned.
func (g *generation) walk(f func(rev revision) bool) int {
	l := len(g.revs)
	for i := range g.revs {
		ok := f(g.revs[l-i-1])
		if !ok {
			return l - i - 1
		}
	}
	return -1
}

func (g *generation) String() string {
	return fmt.Sprintf("g: created[%d] ver[%d], revs %#v\n", g.created, g.ver, g.revs)
}

func (a generation) equal(b generation) bool {
	if a.ver != b.ver {
		return false
	}
	if len(a.revs) != len(b.revs) {
		return false
	}

	for i := range a.revs {
		ar, br := a.revs[i], b.revs[i]
		if ar != br {
			return false
		}
	}
	return true
}
