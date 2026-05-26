// Package stack provides an implementation of a LIFO stack built using a
// resizing array.
package stack

// Stack implements a LIFO stack with peeking.
type Stack[T any] struct {
	entries []T
}

// New returns an empty stack.
func New[T any]() *Stack[T] { _ = "STUB: not implemented"; return nil }

// Push places 'value' at the top of the stack.
func (s *Stack[T]) Push(value T) { _ = "STUB: not implemented"; return }

// Pop removes the stack's top element and returns it. If the stack is empty it
// returns the zero value.
func (s *Stack[T]) Pop() (t T) { _ = "STUB: not implemented"; return *new(T) }

// Peek returns the stack's top element but does not remove it. If the stack is
// empty the zero value is returned.
func (s *Stack[T]) Peek() (t T) { _ = "STUB: not implemented"; return *new(T) }

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int { _ = "STUB: not implemented"; return 0 }

// Copy returns a copy of this stack.
func (s *Stack[T]) Copy() *Stack[T] { _ = "STUB: not implemented"; return nil }
