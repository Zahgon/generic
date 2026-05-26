// Package array2d contains an implementation of a 2D array.
package array2d

// New initializes a 2-dimensional array with all zero values.
func New[T any](width, height int) Array2D[T] { _ = "STUB: not implemented"; return nil }

// NewFilled initializes a 2-dimensional array with a value.
func NewFilled[T any](width, height int, value T) Array2D[T] { _ = "STUB: not implemented"; return nil }

// OfJagged initializes a 2-dimensional array based on a jagged
// slice of rows of values. Values from the jagged slice that are out of bounds
// are ignored.
func OfJagged[J ~[]S, S ~[]E, E any](width, height int, jagged J) Array2D[E] {
	_ = "STUB: not implemented"
	return nil
}

// Array2D is a 2-dimensional array.
type Array2D[T any] struct {
	width, height int
	slice         []T
}

// String returns a string representation of this array.
func (a Array2D[T]) String() string { _ = "STUB: not implemented"; return "" }

// Get returns a value from the array.
//
// The function will panic on out-of-bounds access.
func (a Array2D[T]) Get(x, y int) T { _ = "STUB: not implemented"; return *new(T) }

func (a Array2D[T]) getUnchecked(x, y int) T { _ = "STUB: not implemented"; return *new(T) }

// Set sets a value in the array.
//
// The function will panic on out-of-bounds access.
func (a Array2D[T]) Set(x, y int, value T) { _ = "STUB: not implemented"; return }

func (a Array2D[T]) setUnchecked(x, y int, value T) { _ = "STUB: not implemented"; return }

// Width returns the width of this array. The maximum x value is Width()-1.
func (a Array2D[T]) Width() int {
	_ = "STUB: not implemented"

	// Height returns the height of this array. The maximum y value is Height()-1.
	return 0
}

func (a Array2D[T]) Height() int {
	_ = "STUB: not implemented"

	// Copy returns a shallow copy of this array.
	return 0
}

func (a Array2D[T]) Copy() Array2D[T] { _ = "STUB: not implemented"; return nil }

// RowSpan returns a mutable slice for part of a row. Changing values in this
// slice will affect the array.
func (a Array2D[T]) RowSpan(x1, x2, y int) []T { _ = "STUB: not implemented"; return nil }

// Row returns a mutable slice for an entire row. Changing values in this slice
// will affect the array.
func (a Array2D[T]) Row(y int) []T { _ = "STUB: not implemented"; return nil }

// Fill will assign all values inside the region to the specified value.
// The coordinates are inclusive, meaning all values from [x1,y1] including
// [x1,y1] to [x2,y2] including [x2,y2] are set.
//
// The method sorts the arguments, so x2 may be lower than x1 and y2 may be
// lower than y1.
func (a Array2D[T]) Fill(x1, y1, x2, y2 int, value T) { _ = "STUB: not implemented"; return }

func fill[E any](slice []E, value E) { _ = "STUB: not implemented"; return }

// Exponential copy to fill a slice
