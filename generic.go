package generic

import (
	"golang.org/x/exp/constraints"
)

// EqualsFn is a function that returns whether 'a' and 'b' are equal.
type EqualsFn[T any] func(a, b T) bool

// LessFn is a function that returns whether 'a' is less than 'b'.
type LessFn[T any] func(a, b T) bool

// HashFn is a function that returns the hash of 't'.
type HashFn[T any] func(t T) uint64

// Equals wraps the '==' operator for comparable types.
func Equals[T comparable](a, b T) bool {
	_ = "STUB: not implemented"

	// Less wraps the '<' operator for ordered types.
	return false
}

func Less[T constraints.Ordered](a, b T) bool {
	_ = "STUB: not implemented"

	// Compare uses a less function to determine the ordering of 'a' and 'b'. It returns:
	//
	// * -1 if a < b
	//
	// * 1 if a > b
	//
	// * 0 if a == b
	return false
}

func Compare[T any](a, b T, less LessFn[T]) int { _ = "STUB: not implemented"; return 0 }

// Max returns the max of a and b.
func Max[T constraints.Ordered](a, b T) T { _ = "STUB: not implemented"; return *new(T) }

// Min returns the min of a and b.
func Min[T constraints.Ordered](a, b T) T { _ = "STUB: not implemented"; return *new(T) }

// Clamp returns x constrained within [lo:hi] range.
// If x compares less than lo, returns lo; otherwise if hi compares less than x, returns hi; otherwise returns v.
func Clamp[T constraints.Ordered](x, lo, hi T) T { _ = "STUB: not implemented"; return *new(T) }

// MaxFunc returns the max of a and b using the less func.
func MaxFunc[T any](a, b T, less LessFn[T]) T { _ = "STUB: not implemented"; return *new(T) }

// MinFunc returns the min of a and b using the less func.
func MinFunc[T any](a, b T, less LessFn[T]) T { _ = "STUB: not implemented"; return *new(T) }

// ClampFunc returns x constrained within [lo:hi] range using the less func.
// If x compares less than lo, returns lo; otherwise if hi compares less than x, returns hi; otherwise returns v.
func ClampFunc[T any](x, lo, hi T, less LessFn[T]) T { _ = "STUB: not implemented"; return *new(T) }

func HashUint64(u uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func HashUint32(u uint32) uint64 { _ = "STUB: not implemented"; return 0 }

func HashUint16(u uint16) uint64 { _ = "STUB: not implemented"; return 0 }

func HashUint8(u uint8) uint64 { _ = "STUB: not implemented"; return 0 }

func HashInt64(i int64) uint64 { _ = "STUB: not implemented"; return 0 }

func HashInt32(i int32) uint64 { _ = "STUB: not implemented"; return 0 }

func HashInt16(i int16) uint64 { _ = "STUB: not implemented"; return 0 }

func HashInt8(i int8) uint64 { _ = "STUB: not implemented"; return 0 }

func HashInt(i int) uint64 { _ = "STUB: not implemented"; return 0 }

func HashUint(i uint) uint64 { _ = "STUB: not implemented"; return 0 }

func HashString(s string) uint64 { _ = "STUB: not implemented"; return 0 }

func HashBytes(b []byte) uint64 { _ = "STUB: not implemented"; return 0 }

func hash(u uint64) uint64 { _ = "STUB: not implemented"; return 0 }
