package hashmap

import (
	"fmt"
	"hash/fnv"
)

type node[T any] struct {
	key   string
	value T
	next  *node[T]
}

type HashMap[T any] struct {
	capacity uint
	size     uint
	table    map[uint64]*node[T]
}

func New[T any](capacity uint) *HashMap[T] {
	return &HashMap[T]{
		capacity: capacity,
		size:     0,
		table:    make(map[uint64]*node[T], capacity),
	}
}

func (hm *HashMap[T]) Put(key string, value T) (hash uint64, err error) {
	if hm.size == hm.capacity {
		err = fmt.Errorf("hashmap is full")
		hash = 0
		return hash, err
	}
	hash = hm.hash(key)
	hmNode := hm.table[hash]
	if hmNode == nil {
		hm.table[hash] = &node[T]{
			key:   key,
			value: value,
			next:  nil,
		}
	} else {
		hm.table[hash] = &node[T]{
			key:   key,
			value: value,
			next:  hmNode,
		}
	}

	hm.size++
	return hash, nil
}

func (hm *HashMap[T]) Get(key string) (T, error) {
	if hm.size == 0 {
		var t T
		return t, fmt.Errorf("HashMap is Empty")
	}
	hash := hm.hash(key)
	return hm.getHelper(key, hash)
}

func (hm *HashMap[T]) getHelper(key string, hash uint64) (T, error) {
	node := hm.table[hash]
	if node == nil {
		var t T
		return t, fmt.Errorf("not found")
	} else if node.key == key {
		return node.value, nil
	} else {
		inList := hm.traverseNode(node.next, key)

		if inList == nil {
			var t T
			return t, fmt.Errorf("not found")
		} else {
			return inList.value, nil
		}
	}
}
func (hm *HashMap[T]) traverseNode(node *node[T], key string) *node[T] {
	current := node
	if current == nil {
		return nil
	} else if current.key == key {
		return current
	} else {
		return hm.traverseNode(current.next, key)
	}
}

func (hm *HashMap[T]) hash(key string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))
	return h.Sum64()
}
