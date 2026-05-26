// Package interval provides an implementation of an interval tree built using
// an augmented AVL tree. An interval tree stores values associated with
// intervals, and can efficiently determine which intervals overlap with
// others. All intervals must have a unique starting position. It supports the
// following operations, where 'n' is the number of
// intervals in the tree:
//
//   - Put: add an interval to the tree. Complexity: O(lg n).
//
//   - Get: find an interval with a given starting position. Complexity O(lg n).
//
//   - Overlaps: find all intervals that overlap with a given interval. Complexity:
//     O(lg n + m), where 'm' is the size of the result (number of overlapping
//     intervals found).
//
//   - Remove: remove the interval at a given position. Complexity: O(lg n).
package interval

import (
	"golang.org/x/exp/constraints"
)

type KV[I constraints.Ordered, V any] struct {
	Low, High I
	Val       V
}

func newKV[I constraints.Ordered, V any](n *node[I, V]) KV[I, V] {
	_ = "STUB: not implemented"
	return nil
}

// intrvl represents an interval over [low, high).
type intrvl[I constraints.Ordered] struct {
	low, high I
}

func newIntrvl[I constraints.Ordered](low, high I) intrvl[I] { _ = "STUB: not implemented"; return nil }

func overlaps[I constraints.Ordered](i1 intrvl[I], i2 intrvl[I]) bool {
	_ = "STUB: not implemented"
	return false
}

// Tree implements an interval tree. All intervals must have unique starting
// positions. Every low bound if an interval is inclusive, while high is
// exclusive.
type Tree[I constraints.Ordered, V any] struct {
	root *node[I, V]
}

// New returns an empty interval tree.
func New[I constraints.Ordered, V any]() *Tree[I, V] { _ = "STUB: not implemented"; return nil }

// Add associates the interval [low, high) with value.
//
// If an interval starting at low already exists in t, this method doesn't
// perform any change of the tree, but returns the conflicting interval.
func (t *Tree[I, V]) Add(low, high I, value V) (KV[I, V], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Put associates the interval [low, high) with value.
//
// If an interval starting at low already exists, this method will replace it.
// In such a case the conflicting (replaced) interval is returned.
func (t *Tree[I, V]) Put(low, high I, value V) (KV[I, V], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Overlaps returns all values that overlap with the given range. List returned
// is sorted by low positions of intervals.
func (t *Tree[I, V]) Overlaps(low, high I) []KV[I, V] { _ = "STUB: not implemented"; return nil }

// Remove deletes the interval starting at low. The removed interval is
// returned. If no such interval existed in a tree, the returned value is false.
func (t *Tree[I, V]) Remove(low I) (KV[I, V], bool) { _ = "STUB: not implemented"; return nil, false }

// Get returns the interval and value associated with the interval starting at
// low, or false if no such value exists.
func (t *Tree[I, V]) Get(low I) (KV[I, V], bool) { _ = "STUB: not implemented"; return nil, false }

// Each calls 'fn' on every element in the tree, and its corresponding
// interval, in order sorted by starting position.
func (t *Tree[I, V]) Each(fn func(low, high I, val V)) {
	_ = "STUB: not implemented"

	// Height returns the height of the tree.
	return
}

func (t *Tree[I, V]) Height() int { _ = "STUB: not implemented"; return 0 }

// Size returns the number of elements in the tree.
func (t *Tree[I, V]) Size() int { _ = "STUB: not implemented"; return 0 }

type node[I constraints.Ordered, V any] struct {
	key   intrvl[I]
	value V

	height int
	left   *node[I, V]
	right  *node[I, V]

	// max is highest upper bound of all intervals stored in subtree which
	// node as its root.
	max I
}

// insert inserts interval key associated with value value to the tree.
//
// If interval starting at key.low already exists in a tree, behaviour of this
// method is defined by overwrite parameter. If it's true, the value is
// replaced. Otherwise whole subtree is left unchanged.
//
// This method returns new root node of a subtree rooted in n after insertion,
// an interval starting at key.low which already exists in the subtree and a
// flag if such an interval exists.
func (n *node[I, V]) insert(
	key intrvl[I],
	value V,
	overwrite bool,
) (*node[I, V], KV[I, V], bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (n *node[I, V]) updateMax() { _ = "STUB: not implemented"; return }

// remove removes interval starting at pos from a subtree. This function returns
// the new root of subtree rooted in n after pos removal, the KV removed and an
// information if any deletion happened (i.e. if interval starting at pos
// exists).
func (n *node[I, V]) remove(low I) (*node[I, V], KV[I, V], bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// removeThis deletes n from subtree rooted in n and returns new root of the
// subtree.
func (n *node[I, V]) removeThis() *node[I, V] {
	_ = "STUB: not implemented"
	// This can return nil if n has no children (n.right == nil).
	return nil
}

func (n *node[I, V]) search(low I) *node[I, V] { _ = "STUB: not implemented"; return nil }

func (n *node[I, V]) overlaps(key intrvl[I], result []KV[I, V]) []KV[I, V] {
	_ = "STUB: not implemented"
	return nil
}

func (n *node[I, V]) each(fn func(low, high I, val V)) { _ = "STUB: not implemented"; return }

func (n *node[I, V]) getHeight() int { _ = "STUB: not implemented"; return 0 }

func (n *node[I, V]) recalculateHeight() { _ = "STUB: not implemented"; return }

func (n *node[I, V]) rebalanceTree() *node[I, V] { _ = "STUB: not implemented"; return nil }

func (n *node[I, V]) rotateLeft() *node[I, V] { _ = "STUB: not implemented"; return nil }

func (n *node[I, V]) rotateRight() *node[I, V] { _ = "STUB: not implemented"; return nil }

func (n *node[I, V]) findSmallest() *node[I, V] { _ = "STUB: not implemented"; return nil }

func (n *node[I, V]) size() int { _ = "STUB: not implemented"; return 0 }
