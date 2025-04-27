package hw02

type Signed interface {
	int | int8 | int16 | int32 | int64
}

type CircularQueueGeneric[T Signed] struct {
	values []T
	size   int
	front  int
	rear   int
	used   int
}

// NewCircularQueueGeneric - создать очередь с определенным размером буффера
func NewCircularQueueGeneric[T Signed](size int) CircularQueueGeneric[T] {
	return CircularQueueGeneric[T]{
		values: make([]T, size),
		size:   size,
	}
}

// Push - добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueueGeneric[T]) Push(value T) bool {
	if q.Full() {
		return false
	}

	q.values[q.rear] = value
	q.rear = (q.rear + 1) % q.size
	q.used++

	return true
}

// Pop - удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueueGeneric[T]) Pop() bool {
	if q.Empty() {
		return false
	}

	q.front = (q.front + 1) % q.size
	q.used--

	return true
}

// Front - получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueueGeneric[T]) Front() T {
	if q.Empty() {
		return -1
	}
	return q.values[q.front]
}

// Back - получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueueGeneric[T]) Back() T {
	if q.Empty() {
		return -1
	}
	return q.values[(q.rear-1+q.size)%q.size]
}

// Empty - проверить пустая ли очередь
func (q *CircularQueueGeneric[T]) Empty() bool {
	return q.used <= 0
}

// Full - проверить заполнена ли очередь
func (q *CircularQueueGeneric[T]) Full() bool {
	return q.used >= q.size
}
