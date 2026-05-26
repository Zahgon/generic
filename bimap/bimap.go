// Package bimap provides an implementation of a bi-directional map.
//
// It is implemented by using two Go maps, which keeps the lookup speed
// identical for both forward and reverse lookups, however it also doubles the
// memory usage of the map.
package bimap

// Of returns a new [Bimap] initiated with the keys and values
// from the given map.
func Of[K, V comparable](m map[K]V) Bimap[K, V] { _ = "STUB: not implemented"; return nil }

// Bimap is a bi-directional map where both the keys and values are indexed
// against each other, allowing performant lookup on both keys and values,
// at the cost of double the memory usage.
type Bimap[K, V comparable] struct {
	forward map[K]V
	reverse map[V]K
}

// Len returns the number of key-value pairs in this map.
func (b *Bimap[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

// Add another key-value pair to be indexed inside this map. Both the key
// and the value is indexed, to allow performant lookups on both key and value.
//
// On collisions, the old values will be overwritten.
func (b *Bimap[K, V]) Add(key K, value V) { _ = "STUB: not implemented"; return }

// RemoveForward removes a key-value pair from this map based on the key.
func (b *Bimap[K, V]) RemoveForward(key K) { _ = "STUB: not implemented"; return }

// RemoveReverse removes a key-value pair from this map based on the value.
func (b *Bimap[K, V]) RemoveReverse(value V) { _ = "STUB: not implemented"; return }

// Each loops over all the values in this map.
func (b *Bimap[K, V]) Each(f func(key K, value V)) { _ = "STUB: not implemented"; return }

// ContainsForward checks if the given key exists.
func (b *Bimap[K, V]) ContainsForward(key K) bool { _ = "STUB: not implemented"; return false }

// GetForward performs a lookup on the key to get the value.
func (b *Bimap[K, V]) GetForward(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

// ContainsReverse checks if the given value exists.
func (b *Bimap[K, V]) ContainsReverse(value V) bool { _ = "STUB: not implemented"; return false }

// GetReverse performs a lookup on the value to get the key.
func (b *Bimap[K, V]) GetReverse(value V) (K, bool) {
	_ = "STUB: not implemented"
	return *new(K), false
}

// Clear empties this bidirectional map, removing all items.
func (b *Bimap[K, V]) Clear() { _ = "STUB: not implemented"; return }

// Copy creates a shallow copy of this bidirectional map.
func (b *Bimap[K, V]) Copy() Bimap[K, V] { _ = "STUB: not implemented"; return nil }

func clear[M ~map[K]V, K comparable, V any](m M) {
	_ = "STUB: not implemented"
	// Relies on the compiler optimization introduced in Go v1.11
	// https://go.dev/doc/go1.11#performance-compiler
	return
}

func shallowCopy[M ~map[K]V, K comparable, V any](m M) M { _ = "STUB: not implemented"; return *new(M) }
