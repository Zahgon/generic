// Package heap provides an implementation of a binary heap.
// A binary heap (binary min-heap) is a tree with the property that each node
// is the minimum-valued node in its subtree.
package heap

import (
	g "github.com/zyedidia/generic"
)

// Heap implements a binary heap.
type Heap[T any] struct {
	data []T
	less func(a, b T) bool
}

// New returns a new heap with the given less function.
func New[T any](less g.LessFn[T]) *Heap[T] { _ = "STUB: not implemented"; return nil }

// From returns a new heap with the given less function and initial data.
func From[T any](less g.LessFn[T], t ...T) *Heap[T] { _ = "STUB: not implemented"; return nil }

// FromSlice returns a new heap with the given less function and initial data.
// The `data` is not copied and used as the inside array.
func FromSlice[T any](less g.LessFn[T], data []T) *Heap[T] { _ = "STUB: not implemented"; return nil }

// Push pushes the given element onto the heap.
func (h *Heap[T]) Push(x T) { _ = "STUB: not implemented"; return }

// Pop removes and returns the minimum element from the heap. If the heap is
// empty, it returns zero value and false.
func (h *Heap[T]) Pop() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// Peek returns the minimum element from the heap without removing it. if the
// heap is empty, it returns zero value and false.
func (h *Heap[T]) Peek() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// Size returns the number of elements in the heap.
func (h *Heap[T]) Size() int { _ = "STUB: not implemented"; return 0 }

func down[T any](h []T, i int, less g.LessFn[T]) { _ = "STUB: not implemented"; return }

// `left < 0` in case of overflow

// find the smallest child

func up[T any](h []T, i int, less g.LessFn[T]) { _ = "STUB: not implemented"; return }
