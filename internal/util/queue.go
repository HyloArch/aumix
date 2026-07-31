package util

import "fmt"

type Queue[T any] struct {
	elements []T
	head     int
	size     int
}

func NewQueueWithCapacity[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		head:     0,
		size:     0,
		elements: make([]T, capacity),
	}
}

func (q *Queue[T]) Enqueue(value T) bool {
	if cap(q.elements) == q.size {
		return false
	}

	tail := (q.head + q.size) % cap(q.elements)
	q.elements[tail] = value
	q.size++
	return true
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var t T
	if q.size == 0 {
		return t, false
	}

	value := q.elements[q.head]
	q.elements[q.head] = t
	q.head = (q.head + 1) % cap(q.elements)
	q.size--
	return value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	var t T
	if q.size == 0 {
		return t, false
	}
	return q.elements[q.head], true
}

func (q *Queue[T]) Clear() {
	var t T
	for index := range cap(q.elements) {
		q.elements[index] = t
	}
	q.head = 0
	q.size = 0
}

func (q *Queue[T]) Len() int {
	return q.size
}

func (q *Queue[T]) String() string {
	return fmt.Sprintf("%v (%d/%d)", q.elements, q.size, cap(q.elements))
}
