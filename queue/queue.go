// Package queue provides an implementation of a First In First Out (FIFO)
// queue. The FIFO queue is implemented using the doubly-linked list from the
// 'list' package.
package queue

import (
	"github.com/zyedidia/generic/list"
)

// Queue is a simple First In First Out (FIFO) queue.
type Queue[T any] struct {
	list   *list.List[T]
	length int
}

// New returns an empty First In First Out (FIFO) queue.
func New[T any]() *Queue[T] { _ = "STUB: not implemented"; return nil }

// Of returns a First In First Out (FIFO) queue that has been populated with
// values from an existing slice.
func Of[S ~[]E, E any](slice S) *Queue[E] { _ = "STUB: not implemented"; return nil }

// Len returns the number of items currently in the queue.
func (q *Queue[T]) Len() int {
	_ = "STUB: not implemented"

	// Enqueue inserts 'value' to the end of the queue.
	return 0
}

func (q *Queue[T]) Enqueue(value T) { _ = "STUB: not implemented"; return }

// Dequeue removes and returns the item at the front of the queue.
//
// A panic occurs if the queue is Empty.
func (q *Queue[T]) Dequeue() T { _ = "STUB: not implemented"; return *new(T) }

// TryDequeue tries to remove and return the item at the front of the queue.
//
// If the queue is empty, then false is returned as the second return value.
func (q *Queue[T]) TryDequeue() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// DequeueAll removes and returns all the items in the queue.
func (q *Queue[T]) DequeueAll() []T { _ = "STUB: not implemented"; return nil }

// Peek returns the item at the front of the queue without removing it.
//
// A panic occurs if the queue is Empty.
func (q *Queue[T]) Peek() T { _ = "STUB: not implemented"; return *new(T) }

// TryPeek tries to return the item at the front of the queue without removing it.
//
// If the queue is empty, then false is returned as the second return value.
func (q *Queue[T]) TryPeek() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// PeekAll returns all the items in the queue without removing them.
func (q *Queue[T]) PeekAll() []T { _ = "STUB: not implemented"; return nil }

// Empty returns true if the queue is empty.
func (q *Queue[T]) Empty() bool { _ = "STUB: not implemented"; return false }

// Clear empties the queue, resetting it to zero elements.
func (q *Queue[T]) Clear() { _ = "STUB: not implemented"; return }

// Copy returns a shallow copy of this queue.
func (q *Queue[T]) Copy() *Queue[T] { _ = "STUB: not implemented"; return nil }

// Each calls 'fn' on every item in the queue, starting with the least
// recently pushed element.
func (q *Queue[T]) Each(fn func(t T)) { _ = "STUB: not implemented"; return }
