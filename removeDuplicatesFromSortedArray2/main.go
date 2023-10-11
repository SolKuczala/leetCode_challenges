package main

import "fmt"

func main() {
	example := []int{1, 1, 1, 2, 2, 3}

	fmt.Println(removeDuplicates(example))

}

func removeDuplicates(nums []int) int {

	var p1 int
	var p2 int
	count := 1

	if nums[0] == nums[1] {
		p1 = 1
		p2 = 2

	} else {
		p1 = 2
		p2 = 3
	}

	// i will be our number to be replaced
	for i := p1; i < len(nums); i++ {
		if nums[p2] == nums[i] {
			// search for the next diff one to replace it
			for nums[p2] == nums[i] {
				p2++
			}
			// once it finds it replace p1
			nums[i] = nums[p2]
			count++
		} else {
			toi
		}
	}

}
