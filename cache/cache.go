// Package cache provides an implementation of a key-value store with a maximum
// size. Once the maximum size is reached, the cache uses a least-recently-used
// policy to evict old entries. The cache is implemented as a combined hashmap
// and linked list. This ensures all operations are constant-time.
package cache

import (
	"github.com/zyedidia/generic/list"
)

// A Cache is an LRU cache for keys and values. Each entry is
// put into the table with an associated key used for looking up the entry.
// The cache has a maximum size, and uses a least-recently-used eviction
// policy when there is not space for a new entry.
type Cache[K comparable, V any] struct {
	capacity int
	lru      list.List[KV[K, V]]
	table    map[K]*list.Node[KV[K, V]]
	evictCb  func(key K, val V)
}

type KV[K comparable, V any] struct {
	Key K
	Val V
}

// New returns a new Cache with the given capacity.
func New[K comparable, V any](capacity int) *Cache[K, V] { _ = "STUB: not implemented"; return nil }

// Get returns the entry associated with a given key, and a boolean indicating
// whether the key exists in the table.
func (t *Cache[K, V]) Get(k K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

// Put adds a new key-entry pair to the table.
func (t *Cache[K, V]) Put(k K, e V) { _ = "STUB: not implemented"; return }

func (t *Cache[K, V]) evict() { _ = "STUB: not implemented"; return }

// Remove causes the entry associated with the given key to be immediately
// evicted from the cache.
func (t *Cache[K, V]) Remove(k K) { _ = "STUB: not implemented"; return }

// Resize changes the maximum capacity for this cache to 'capacity'.
func (t *Cache[K, V]) Resize(capacity int) { _ = "STUB: not implemented"; return }

// Size returns the number of active elements in the cache.
func (t *Cache[K, V]) Size() int { _ = "STUB: not implemented"; return 0 }

// Capacity returns the maximum capacity of the cache.
func (t *Cache[K, V]) Capacity() int {
	_ = "STUB: not implemented"

	// Each calls 'fn' on every value in the cache, from most recently used to
	// least recently used.
	return 0
}

func (t *Cache[K, V]) Each(fn func(key K, val V)) { _ = "STUB: not implemented"; return }

// SetEvictCallback sets a callback to be invoked before an entry is evicted.
// This replaces any prior callback set by this method.
func (t *Cache[K, V]) SetEvictCallback(fn func(key K, val V)) { _ = "STUB: not implemented"; return }
