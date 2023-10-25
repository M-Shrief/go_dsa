package queue

import "github.com/M-Shrief/go-dsa-practice/linkedlist"

type Queue[T any] struct {
	list linkedlist.Singly[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (queue *Queue[T]) Enqueue(data T) {
	queue.list.AddLast(data)
	queue.size++
}

func (queue *Queue[T]) Dequeue() (T, bool) {
	val, ok := queue.list.DeleteFirst()
	if ok {
		queue.size--
	}

	return val, ok
}

func (queue *Queue[T]) GetFirst() T {
	return queue.list.GetHead().GetVal()
}

func (queue *Queue[T]) GetSize() int {
	return queue.size
}
