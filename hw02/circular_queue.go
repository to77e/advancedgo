package hw02

import (
	"log/slog"
	"sync"
)

type CircularQueue struct {
	values []int
	mu     sync.Mutex
	size   int
	front  int
	rear   int
}

// NewCircularQueue - создать очередь с определенным размером буффера
func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		mu:     sync.Mutex{},
		size:   size,
		front:  -1,
		rear:   -1,
	}
}

// Push - добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Push(value int) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.Full() {
		return false
	}

	if q.front == -1 {
		q.front, q.rear = 0, 0
	} else {
		q.rear = (q.rear + 1) % q.size
	}
	q.values[q.rear] = value

	return true
}

// Pop - удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue) Pop() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.Empty() {
		return false
	}

	if q.front == q.rear {
		q.front, q.rear = -1, -1
	} else {
		q.front = (q.front + 1) % q.size
	}

	return true
}

// Front - получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue) Front() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.Empty() {
		return -1
	}
	return q.values[q.front]
}

// Back - получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue) Back() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	
	if q.Empty() {
		return -1
	}
	return q.values[q.rear]
}

// Empty - проверить пустая ли очередь
func (q *CircularQueue) Empty() bool {
	if q.front == -1 {
		slog.Debug("the circular queue is empty")
		return true
	}
	return false
}

// Full - проверить заполнена ли очередь
func (q *CircularQueue) Full() bool {
	if (q.rear+1)%q.size == q.front {
		slog.Debug("the circular queue is full")
		return true
	}
	return false
}
