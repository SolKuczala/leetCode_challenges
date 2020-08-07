package main

import(
	"fmt"
)

func main() {
	a:= []int{1,2,3,4,5}
	fmt.Println(runningSum(a))
}

func runningSum(nums []int) []int {
	if len(nums) == 0{
		return nums
	}
	var newSlice []int
	temp:= 0
	for i := 0; i < len(nums); i++ {
	  newSlice = append(newSlice, nums[i] + temp)
	  temp = newSlice[i] 
	}
	return newSlice
}