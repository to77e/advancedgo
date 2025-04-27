package hw04

type OrderedMap struct {
	key   int
	value int
	size  *int
	left  *OrderedMap
	right *OrderedMap
}

// NewOrderedMap - создать новый упорядоченный словарь
func NewOrderedMap() OrderedMap {
	return OrderedMap{
		size: new(int),
	}
}

// Insert - добавить элемент в словарь
func (m *OrderedMap) Insert(key, value int) {
	if *m.size == 0 {
		*m.size++
		m.key, m.value = key, value
		return
	}
	if key > m.key {
		if m.right == nil {
			*m.size++
			m.right = &OrderedMap{key: key, value: value, size: m.size}
		} else {
			m.right.Insert(key, value)
		}
	}
	if key < m.key {
		if m.left == nil {
			*m.size++
			m.left = &OrderedMap{key: key, value: value, size: m.size}
		} else {
			m.left.Insert(key, value)
		}
	}
}

// Erase - удалить элемент из словари
func (m *OrderedMap) Erase(key int) {
	if *m.size == 0 || !m.Contains(key) {
		return
	}

	if *m.size == 1 && m.key == key {
		*m.size--
		m.key, m.value = 0, 0
		return
	}

	tempMap := NewOrderedMap()
	m.ForEach(func(k, v int) {
		if k != key {
			tempMap.Insert(k, v)
		}
	})

	*m = tempMap
}

// Contains - проверить существование элемента в словаре
func (m *OrderedMap) Contains(key int) bool {
	if key == m.key && key != 0 {
		return true
	}
	if key > m.key {
		if m.right == nil {
			return false
		} else {
			return m.right.Contains(key)
		}
	}
	if key < m.key {
		if m.left == nil {
			return false
		} else {
			return m.left.Contains(key)
		}
	}
	return false
}

// Size - получить количество элементов в словаре
func (m *OrderedMap) Size() int {
	return *m.size
}

// ForEach - применить функцию к каждому элементу словаря от меньшего к большему
func (m *OrderedMap) ForEach(action func(int, int)) {
	if m.left != nil {
		m.left.ForEach(action)
	}
	action(m.key, m.value)
	if m.right != nil {
		m.right.ForEach(action)
	}
}
