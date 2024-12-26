// Check this: https://www.geeksforgeeks.org/mru-cache-implementation/#mru-cache-implementation-using-queue-and-hashing
package cache

import (
	"github.com/M-Shrief/go-dsa-practice/ds/linkedlist"
)

// MRU cache
type MRU[T any] struct {
	// Doubly Linkedlist
	dl *linkedlist.Doubly[*item[T]]
	// MRU's size
	size int
	// MRU's capacity
	capacity int
	// MRU's fast access storage (hashmap)
	storage map[string]*linkedlist.DoublyNode[*item[T]]
}

// Create a new MRU with a certain capacity.
func NewMRU[T any](capacity int) *MRU[T] {
	return &MRU[T]{
		dl:       linkedlist.NewDoubly[*item[T]](),
		storage:  make(map[string]*linkedlist.DoublyNode[*item[T]], capacity),
		size:     0,
		capacity: capacity,
	}
}

/*
Add a new item to MRU.
*/
func (mru *MRU[T]) Put(key string, value T) {
	newNode := &item[T]{key, value}
	if mru.size == mru.capacity {
		mru.evict()
	}
	mru.dl.AddFirst(newNode)
	mru.storage[key] = mru.dl.GetHead()
	mru.size++
}

// Evict last item from Doubly Linkedlist and from MRU's storage.
func (mru *MRU[T]) evict() bool {
	if mru.size == 0 {
		return false
	}
	node, _ := mru.dl.DeleteLast()
	delete(mru.storage, node.key)
	mru.size--
	return true
}

// Get MRU's item by key
// Returns item's value, and boolean to report if it exists or not.
func (mru *MRU[T]) Get(key string) (T, bool) {
	if mru.size == 0 {
		var r T
		return r, false
	}
	node, ok := mru.storage[key]
	if !ok {
		var r T
		return r, false
	}
	mru.dl.DeleteByNode(node)
	mru.dl.AddLast(node.GetVal())
	mru.storage[key] = mru.dl.GetTail()
	nodeVal := node.GetVal().value
	return nodeVal, true
}
