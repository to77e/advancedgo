package hw12

type Task struct {
	Identifier int
	Priority   int
}

type Scheduler struct {
	tasks      []Task
	idToIndex  map[int]int
	priorities map[int]int
}

func NewScheduler() Scheduler {
	return Scheduler{
		tasks:      make([]Task, 0),
		idToIndex:  make(map[int]int),
		priorities: make(map[int]int),
	}
}

func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
	index := len(s.tasks) - 1
	s.idToIndex[task.Identifier] = index
	s.priorities[task.Identifier] = task.Priority

	s.siftUp(index)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	index, exists := s.idToIndex[taskID]
	if !exists {
		return
	}

	oldPriority := s.priorities[taskID]
	s.priorities[taskID] = newPriority

	if newPriority > oldPriority {
		s.siftUp(index)
	} else if newPriority < oldPriority {
		s.siftDown(index)
	}
}

func (s *Scheduler) GetTask() Task {
	if len(s.tasks) == 0 {
		return Task{}
	}

	task := s.tasks[0]

	lastIndex := len(s.tasks) - 1
	s.tasks[0] = s.tasks[lastIndex]
	s.tasks = s.tasks[:lastIndex]

	delete(s.idToIndex, task.Identifier)
	delete(s.priorities, task.Identifier)

	if lastIndex > 0 {
		s.idToIndex[s.tasks[0].Identifier] = 0
		s.siftDown(0)
	}

	return task
}

func (s *Scheduler) siftUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if s.getPriority(s.tasks[parentIndex].Identifier) >= s.getPriority(s.tasks[index].Identifier) {
			break
		}
		s.swap(index, parentIndex)
		index = parentIndex
	}
}

func (s *Scheduler) siftDown(index int) {
	lastIndex := len(s.tasks) - 1
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		largestIndex := index

		if leftChildIndex <= lastIndex &&
			s.getPriority(s.tasks[leftChildIndex].Identifier) > s.getPriority(s.tasks[largestIndex].Identifier) {
			largestIndex = leftChildIndex
		}
		if rightChildIndex <= lastIndex &&
			s.getPriority(s.tasks[rightChildIndex].Identifier) > s.getPriority(s.tasks[largestIndex].Identifier) {
			largestIndex = rightChildIndex
		}

		if largestIndex == index {
			break
		}

		s.swap(index, largestIndex)
		index = largestIndex
	}
}

func (s *Scheduler) getPriority(taskID int) int {
	return s.priorities[taskID]
}

func (s *Scheduler) swap(i, j int) {
	s.idToIndex[s.tasks[i].Identifier] = j
	s.idToIndex[s.tasks[j].Identifier] = i

	s.tasks[i], s.tasks[j] = s.tasks[j], s.tasks[i]
}
