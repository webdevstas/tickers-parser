package updater

import (
	"sync"
)

type Queue[T comparable] struct {
	queue       []T
	unconfirmed []T
	mutex       sync.Mutex
}

func NewQueue[T comparable](items []T) Queue[T] {
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
	q.unconfirmed = append(q.unconfirmed, item)

	return item, true
}

func (q *Queue[T]) Confirm(item T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.queue = append(q.queue, item)
	for i, el := range q.unconfirmed {
		if el == item {
			q.unconfirmed = append(q.unconfirmed[:i], q.unconfirmed[i+1:]...)
		}
	}
}

func (q *Queue[T]) RestoreUnconfirmed() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.queue = append(q.queue, q.unconfirmed...)
	q.unconfirmed = nil
}
