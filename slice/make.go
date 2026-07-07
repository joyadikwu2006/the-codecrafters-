package main

func BuildAndGrow(length int, capacity int, extra int) []int {
	start := make([]int, length, capacity)
	start = append(start, extra)
	return start
}
