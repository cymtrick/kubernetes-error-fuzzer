// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package traverse

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/linear"
	"gonum.org/v1/gonum/graph/internal/set"
)

var _ Graph = graph.Graph(nil)

// Graph is the subset of graph.Graph necessary for graph traversal.
type Graph interface {
	// From returns all nodes that can be reached directly
	// from the node with the given ID.
	From(id int64) []graph.Node

	// Edge returns the edge from u to v, with IDs uid and vid,
	// if such an edge exists and nil otherwise. The node v
	// must be directly reachable from u as defined by the
	// From method.
	Edge(uid, vid int64) graph.Edge
}

// BreadthFirst implements stateful breadth-first graph traversal.
type BreadthFirst struct {
	EdgeFilter func(graph.Edge) bool
	Visit      func(u, v graph.Node)
	queue      linear.NodeQueue
	visited    set.Int64s
}

// Walk performs a breadth-first traversal of the graph g starting from the given node,
// depending on the the EdgeFilter field and the until parameter if they are non-nil. The
// traversal follows edges for which EdgeFilter(edge) is true and returns the first node
// for which until(node, depth) is true. During the traversal, if the Visit field is
// non-nil, it is called with the nodes joined by each followed edge.
func (b *BreadthFirst) Walk(g Graph, from graph.Node, until func(n graph.Node, d int) bool) graph.Node {
	if b.visited == nil {
		b.visited = make(set.Int64s)
	}
	b.queue.Enqueue(from)
	b.visited.Add(from.ID())

	var (
		depth     int
		children  int
		untilNext = 1
	)
	for b.queue.Len() > 0 {
		t := b.queue.Dequeue()
		if until != nil && until(t, depth) {
			return t
		}
		tid := t.ID()
		for _, n := range g.From(tid) {
			nid := n.ID()
			if b.EdgeFilter != nil && !b.EdgeFilter(g.Edge(tid, nid)) {
				continue
			}
			if b.visited.Has(nid) {
				continue
			}
			if b.Visit != nil {
				b.Visit(t, n)
			}
			b.visited.Add(nid)
			children++
			b.queue.Enqueue(n)
		}
		if untilNext--; untilNext == 0 {
			depth++
			untilNext = children
			children = 0
		}
	}

	return nil
}

// WalkAll calls Walk for each unvisited node of the graph g using edges independent
// of their direction. The functions before and after are called prior to commencing
// and after completing each walk if they are non-nil respectively. The function
// during is called on each node as it is traversed.
func (b *BreadthFirst) WalkAll(g graph.Undirected, before, after func(), during func(graph.Node)) {
	b.Reset()
	for _, from := range g.Nodes() {
		if b.Visited(from) {
			continue
		}
		if before != nil {
			before()
		}
		b.Walk(g, from, func(n graph.Node, _ int) bool {
			if during != nil {
				during(n)
			}
			return false
		})
		if after != nil {
			after()
		}
	}
}

// Visited returned whether the node n was visited during a traverse.
func (b *BreadthFirst) Visited(n graph.Node) bool {
	return b.visited.Has(n.ID())
}

// Reset resets the state of the traverser for reuse.
func (b *BreadthFirst) Reset() {
	b.queue.Reset()
	b.visited = nil
}

// DepthFirst implements stateful depth-first graph traversal.
type DepthFirst struct {
	EdgeFilter func(graph.Edge) bool
	Visit      func(u, v graph.Node)
	stack      linear.NodeStack
	visited    set.Int64s
}

// Walk performs a depth-first traversal of the graph g starting from the given node,
// depending on the the EdgeFilter field and the until parameter if they are non-nil. The
// traversal follows edges for which EdgeFilter(edge) is true and returns the first node
// for which until(node) is true. During the traversal, if the Visit field is non-nil, it
// is called with the nodes joined by each followed edge.
func (d *DepthFirst) Walk(g Graph, from graph.Node, until func(graph.Node) bool) graph.Node {
	if d.visited == nil {
		d.visited = make(set.Int64s)
	}
	d.stack.Push(from)
	d.visited.Add(from.ID())

	for d.stack.Len() > 0 {
		t := d.stack.Pop()
		if until != nil && until(t) {
			return t
		}
		tid := t.ID()
		for _, n := range g.From(tid) {
			nid := n.ID()
			if d.EdgeFilter != nil && !d.EdgeFilter(g.Edge(tid, nid)) {
				continue
			}
			if d.visited.Has(nid) {
				continue
			}
			if d.Visit != nil {
				d.Visit(t, n)
			}
			d.visited.Add(nid)
			d.stack.Push(n)
		}
	}

	return nil
}

// WalkAll calls Walk for each unvisited node of the graph g using edges independent
// of their direction. The functions before and after are called prior to commencing
// and after completing each walk if they are non-nil respectively. The function
// during is called on each node as it is traversed.
func (d *DepthFirst) WalkAll(g graph.Undirected, before, after func(), during func(graph.Node)) {
	d.Reset()
	for _, from := range g.Nodes() {
		if d.Visited(from) {
			continue
		}
		if before != nil {
			before()
		}
		d.Walk(g, from, func(n graph.Node) bool {
			if during != nil {
				during(n)
			}
			return false
		})
		if after != nil {
			after()
		}
	}
}

// Visited returned whether the node n was visited during a traverse.
func (d *DepthFirst) Visited(n graph.Node) bool {
	return d.visited.Has(n.ID())
}

// Reset resets the state of the traverser for reuse.
func (d *DepthFirst) Reset() {
	d.stack = d.stack[:0]
	d.visited = nil
}
