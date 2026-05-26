package set

import (
	"github.com/zyedidia/generic"
)

func NewMapset[K comparable](in ...K) Set[K] { _ = "STUB: not implemented"; return nil }

func NewHashset[K comparable](cap uint64, equals generic.EqualsFn[K], hash generic.HashFn[K], in ...K) Set[K] {
	_ = "STUB: not implemented"
	return nil
}

func NewSet[K comparable, S func() SetOf[K]](con S, in ...K) Set[K] {
	_ = "STUB: not implemented"
	return nil
}

type SetOf[K comparable] interface {
	Put(val K)
	Has(val K) bool
	Remove(val K)
	Clear()
	Size() int
	Each(fn func(key K))
}

type Set[K comparable] struct {
	SetOf[K]
	new func() SetOf[K]
}

func (s Set[K]) Intersection(others ...SetOf[K]) Set[K] { _ = "STUB: not implemented"; return nil }

func (s Set[K]) Difference(others ...SetOf[K]) Set[K] { _ = "STUB: not implemented"; return nil }

func (s Set[K]) Union(others ...SetOf[K]) Set[K] { _ = "STUB: not implemented"; return nil }

func (s Set[K]) ConstSymmetricDifference(with ...K) Set[K] { _ = "STUB: not implemented"; return nil }

func (s Set[K]) ConstIntersection(with ...K) Set[K] { _ = "STUB: not implemented"; return nil }

func (s Set[K]) ConstDifference(with ...K) Set[K] { _ = "STUB: not implemented"; return nil }

func (s Set[K]) ConstUnion(with ...K) Set[K] { _ = "STUB: not implemented"; return nil }

func (s Set[K]) Clone() Set[K] { _ = "STUB: not implemented"; return nil }

func (s Set[K]) String() string { _ = "STUB: not implemented"; return "" }

func (s Set[K]) Map() map[K]struct{} { _ = "STUB: not implemented"; return nil }

func (s Set[K]) SymmetricDifference(others ...SetOf[K]) Set[K] {
	_ = "STUB: not implemented"
	return nil
}

func (s Set[K]) InPlaceIntersection(others ...SetOf[K]) Set[K] {
	_ = "STUB: not implemented"
	return nil
}

func (s Set[K]) InPlaceDifference(others ...SetOf[K]) Set[K] { _ = "STUB: not implemented"; return nil }

func (s Set[K]) InPlaceUnion(others ...SetOf[K]) Set[K] { _ = "STUB: not implemented"; return nil }

func (s Set[K]) Keys() []K { _ = "STUB: not implemented"; return nil }

func (s Set[K]) IsDisjoint(other SetOf[K]) bool {
	_ = "STUB: not implemented"
	// TODO: maybe optimize?
	return false
}

func (s Set[K]) IsSubset(of SetOf[K]) bool { _ = "STUB: not implemented"; return false }

func (s Set[K]) IsSuperset(of SetOf[K]) bool { _ = "STUB: not implemented"; return false }

func (s Set[K]) Equal(to SetOf[K]) bool { _ = "STUB: not implemented"; return false }

func (s Set[K]) IsProperSubset(to SetOf[K]) bool { _ = "STUB: not implemented"; return false }

func (s Set[K]) IsProperSuperset(to SetOf[K]) bool { _ = "STUB: not implemented"; return false }
