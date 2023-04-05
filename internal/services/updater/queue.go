package updater

import (
	"sync"
)

type Queue[T any] struct {
	queue []T
	mutex sync.Mutex
}

func NewQueue[T any](items []T) Queue[T] {
	return Queue[T]{
		queue: items,
		mutex: sync.Mutex{},
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.queue = append(q.queue, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.queue) == 0 {
		var res T
		return res, false
	}
	item := q.queue[0]
	q.queue = q.queue[1:]

	return item, true
}

func (q *Queue[T]) Confirm(item T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.queue = append(q.queue, item)
}
