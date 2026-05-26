// Package avl provides an implementation of an AVL tree. An AVL tree is a
// self-balancing binary search tree. It stores key-value pairs that are sorted
// based on the key, and maintains that the tree is always balanced, ensuring
// logarithmic-time for all operations.
package avl

import (
	g "github.com/zyedidia/generic"
)

// Tree implements an AVL tree.
type Tree[K, V any] struct {
	root *node[K, V]
	less g.LessFn[K]
}

// New returns an empty AVL tree.
func New[K, V any](less g.LessFn[K]) *Tree[K, V] { _ = "STUB: not implemented"; return nil }

// Put associates 'key' with 'value'.
func (t *Tree[K, V]) Put(key K, value V) { _ = "STUB: not implemented"; return }

// Remove removes the value associated with 'key'.
func (t *Tree[K, V]) Remove(key K) { _ = "STUB: not implemented"; return }

// Get returns the value associated with 'key'.
func (t *Tree[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

// Each calls 'fn' on every node in the tree in order
func (t *Tree[K, V]) Each(fn func(key K, val V)) {
	_ = "STUB: not implemented"

	// Height returns the height of the tree.
	return
}

func (t *Tree[K, V]) Height() int { _ = "STUB: not implemented"; return 0 }

// Size returns the number of elements in the tree.
func (t *Tree[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

type node[K, V any] struct {
	key   K
	value V

	height int
	left   *node[K, V]
	right  *node[K, V]
}

func (n *node[K, V]) add(key K, value V, less g.LessFn[K]) *node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (n *node[K, V]) remove(key K, less g.LessFn[K]) *node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (n *node[K, V]) search(key K, less g.LessFn[K]) *node[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (n *node[K, V]) each(fn func(key K, val V)) { _ = "STUB: not implemented"; return }

func (n *node[K, V]) getHeight() int { _ = "STUB: not implemented"; return 0 }

func (n *node[K, V]) recalculateHeight() { _ = "STUB: not implemented"; return }

func (n *node[K, V]) rebalanceTree() *node[K, V] { _ = "STUB: not implemented"; return nil }

func (n *node[K, V]) rotateLeft() *node[K, V] { _ = "STUB: not implemented"; return nil }

func (n *node[K, V]) rotateRight() *node[K, V] { _ = "STUB: not implemented"; return nil }

func (n *node[K, V]) findSmallest() *node[K, V] { _ = "STUB: not implemented"; return nil }

func (n *node[K, V]) size() int { _ = "STUB: not implemented"; return 0 }
