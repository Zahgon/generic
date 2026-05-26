// Package btree provides an implementation of a B-tree. A B-tree is a
// logarithmic search tree that maintains key-value pairs in sorted order. It
// is not binary because it stores more than 2 data entries per node. The
// branching factor for this tree is 64.
package btree

import (
	g "github.com/zyedidia/generic"
)

const maxChildren = 64 // must be even and > 2

// Adapted from the B-tree implementation in Algorithms, 4th ed., by Robert
// Sedgewick and Kevin Wayne.
// https://algs4.cs.princeton.edu/62btree/BTree.java.html.

// Tree implements a B-tree.
type Tree[K, V any] struct {
	root   *node[K, V]
	height int
	n      int

	less g.LessFn[K]
}

type node[K, V any] struct {
	m        int
	children [maxChildren]entry[K, V]
}

type entry[K, V any] struct {
	key   K
	val   V
	valid bool
	next  *node[K, V]
}

// New returns an empty B-tree.
func New[K, V any](less g.LessFn[K]) *Tree[K, V] { _ = "STUB: not implemented"; return nil }

// Size returns the number of elements in the tree.
func (t *Tree[K, V]) Size() int {
	_ = "STUB: not implemented"

	// Get returns the value associated with 'key'.
	return 0
}

func (t *Tree[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (t *Tree[K, V]) search(x *node[K, V], key K, height int) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// leaf node

// internal node

// Put associates 'key' with 'val'.
func (t *Tree[K, V]) Put(key K, val V) { _ = "STUB: not implemented"; return }

// Remove removes the value associated with 'key'.
func (t *Tree[K, V]) Remove(key K) { _ = "STUB: not implemented"; return }

// insert a tombstone to remove an existing value

func (t *Tree[K, V]) insert(h *node[K, V], key K, val V, height int, valid bool) *node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// leaf node

// internal node

func (t *Tree[K, V]) split(h *node[K, V]) *node[K, V] { _ = "STUB: not implemented"; return nil }

// Each calls 'fn' on every node in the tree in order.
func (t *Tree[K, V]) Each(fn func(key K, val V)) { _ = "STUB: not implemented"; return }

func (t *Tree[K, V]) each(n *node[K, V], height int, fn func(key K, val V)) {
	_ = "STUB: not implemented"
	return
}
