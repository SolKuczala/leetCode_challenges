package main

import "fmt"

//Given an array of integers, find if the array contains any duplicates.
//Your function should return true if any value appears at least twice in the array,
// and it should return false if every element is distinct.

func main() {
	a := []int{0, 2, 1, 3, 4, 5, 6}
	fmt.Println(containsDuplicate(a))
}

func containsDuplicate(nums []int) bool {
	//sort and compare
	numbers := map[int]string{}
	for _, n := range nums {
		if _, duplicate := numbers[n]; duplicate {
			return duplicate
		}
		numbers[n] = "pepito"
	}
	return false
}

//passed
