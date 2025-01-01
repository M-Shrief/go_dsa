package cache

import (
	"github.com/M-Shrief/go-dsa-practice/ds/heap"
	"github.com/M-Shrief/go-dsa-practice/utils"
)

type mfuNode[T any] struct {
	key  string
	val  T
	freq int
}

type MFU[T any] struct {
	// MFU heap
	maxHeap *heap.Heap[*mfuNode[T]]
	// MFU's size
	size int
	// MFU's capacity
	capacity int
	// MFU's fast access storage (hashmap)
	storage map[string]*mfuNode[T]
}

// Create a new MFU with a certain capacity.
func NewMFU[T any](capacity int) *MFU[T] {
	isMostFrequent := func(a, b *mfuNode[T]) bool {
		if utils.Compare(a.freq, b.freq) == 1 {
			return true
		} else {
			return false
		}
	}

	return &MFU[T]{
		maxHeap:  heap.NewHeap(isMostFrequent),
		storage:  make(map[string]*mfuNode[T], capacity),
		size:     0,
		capacity: capacity,
	}
}

func (mfu *MFU[T]) GetSize() int {
	return mfu.size
}

func (mfu MFU[T]) GetTop() *mfuNode[T] {
	return mfu.maxHeap.Top()
}

/*
Add a new item to MFU.
*/
func (mfu *MFU[T]) Put(key string, value T) {
	newNode := &mfuNode[T]{key: key, val: value, freq: 0}
	if mfu.size == mfu.capacity {
		mfu.evict()
	}
	mfu.maxHeap.Push(newNode)
	mfu.storage[key] = newNode
	mfu.size++
}

// Evict Top item from MFU's MinHEAP and Hashmap.
func (mfu *MFU[T]) evict() bool {
	if mfu.size == 0 {
		return false
	}
	top := mfu.maxHeap.Top()
	delete(mfu.storage, top.key)
	mfu.maxHeap.Pop()
	mfu.size--
	return true
}

// Get MFU's item by key
// Returns item's value, and boolean to report if it exists or not.
func (mfu *MFU[T]) Get(key string) (T, bool) {
	if mfu.size == 0 {
		var r T
		return r, false
	}
	node, ok := mfu.storage[key]
	if !ok {
		var r T
		return r, false
	}

	// 1- Delete the node
	// 2- Increase it's freq
	// 3- Push it again and the heap correct its position
	mfu.maxHeap.Delete(node)
	node.freq++
	mfu.maxHeap.Push(node)

	return node.val, true
}
