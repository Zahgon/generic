// Package prope provides an implementation of a persistent rope data structure.
// It is similar to the base rope data structure, but the changes
// are saved separately without modifying the original data structure by
// sharing data between multiple versions. The time complexity of operations
// stay the same, but they are generally a bit slower:
//
// * Remove: O(lg n).
//
// * Insert: O(lg n).
//
// * Slice: O(lg n + m), where m is the size of the slice.
//
// * At: O(lg n).
//
// The main difference is in space complexity, as the persistent data structure
// allows creating a copy with an insertion or removal in O(lg n) space,
// instead of duplicating the entire rope for each change.
// This also prevents the O(n) time complexity of cloning the rope
// to save a version, as this is done inside the operations in a more efficient
// manner.
package prope

var (
	// SplitLength is the threshold above which slices will be split into
	// separate nodes. Larger values will take make operations take more
	// memory.
	SplitLength = 256
	// JoinLength is the threshold below which nodes will be merged into
	// slices.
	JoinLength = SplitLength / 2
	// RebalanceRatio is the threshold used to trigger a rebuild during a
	// rebalance operation.
	RebalanceRatio = 1.5
)

type nodeType byte

const (
	tLeaf nodeType = iota
	tNode
)

// A Node in the rope structure. If the kind is tLeaf, only the value and
// length are valid, and if the kind is tNode, only length, left, right are
// valid.
type Node[V any] struct {
	kind        nodeType
	value       []V
	length      int
	left, right *Node[V]
}

// New returns a new rope node from the given byte slice. The underlying
// data is not copied so the user should ensure that the slice will
// not be modified after the rope is created.
func New[V any](b []V) *Node[V] { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements stored in the rope.
func (n *Node[V]) Len() int {
	_ = "STUB: not implemented"

	// Value returns the elements of this node concatenated into a slice.
	return 0
}

func (n *Node[V]) Value() []V { _ = "STUB: not implemented"; return nil }

// Slice returns the range of the rope from [start:end).
func (n *Node[V]) Slice(start, end int) []V { _ = "STUB: not implemented"; return nil }

// At returns the element at the given position.
func (n *Node[V]) At(pos int) V { _ = "STUB: not implemented"; return *new(V) }

// Insert returns a new version of the rope with the given
// value inserted at pos.
func (n *Node[V]) Insert(pos int, value []V) *Node[V] { _ = "STUB: not implemented"; return nil }

// Adjusting is done here

// Remove returns a new version of the rope with the elements
// in the [start:end) range removed.
func (n *Node[V]) Remove(start, end int) *Node[V] { _ = "STUB: not implemented"; return nil }

// SplitAt splits the node at the given index and returns two new ropes
// corresponding to the left and right portions of the split.
func (n *Node[V]) SplitAt(i int) (*Node[V], *Node[V]) { _ = "STUB: not implemented"; return nil, nil }

// Rebalance finds unbalanced nodes and rebuilds them.
// Rebuilded nodes does not share memory with their old versions,
// so sometimes this operation will take up a lot of memory.
func (n *Node[V]) Rebalance() { _ = "STUB: not implemented"; return }

// Rebuild rebuilds the entire rope structure, resulting in a balanced tree.
// The rebuilded node does not share memory with its old versions,
// so this operation will take the same space as creating the node from scratch.
func (n *Node[V]) Rebuild() { _ = "STUB: not implemented"; return }

// Join creates a merged version of all of the ropes.
func Join[V any](nodes ...*Node[V]) *Node[V] { _ = "STUB: not implemented"; return nil }

func (n *Node[V]) copy(dst []V) { _ = "STUB: not implemented"; return }

func (n *Node[V]) copySlice(dst []V, start, end int) { _ = "STUB: not implemented"; return }

func (n *Node[V]) adjust() { _ = "STUB: not implemented"; return }

// Bounds the start and end indices to a given length.
func bound(start, end int, length int) (newStart, newEnd int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Cannot modify slices, so no tricks are possible
func insert[V any](slice []V, k int, insertion []V) []V { _ = "STUB: not implemented"; return nil }

func remove[V any](slice []V, start, end int) []V { _ = "STUB: not implemented"; return nil }
