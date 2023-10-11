package main

import "sort"

func main() {
	slice := []int{2, 3, 4, 5, 5, 6, 0, 7, 6, 8, 5}
	val := 5

	println(removeElement(slice, val))

}

func removeElement(nums []int, val int) int {

	var k int
	slice := len(nums)
	for i := 0; i < slice; i++ {
		if nums[i] == val {
			nums[i] = 0
		} else {
			k++
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	return k

}
