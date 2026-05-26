package ulist

import (
	"github.com/zyedidia/generic/list"
)

// A UListIter points to an element in the UList.
type UListIter[V any] struct {
	node  *list.Node[ulistBlk[V]]
	index int
}

// newIterFront returns a UListIter pointing to the first entry in 'ul'.
// If 'ul' is empty, an invalid iterator is returned.
func newIterFront[V any](ul *UList[V]) *UListIter[V] { _ = "STUB: not implemented"; return nil }

// newIterBack returns a UListIter pointing to the last entry in 'ul'.
// If 'ul' is empty, an invalid iterator is returned.
func newIterBack[V any](ul *UList[V]) *UListIter[V] { _ = "STUB: not implemented"; return nil }

// IsValid returns true if the iterator points to a valid entry in the UList.
func (iter *UListIter[V]) IsValid() bool { _ = "STUB: not implemented"; return false }

// Get returns the entry in the UList that the 'iter' is pointing to.
// This call should only ever be made when iter.IsValid() is true.
func (iter *UListIter[V]) Get() V { _ = "STUB: not implemented"; return *new(V) }

// Next moves the iterator one step forward and returns true if the iterator is valid.
func (iter *UListIter[V]) Next() bool { _ = "STUB: not implemented"; return false }

// By not going past len, we can recover to the end using Prev().

// Prev moves the iterator one step back and returns true if the iterator is valid.
func (iter *UListIter[V]) Prev() bool { _ = "STUB: not implemented"; return false }

// By not going further past -1, we can recover to the begin using Next().
