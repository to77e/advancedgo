package hw05

func Map(data []int, action func(int) int) []int {
	if data == nil {
		return nil
	}
	if len(data) == 0 {
		return []int{}
	}
	newData := make([]int, 0, len(data))
	for _, v := range data {
		newData = append(newData, action(v))
	}
	return newData
}

func Filter(data []int, action func(int) bool) []int {
	if data == nil {
		return nil
	}
	if len(data) == 0 {
		return []int{}
	}
	newData := make([]int, 0, len(data))
	for _, v := range data {
		if ok := action(v); !ok {
			continue
		}
		newData = append(newData, v)
	}
	return newData
}

func Reduce(data []int, initial int, action func(int, int) int) int {
	if data == nil {
		return 0
	}
	result := initial
	for _, v := range data {
		result = action(v, result)
	}
	return result
}
