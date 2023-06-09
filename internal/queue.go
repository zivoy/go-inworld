package internal

import (
	"fmt"
	"sync"
)

// Queue is fifo
type Queue[T any] struct {
	size  uint
	queue *queueElm[T]
	last  *queueElm[T]

	mu sync.Mutex
}

// NewPointer is a helper for making pointers to stuff like numbers
func NewPointer[T any](v T) *T {
	return &v
}

type queueElm[T any] struct {
	next *queueElm[T]
	data *T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Append(data *T) {
	e := &queueElm[T]{
		data: data,
	}

	q.mu.Lock()
	defer q.mu.Unlock()
	q.size++
	if q.queue == nil {
		q.last = e
		q.queue = e
		return
	}

	q.last.next = e
	q.last = e
}

func (q *Queue[T]) Pop() *T {
	if q.queue == nil {
		return nil
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	e := q.queue
	q.queue = q.queue.next
	if q.queue == nil {
		q.last = nil
	}
	q.size--

	return e.data
}

func (q *Queue[T]) Clear() {
	q.mu.Lock()
	q.queue = nil
	q.last = nil
	q.size = 0
	q.mu.Unlock()
}

// Size returns an int rather than an uint for convenience
func (q *Queue[T]) Size() int {
	return int(q.size)
}

func (q *Queue[T]) String() string {
	return fmt.Sprintf("Queue of %T with %d elements", *new(T), q.size)
}
