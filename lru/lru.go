// Check this: https://www.geeksforgeeks.org/lru-cache-implementation/#lru-cache-implementation-using-queue-and-hashing
package lru

import (
	"github.com/M-Shrief/go-dsa-practice/linkedlist"
)

// LRU item
type item[T any] struct {
	key   string
	value T
}

// LRU cache
type LRU[T any] struct {
	// Doubly Linkedlist
	dl *linkedlist.Doubly[*item[T]]
	// LRU's size
	size int
	// LRU's capacity
	capacity int
	// LRU's fast access storage (hashmap)
	storage map[string]*linkedlist.DoublyNode[*item[T]]
}

// Create a new LRU with a certain capacity.
func NewLRU[T any](capacity int) *LRU[T] {
	return &LRU[T]{
		dl:       linkedlist.NewDoubly[*item[T]](),
		storage:  make(map[string]*linkedlist.DoublyNode[*item[T]], capacity),
		size:     0,
		capacity: capacity,
	}
}

/*
Add a new item to LRU.
*/
func (lru *LRU[T]) Put(key string, value T) {
	newNode := &item[T]{key, value}
	if lru.size == lru.capacity {
		lru.evict()
	}
	lru.dl.AddFirst(newNode)
	lru.storage[key] = lru.dl.GetHead()
	lru.size++
}

// Evict last item from Doubly Linkedlist and from LRU's storage.
func (lru *LRU[T]) evict() bool {
	if lru.size == 0 {
		return false
	}
	node, _ := lru.dl.DeleteLast()
	delete(lru.storage, node.key)
	lru.size--
	return true
}

// Get LRU's item by key
// Returns item's value, and boolean to report if it exists or not.
func (lru *LRU[T]) Get(key string) (T, bool) {
	if lru.size == 0 {
		var r T
		return r, false
	}
	node, ok := lru.storage[key]
	if !ok {
		var r T
		return r, false
	}
	lru.dl.DeleteByNode(node)
	lru.dl.AddFirst(node.GetVal())
	nodeVal := node.GetVal().value
	return nodeVal, true
}
