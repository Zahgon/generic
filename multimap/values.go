package multimap

import (
	"github.com/zyedidia/generic/avl"
)

type valuesContainer[V any] interface {
	Empty() bool
	Size() int
	Put(value V) int
	Remove(value V) int
	List() []V
	Each(fn func(value V))
}

var (
	_ valuesContainer[int] = valuesSet[int]{}
	_ valuesContainer[int] = (*valuesSlice[int])(nil)
)

type valuesSet[V any] struct {
	t *avl.Tree[V, struct{}]
}

func (vs valuesSet[V]) Empty() bool { _ = "STUB: not implemented"; return false }

func (vs valuesSet[V]) Size() int { _ = "STUB: not implemented"; return 0 }

func (vs valuesSet[V]) has(value V) bool { _ = "STUB: not implemented"; return false }

func (vs valuesSet[V]) Put(value V) int { _ = "STUB: not implemented"; return 0 }

func (vs valuesSet[V]) Remove(value V) int { _ = "STUB: not implemented"; return 0 }

func (vs valuesSet[V]) List() (values []V) { _ = "STUB: not implemented"; return nil }

func (vs valuesSet[V]) Each(fn func(value V)) { _ = "STUB: not implemented"; return }

type valuesSlice[V comparable] []V

func (vs *valuesSlice[V]) Empty() bool { _ = "STUB: not implemented"; return false }

func (vs *valuesSlice[V]) Size() int { _ = "STUB: not implemented"; return 0 }

func (vs *valuesSlice[V]) Put(value V) int { _ = "STUB: not implemented"; return 0 }

func (vs *valuesSlice[V]) Remove(value V) int { _ = "STUB: not implemented"; return 0 }

func (vs *valuesSlice[V]) List() []V { _ = "STUB: not implemented"; return nil }

func (vs *valuesSlice[V]) Each(fn func(value V)) { _ = "STUB: not implemented"; return }
