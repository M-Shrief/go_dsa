// Check this: https://www.geeksforgeeks.org/lru-cache-implementation/#lru-cache-implementation-using-queue-and-hashing
package lru

import (
	"github.com/M-Shrief/go-dsa-practice/linkedlist"
)

type item[T any] struct {
	key   string
	value T
}

type LRU[T any] struct {
	dl       *linkedlist.DoublyLinkedList[*item[T]]
	size     int
	capacity int
	storage  map[string]*linkedlist.DoublyNode[*item[T]]
}

func NewLRU[T any](capacity int) *LRU[T] {
	return &LRU[T]{
		dl:       linkedlist.NewDoublyLinkedList[*item[T]](),
		storage:  make(map[string]*linkedlist.DoublyNode[*item[T]], capacity),
		size:     0,
		capacity: capacity,
	}
}

/* Put

 */

func (lru *LRU[T]) Put(key string, value T) {
	newNode := &item[T]{key, value}
	if lru.size == lru.capacity {
		lru.evict()
		lru.dl.AddFirst(newNode)
		lru.storage[key] = lru.dl.GetHead()
		lru.size++
		return
	}
	lru.dl.AddFirst(newNode)
	lru.storage[key] = lru.dl.GetHead()
	lru.size++
}

func (lru *LRU[T]) evict() bool {
	if lru.size == 0 {
		return false
	}
	node, _ := lru.dl.DeleteLast()
	delete(lru.storage, node.key)
	lru.size--
	return true
}
