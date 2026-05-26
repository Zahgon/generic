// Package mapset provides an implementation of a set using the built-in map.
package mapset

// Set implements a hashset, using the hashmap as the underlying storage.
type Set[K comparable] struct {
	m map[K]struct{}
}

// New returns an empty hashset.
func New[K comparable]() Set[K] { _ = "STUB: not implemented"; return nil }

// Of returns a new hashset initialized with the given 'vals'
func Of[K comparable](vals ...K) Set[K] { _ = "STUB: not implemented"; return nil }

// Put adds 'val' to the set.
func (s Set[K]) Put(val K) { _ = "STUB: not implemented"; return }

// Has returns true only if 'val' is in the set.
func (s Set[K]) Has(val K) bool { _ = "STUB: not implemented"; return false }

// Remove removes 'val' from the set.
func (s Set[K]) Remove(val K) {
	_ = "STUB: not implemented"

	// Clear removes all elements from the set.
	return
}

func (s Set[K]) Clear() {
	_ = "STUB: not implemented"

	// Size returns the number of elements in the set.
	return
}

func (s Set[K]) Size() int {
	_ = "STUB: not implemented"

	// Each calls 'fn' on every item in the set in no particular order.
	return 0
}

func (s Set[K]) Each(fn func(key K)) { _ = "STUB: not implemented"; return }
