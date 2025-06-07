package hw12

type Task struct {
	Identifier int
	Priority   int
}

type PriorityHeap struct {
	tasks      []Task
	idToIndex  map[int]int
	priorities map[int]int
}

type Scheduler struct {
	heap PriorityHeap
}

func NewScheduler() Scheduler {
	return Scheduler{
		heap: PriorityHeap{
			tasks:      []Task{},
			idToIndex:  make(map[int]int),
			priorities: make(map[int]int),
		},
	}
}

func (s *Scheduler) AddTask(task Task) {
	h := &s.heap
	h.tasks = append(h.tasks, task)
	index := len(h.tasks) - 1
	h.idToIndex[task.Identifier] = index
	h.priorities[task.Identifier] = task.Priority

	h.siftUp(index)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	h := &s.heap
	index, exists := h.idToIndex[taskID]
	if !exists {
		return
	}

	oldPriority := h.priorities[taskID]
	h.priorities[taskID] = newPriority

	if newPriority > oldPriority {
		h.siftUp(index)
	} else if newPriority < oldPriority {
		h.siftDown(index)
	}
}

func (s *Scheduler) GetTask() Task {
	h := &s.heap
	if len(h.tasks) == 0 {
		return Task{}
	}

	task := h.tasks[0]

	lastIndex := len(h.tasks) - 1
	h.tasks[0] = h.tasks[lastIndex]
	h.tasks = h.tasks[:lastIndex]

	delete(h.idToIndex, task.Identifier)
	delete(h.priorities, task.Identifier)

	if lastIndex > 0 {
		h.idToIndex[h.tasks[0].Identifier] = 0
		h.siftDown(0)
	}

	return task
}

func (h *PriorityHeap) siftUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.getPriority(h.tasks[parentIndex].Identifier) >= h.getPriority(h.tasks[index].Identifier) {
			break
		}
		h.swap(index, parentIndex)
		index = parentIndex
	}
}

func (h *PriorityHeap) siftDown(index int) {
	lastIndex := len(h.tasks) - 1
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		largestIndex := index

		if leftChildIndex <= lastIndex &&
			h.getPriority(h.tasks[leftChildIndex].Identifier) > h.getPriority(h.tasks[largestIndex].Identifier) {
			largestIndex = leftChildIndex
		}
		if rightChildIndex <= lastIndex &&
			h.getPriority(h.tasks[rightChildIndex].Identifier) > h.getPriority(h.tasks[largestIndex].Identifier) {
			largestIndex = rightChildIndex
		}

		if largestIndex == index {
			break
		}

		h.swap(index, largestIndex)
		index = largestIndex
	}
}

func (h *PriorityHeap) getPriority(taskID int) int {
	return h.priorities[taskID]
}

func (h *PriorityHeap) swap(i, j int) {
	h.idToIndex[h.tasks[i].Identifier] = j
	h.idToIndex[h.tasks[j].Identifier] = i

	h.tasks[i], h.tasks[j] = h.tasks[j], h.tasks[i]
}
