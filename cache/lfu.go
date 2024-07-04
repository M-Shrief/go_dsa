package cache

import (
	"github.com/M-Shrief/go-dsa-practice/heap"
	"github.com/M-Shrief/go-dsa-practice/utils"
)

type lfuNode[T any] struct {
	key  string
	val  T
	freq int
}

type LFU[T any] struct {
	// LFU heap
	mh *heap.Heap[*lfuNode[T]]
	// LFU's size
	size int
	// LFU's capacity
	capacity int
	// LFU's fast access storage (hashmap)
	storage map[string]*lfuNode[T]
}

// Create a new LFU with a certain capacity.
func NewLFU[T any](capacity int) *LFU[T] {
	isSmaller := func(a, b *lfuNode[T]) bool {
		if utils.Compare(a.freq, b.freq) == -1 {
			return true
		} else {
			return false
		}
	}
	return &LFU[T]{
		mh:       heap.NewHeap(isSmaller),
		storage:  make(map[string]*lfuNode[T], capacity),
		size:     0,
		capacity: capacity,
	}
}

/*
Add a new item to LFU.
*/
func (lfu *LFU[T]) Put(key string, value T) {
	newNode := &lfuNode[T]{key: key, val: value, freq: 0}
	if lfu.size == lfu.capacity {
		lfu.evict()
	}
	lfu.mh.Push(newNode)
	lfu.storage[key] = newNode
	lfu.size++
}

// Evict Top item from LFU's MinHEAP and Hashmap.
func (lfu *LFU[T]) evict() bool {
	if lfu.size == 0 {
		return false
	}
	top := lfu.mh.Top()
	delete(lfu.storage, top.key)
	lfu.mh.Pop()
	lfu.size--
	return true
}

// Get LFU's item by key
// Returns item's value, and boolean to report if it exists or not.
func (lfu *LFU[T]) Get(key string) (T, bool) {
	if lfu.size == 0 {
		var r T
		return r, false
	}
	node, ok := lfu.storage[key]
	if !ok {
		var r T
		return r, false
	}

	// 1- Delete the node
	// 2- Increase it's freq
	// 3- Push it again and the heap correct its position
	lfu.mh.Delete(node)
	node.freq++
	lfu.mh.Push(node)

	return node.val, true
}
