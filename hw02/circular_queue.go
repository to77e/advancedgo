package hw02

type CircularQueue struct {
	values []int
	size   int
	front  int
	rear   int
	used   int
}

// NewCircularQueue - создать очередь с определенным размером буффера
func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		values: make([]int, size),
		size:   size,
		used:   0,
		front:  0,
		rear:   0,
	}
}

// Push - добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	q.values[q.rear] = value
	q.rear = (q.rear + 1) % q.size
	q.used++

	return true
}

// Pop - удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	q.front = (q.front + 1) % q.size
	q.used--

	return true
}

// Front - получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.front]
}

// Back - получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	return q.values[(q.rear-1+q.size)%q.size]
}

// Empty - проверить пустая ли очередь
func (q *CircularQueue) Empty() bool {
	return q.used <= 0
}

// Full - проверить заполнена ли очередь
func (q *CircularQueue) Full() bool {
	return q.used >= q.size
}
