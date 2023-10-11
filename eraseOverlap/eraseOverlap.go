package main

import "sort"

func main() {
	a := [][]int{
		{1, 2},
		{-1, 1},
		{3, 6},
		{1, 3},
		{2, 5},
	}
	b := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}
	c := [][]int{
		{2, 3},
		{5, 6},
		{10, 14},
		{1, 16},
	}
	print(eraseOverlapIntervals(a), "\n")
	print(eraseOverlapIntervals(b), "\n")
	print(eraseOverlapIntervals(c))
}
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 1 {
		return 0
	}
	//order slice
	sort.Ints()
	var toErase int
	baseInterval := intervals[0] //[n,n]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < baseInterval[1] || intervals[i][1] < baseInterval[0] {
			toErase++

		}
	}
	return toErase
}
