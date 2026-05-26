// Package trie provides an implementation of a ternary search trie.
package trie

// Adapted from the TST implementation in Algorithms, 4th ed., by Robert
// Sedgewick and Kevin Wayne.
// https://algs4.cs.princeton.edu/52trie/TST.java.html.

// A Trie is a data structure that supports common prefix operations.
type Trie[V any] struct {
	n    int
	root *node[V]
}

type node[V any] struct {
	c                byte
	left, mid, right *node[V]
	val              V
	valid            bool
}

func (n *node[V]) isUnused() bool { _ = "STUB: not implemented"; return false }

func (n *node[V]) delete() *node[V] { _ = "STUB: not implemented"; return nil }

func (n *node[V]) minChild() *node[V] { _ = "STUB: not implemented"; return nil }

func (n *node[V]) deleteMinChild() *node[V] { _ = "STUB: not implemented"; return nil }

// n is the min node

// New returns an empty trie.
func New[V any]() *Trie[V] {
	_ = "STUB: not implemented"

	// Size returns the size of the trie.
	return nil
}

func (t *Trie[V]) Size() int {
	_ = "STUB: not implemented"

	// Contains returns whether this trie contains 'key'.
	return 0
}

func (t *Trie[V]) Contains(key string) bool { _ = "STUB: not implemented"; return false }

// Get returns the value associated with 'key'.
func (t *Trie[V]) Get(key string) (v V, ok bool) { _ = "STUB: not implemented"; return *new(V), false }

func (t *Trie[V]) get(x *node[V], key string, d int) *node[V] {
	_ = "STUB: not implemented"
	return nil
}

// Put associates 'val' with 'key'.
func (t *Trie[V]) Put(key string, val V) { _ = "STUB: not implemented"; return }

func (t *Trie[V]) put(x *node[V], key string, val V, d int) *node[V] {
	_ = "STUB: not implemented"
	return nil
}

// Remove removes the value associated with 'key', along with any nodes of the key that are no
// longer used.
func (t *Trie[V]) Remove(key string) { _ = "STUB: not implemented"; return }

func (t *Trie[V]) remove(x *node[V], key string, d int) *node[V] {
	_ = "STUB: not implemented"
	return nil
}

// LongestPrefix returns the key that is the longest prefix of 'query'.
func (t *Trie[V]) LongestPrefix(query string) string { _ = "STUB: not implemented"; return "" }

// Keys returns all keys in the trie.
func (t *Trie[V]) Keys() (queue []string) { _ = "STUB: not implemented"; return nil }

// KeysWithPrefix returns all keys with prefix 'prefix'.
func (t *Trie[V]) KeysWithPrefix(prefix string) (queue []string) {
	_ = "STUB: not implemented"
	return nil
}

func (t *Trie[V]) collect(x *node[V], prefix []byte, queue []string) []string {
	_ = "STUB: not implemented"
	return nil
}
