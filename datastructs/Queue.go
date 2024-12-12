package datastructs

import "errors"

type Queue[T any] []T

var emptyError = errors.New("queue is empty")

func (q *Queue[T]) Enqueue(item T) {
	*q = append(*q, item)
}

func (q *Queue[T]) Dequeue() (val T, err error) {
	if len(*q) == 0 {
		err = emptyError
		return
	}

	val, *q = (*q)[0], (*q)[1:]
	return
}

func (q *Queue[T]) Front() (val T, err error) {
	if len(*q) == 0 {
		err = emptyError
		return
	}
	val = (*q)[0]
	return
}

func (q *Queue[T]) Size() int {
	return len(*q)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}
