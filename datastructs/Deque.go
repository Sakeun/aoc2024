package datastructs

import "errors"

type Deque[T any] []T

var EmptyDequeError = errors.New("Deque is empty")

func (d *Deque[T]) PopBack() (val T, err error) {
	size := d.Size()
	if size == 0 {
		err = EmptyDequeError
		return
	}

	val, *d = (*d)[size-1], (*d)[:size-1]
	return
}

func (d *Deque[T]) PopFront() (val T, err error) {
	if len(*d) == 0 {
		err = EmptyDequeError
		return
	}

	val, *d = (*d)[0], (*d)[1:]
	return
}

func (d *Deque[T]) PushBack(item T) {
	*d = append(*d, item)
}

func (d *Deque[T]) PushFront(item T) {
	*d = append([]T{item}, *d...)
}

func (d *Deque[T]) Size() int {
	return len(*d)
}

func (d *Deque[T]) PeekFront() T {
	return (*d)[0]
}

func (d *Deque[T]) PeekBack() T {
	var val T
	if d.Size() == 0 {
		return val
	}
	return (*d)[len(*d)-1]
}
