package main

import (
	"fmt"
)

func main() {
	a := []int{4, 3, 2, 7, 8, 2, 3, 4, 1}
	fmt.Println(findDuplicates(a))
}

func findDuplicates(nums []int) []int {
	encountered := map[int]bool{}
	duplicates := []int{}

	for i := range nums {
		if encountered[nums[i]] != true {
			encountered[nums[i]] = true
		} else {
			duplicates = append(duplicates, nums[i])
		}
	}
	return duplicates
}
