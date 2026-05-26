// Package hashmap provides an implementation of a hashmap. The map uses linear
// probing and automatically resizes. The map can also be efficiently copied,
// and will perform copies lazily, using copy-on-write. However, the
// copy-on-write will copy the entire map after the first write. One can imagine
// a more efficient implementation that would split the map into chunks and use
// copy-on-write selectively for each chunk.
package hashmap

import (
	g "github.com/zyedidia/generic"
)

type entry[K, V any] struct {
	key    K
	filled bool
	value  V
}

// A Map is a hashmap that supports copying via copy-on-write.
type Map[K, V any] struct {
	entries  []entry[K, V]
	capacity uint64
	length   uint64
	readonly bool

	ops ops[K]
}

type ops[T any] struct {
	equals func(a, b T) bool
	hash   func(t T) uint64
}

func pow2ceil(num uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// New constructs a new map with the given capacity.
func New[K, V any](capacity uint64, equals g.EqualsFn[K], hash g.HashFn[K]) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Get returns the value stored for this key, or false if there is no such
// value.
func (m *Map[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (m *Map[K, V]) resize(newcap uint64) { _ = "STUB: not implemented"; return }

// Put maps the given key to the given value. If the key already exists its
// value will be overwritten with the new value.
func (m *Map[K, V]) Put(key K, val V) { _ = "STUB: not implemented"; return }

func (m *Map[K, V]) remove(idx uint64) { _ = "STUB: not implemented"; return }

// Remove removes the specified key-value pair from the map.
func (m *Map[K, V]) Remove(key K) { _ = "STUB: not implemented"; return }

// halves the array if it is 12.5% full or less

// Clear removes all key-value pairs from the map.
func (m *Map[K, V]) Clear() { _ = "STUB: not implemented"; return }

// Size returns the number of items in the map.
func (m *Map[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

// Copy returns a copy of this map. The copy will not allocate any memory until
// the first write, so any number of read-only copies can be made without any
// additional allocations.
func (m *Map[K, V]) Copy() *Map[K, V] { _ = "STUB: not implemented"; return nil }

// Each calls 'fn' on every key-value pair in the hashmap in no particular
// order.
func (m *Map[K, V]) Each(fn func(key K, val V)) { _ = "STUB: not implemented"; return }
