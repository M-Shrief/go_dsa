package queue

import "github.com/M-Shrief/go-dsa-practice/linkedlist"

type Queue[T any] struct {
	list linkedlist.Singly[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (queue *Queue[T]) Enqueue(data T) {
	queue.list.AddLast(data)
}

func (queue *Queue[T]) Dequeue() (T, bool) {
	return queue.list.DeleteFirst()
}

func (queue *Queue[T]) GetFirst() *linkedlist.SinglyNode[T] {
	return queue.list.GetHead()
}
