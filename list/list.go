// Package list provides an implementation of a doubly-linked list with a front
// and back. The individual nodes of the list are publicly exposed so that the
// user can have fine-grained control over the list.
package list

// List implements a doubly-linked list.
type List[V any] struct {
	Front, Back *Node[V]
}

// Node is a node in the linked list.
type Node[V any] struct {
	Value      V
	Prev, Next *Node[V]
}

// New returns an empty linked list.
func New[V any]() *List[V] {
	_ = "STUB: not implemented"

	// PushBack adds 'v' to the end of the list.
	return nil
}

func (l *List[V]) PushBack(v V) { _ = "STUB: not implemented"; return }

// PushFront adds 'v' to the beginning of the list.
func (l *List[V]) PushFront(v V) { _ = "STUB: not implemented"; return }

// PushBackNode adds the node 'n' to the back of the list.
func (l *List[V]) PushBackNode(n *Node[V]) { _ = "STUB: not implemented"; return }

// PushFrontNode adds the node 'n' to the front of the list.
func (l *List[V]) PushFrontNode(n *Node[V]) { _ = "STUB: not implemented"; return }

// InsertAfter adds 'next' into the list after 'n'. Returns the added node.
func (l *List[V]) InsertAfter(n *Node[V], next *Node[V]) *Node[V] {
	_ = "STUB: not implemented"
	return nil
}

// InsertBefore adds 'prev' into the list before 'n'. Returns the added node.
func (l *List[V]) InsertBefore(n *Node[V], prev *Node[V]) *Node[V] {
	_ = "STUB: not implemented"
	return nil
}

// Remove removes the node 'n' from the list.
func (l *List[V]) Remove(n *Node[V]) { _ = "STUB: not implemented"; return }

// Each calls 'fn' on every element from this node onward in the list.
func (n *Node[V]) Each(fn func(val V)) { _ = "STUB: not implemented"; return }

// EachReverse calls 'fn' on every element from this node backward in the list.
func (n *Node[V]) EachReverse(fn func(val V)) { _ = "STUB: not implemented"; return }

// EachNode calls 'fn' on every node from this node onward in the list.
func (n *Node[V]) EachNode(fn func(n *Node[V])) { _ = "STUB: not implemented"; return }

// EachReverseNode calls 'fn' on every node from this node backward in the list.
func (n *Node[V]) EachReverseNode(fn func(n *Node[V])) { _ = "STUB: not implemented"; return }
