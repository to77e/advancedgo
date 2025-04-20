package hw04

type OrderedMap struct {
	key   *int
	value *int
	left  *OrderedMap
	right *OrderedMap
}

// NewOrderedMap - создать новый упорядоченный словарь
func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

// Insert - добавить элемент в словарь
func (m *OrderedMap) Insert(key, value int) {
	if m == nil || m.key == nil || m.value == nil {
		m.key = &key
		m.value = &value
		return
	}

	curr := m
	var prev *OrderedMap
	for curr != nil {
		if key == *curr.key {
			*curr.value = value
			return
		}
		prev = curr
		if key < *curr.key {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	newMap := OrderedMap{
		key:   &key,
		value: &value,
	}
	if prev == nil {
		prev = &newMap
	} else if key < *prev.key {
		prev.left = &newMap
	} else {
		prev.right = &newMap
	}
}

// Erase - удалить элемент из словари
func (m *OrderedMap) Erase(key int) {
	if m == nil || m.key == nil || m.value == nil {
		return
	}
	curr := m
	var prev *OrderedMap
	for curr != nil && *curr.key != key {
		prev = curr
		if key < *curr.key {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	if curr == nil {
		return
	}

	if curr.right == nil {
		if prev == nil {
			*m = *curr.left
		} else if prev.left == curr {
			prev.left = curr.left
		} else {
			prev.right = curr.left
		}
		return
	}

	leftMost := curr.right
	var leftMostPrev *OrderedMap

	for leftMost.left != nil {
		leftMostPrev = leftMost
		leftMost = leftMost.left
	}

	curr.key = leftMost.key
	curr.value = leftMost.value

	if leftMostPrev != nil {
		leftMostPrev.left = leftMost.right
	} else {
		curr.right = leftMost.right
	}
}

// Contains - проверить существование элемента в словаре
func (m *OrderedMap) Contains(key int) bool {
	if m == nil || m.key == nil || m.value == nil {
		return false
	}
	curr := m
	for curr != nil {
		if key == *curr.key {
			return true
		}
		if key < *curr.key {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

// Size - получить количество элементов в словаре
func (m *OrderedMap) Size() int {
	if m == nil || m.key == nil || m.value == nil {
		return 0
	}
	return 1 + m.left.Size() + m.right.Size()
}

// ForEach - применить функцию к каждому элементу словаря от меньшего к большему
func (m *OrderedMap) ForEach(action func(int, int)) {
	if m == nil || m.key == nil || m.value == nil {
		return
	}
	m.left.ForEach(action)
	action(*m.key, *m.value)
	m.right.ForEach(action)
}
