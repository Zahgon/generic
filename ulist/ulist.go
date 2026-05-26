package ulist

import (
	"github.com/zyedidia/generic/list"
)

//          ---------    ---------
//  UList:  | Block | <-> | Block | <-> ...
//          ---------    ---------
// A UList is represented internally as a list.List (doubly linked list)
// of pointers to blocks of entries.
//
// Block:  []V{ ... }
// A Block is a slice of entries and forms the Node in the list.List.

// Type alias for a block of entries.
type ulistBlk[V any] []V

// UList implements a doubly-linked unolled list.
type UList[V any] struct {
	ll              list.List[ulistBlk[V]]
	entriesPerBlock int
	size            int
}

// New returns an empty unrolled linked list.
// 'entriesPerBlock' is the number of entries to store in each block.
// This value should ideally be the size of a cache-line or multiples there-of.
// See: https://en.wikipedia.org/wiki/Unrolled_linked_list
func New[V any](entriesPerBlock int) *UList[V] { _ = "STUB: not implemented"; return nil }

// Size returns the number of entries in 'ul'.
func (ul *UList[V]) Size() int {
	_ = "STUB: not implemented"

	// PushBack adds 'v' to the end of the ulist.
	return 0
}

func (ul *UList[V]) PushBack(v V) { _ = "STUB: not implemented"; return }

// PushFront adds 'v' to the beginning of the ulist.
func (ul *UList[V]) PushFront(v V) { _ = "STUB: not implemented"; return }

// Begin returns an UListIter pointing to the first entry in the UList.
func (ul *UList[V]) Begin() *UListIter[V] { _ = "STUB: not implemented"; return nil }

// End returns an UListIter pointing to the last entry in the UList.
func (ul *UList[V]) End() *UListIter[V] { _ = "STUB: not implemented"; return nil }

// AddAfter adds 'v' to 'ul' after the entry pointed to by 'iter'.
// 'iter' is expected to be valid, i.e. iter->IsValid() == true.
// 'iter' is updated to now point to the new entry added, such that
// iter->Get() == 'v'.
func (ul *UList[V]) AddAfter(iter *UListIter[V], v V) {
	_ = "STUB: not implemented"

	// Adding to a block with spare capacity.
	return
}

// Adding to an already full block.

// When adding to the end of a block, 'v' is the overflow.

// When adding 'v' in the middle, the last entry in the block is the overflow.

// Slide entries beyond the write index right by one spot and write the value.

// AddBefore adds 'v' to 'ul' before the entry pointed to by 'iter'.
// 'iter' is expected to be valid, i.e. iter->IsValid() == true.
// 'iter' is updated to now point to the new entry added, such that
// iter->Get() == 'v'.
func (ul *UList[V]) AddBefore(iter *UListIter[V], v V) { _ = "STUB: not implemented"; return }

// Remove deletes the entry in 'ul' pointed to by 'iter'.
// 'iter' is moved forward in the process. i.e. iter.Get() returns the element in 'ul'
// that occurs after the deleted entry.
func (ul *UList[V]) Remove(iter *UListIter[V]) { _ = "STUB: not implemented"; return }

// Block got emptied.

func hasCapacity[V any](llNode *list.Node[ulistBlk[V]]) bool {
	_ = "STUB: not implemented"
	return false
}

func (ul *UList[V]) newBlock() ulistBlk[V] { _ = "STUB: not implemented"; return nil }

func (ul *UList[V]) prependToBlock(v V, blkPtr *ulistBlk[V]) { _ = "STUB: not implemented"; return }

// 'append' returns a slice with capacity of the first variable.
// To maintain the propoer capacity, we use 'tmp' with an explicitly defined capacity.

func (iter *UListIter[V]) addOverflowToNextBlock(ul *UList[V], v V) {
	_ = "STUB: not implemented"
	return
}
