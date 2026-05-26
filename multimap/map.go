package multimap

import (
	g "github.com/zyedidia/generic"
)

type mapMultiMap[K comparable, V any, C valuesContainer[V]] struct {
	baseMultiMap
	keys       map[K]C
	makeValues func() C
}

func (m *mapMultiMap[K, V, C]) Dimension() int { _ = "STUB: not implemented"; return 0 }

func (m *mapMultiMap[K, V, C]) Count(key K) int { _ = "STUB: not implemented"; return 0 }

func (m *mapMultiMap[K, V, C]) Has(key K) bool { _ = "STUB: not implemented"; return false }

func (m *mapMultiMap[K, V, C]) Get(key K) []V { _ = "STUB: not implemented"; return nil }

func (m *mapMultiMap[K, V, C]) Put(key K, value V) { _ = "STUB: not implemented"; return }

func (m *mapMultiMap[K, V, C]) Remove(key K, value V) { _ = "STUB: not implemented"; return }

func (m *mapMultiMap[K, V, C]) RemoveAll(key K) { _ = "STUB: not implemented"; return }

func (m *mapMultiMap[K, V, C]) Clear() { _ = "STUB: not implemented"; return }

func (m *mapMultiMap[K, V, C]) Each(fn func(key K, value V)) { _ = "STUB: not implemented"; return }

func (m *mapMultiMap[K, V, C]) EachAssociation(fn func(key K, values []V)) {
	_ = "STUB: not implemented"
	return
}

// NewMapSlice creates a MultiMap using builtin map and builtin slice.
//   - Both key type and value type must be comparable.
//   - Duplicate entries are permitted.
//   - Both keys and values are unsorted.
func NewMapSlice[K, V comparable]() MultiMap[K, V] { _ = "STUB: not implemented"; return nil }

// NewMapSet creates a MultiMap using builtin map and AVL set.
//   - Key type must be comparable.
//   - Duplicate entries are not permitted.
//   - Values are sorted, but keys are unsorted.
func NewMapSet[K comparable, V any](valueLess g.LessFn[V]) MultiMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}
