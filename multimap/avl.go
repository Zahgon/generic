package multimap

import (
	g "github.com/zyedidia/generic"
	"github.com/zyedidia/generic/avl"
)

type avlMultiMap[K, V any, C valuesContainer[V]] struct {
	baseMultiMap
	keyLess    g.LessFn[K]
	keys       *avl.Tree[K, C]
	makeValues func() C
}

func (m *avlMultiMap[K, V, C]) Dimension() int { _ = "STUB: not implemented"; return 0 }

func (m *avlMultiMap[K, V, C]) Count(key K) int { _ = "STUB: not implemented"; return 0 }

func (m *avlMultiMap[K, V, C]) Has(key K) bool { _ = "STUB: not implemented"; return false }

func (m *avlMultiMap[K, V, C]) Get(key K) []V { _ = "STUB: not implemented"; return nil }

func (m *avlMultiMap[K, V, C]) Put(key K, value V) { _ = "STUB: not implemented"; return }

func (m *avlMultiMap[K, V, C]) Remove(key K, value V) { _ = "STUB: not implemented"; return }

func (m *avlMultiMap[K, V, C]) RemoveAll(key K) { _ = "STUB: not implemented"; return }

func (m *avlMultiMap[K, V, C]) Clear() { _ = "STUB: not implemented"; return }

func (m *avlMultiMap[K, V, C]) Each(fn func(key K, value V)) { _ = "STUB: not implemented"; return }

func (m *avlMultiMap[K, V, C]) EachAssociation(fn func(key K, values []V)) {
	_ = "STUB: not implemented"
	return
}

// NewAvlSlice creates a MultiMap using AVL tree and builtin slice.
//   - Value type must be comparable.
//   - Duplicate entries are permitted.
//   - Keys are sorted, but values are unsorted.
func NewAvlSlice[K any, V comparable](keyLess g.LessFn[K]) MultiMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// NewAvlSet creates a MultiMap using AVL tree and AVL set.
//   - Duplicate entries are not permitted.
//   - Both keys and values are sorted.
func NewAvlSet[K, V any](keyLess g.LessFn[K], valueLess g.LessFn[V]) MultiMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}
