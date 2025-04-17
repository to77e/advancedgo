package hw02

import (
	"log/slog"
	"sync"
)

type Signed interface {
	int | int8 | int16 | int32 | int64
}

type CircularQueueGeneric[T Signed] struct {
	values []T
	mu     sync.Mutex
	size   int
	front  int
	rear   int
}

// NewCircularQueueGeneric - создать очередь с определенным размером буффера
func NewCircularQueueGeneric[T Signed](size int) CircularQueueGeneric[T] {
	return CircularQueueGeneric[T]{
		values: make([]T, size),
		mu:     sync.Mutex{},
		size:   size,
		front:  -1,
		rear:   -1,
	}
}

// Push - добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueueGeneric[T]) Push(value T) bool {
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
func (q *CircularQueueGeneric[T]) Pop() bool {
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
func (q *CircularQueueGeneric[T]) Front() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.Empty() {
		return -1
	}
	return q.values[q.front]
}

// Back - получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueueGeneric[T]) Back() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.Empty() {
		return -1
	}
	return q.values[q.rear]
}

// Empty - проверить пустая ли очередь
func (q *CircularQueueGeneric[T]) Empty() bool {
	if q.front == -1 {
		slog.Debug("the circular queue is empty")
		return true
	}
	return false
}

// Full - проверить заполнена ли очередь
func (q *CircularQueueGeneric[T]) Full() bool {
	if (q.rear+1)%q.size == q.front {
		slog.Debug("the circular queue is full")
		return true
	}
	return false
}
