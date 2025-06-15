package hw11

import "unsafe"

func Trace(stacks [][]uintptr) []uintptr {
	seen := make(map[uintptr]bool)
	var result []uintptr

	var dfs func(addr uintptr)
	dfs = func(addr uintptr) {
		next := *(*uintptr)(unsafe.Pointer(addr))
		if next == 0 {
			return
		}
		if seen[next] {
			return
		}
		result = append(result, next)
		seen[next] = true
		dfs(next)
	}

	for _, stack := range stacks {
		for _, u := range stack {
			if u == 0 {
				continue
			}
			if !seen[u] {
				result = append(result, u)
				seen[u] = true
				dfs(u)
			}
		}
	}
	return result
}
