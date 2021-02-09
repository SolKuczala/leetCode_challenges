package main

import "math"

//Given an array of positive integers nums and a positive integer target, return the minimal
//length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum
//is greater than or equal to target. If there is no such subarray, return 0 instead.
func main() {
	target2 := 7
	nums2 := []int{2, 3, 1, 2, 4, 3}
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	nums3 := []int{1, 2, 3, 4, 5}

	target4 := 4
	nums4 := []int{1, 4, 4}
	println(minSubArrayLen(target, nums))   //0
	println(minSubArrayLen(target, nums3))  //3
	println(minSubArrayLen(target2, nums2)) //2
	println(minSubArrayLen(target4, nums4)) //1
}

func minSubArrayLen(target int, nums []int) int {
	min := math.MaxInt64
	arrayLen := len(nums)
	found := false
	for i := 0; i < arrayLen; i++ {
		secondIndex := i + 1
		if nums[i] == target {
			return 1
		}
		var count int
		tempSum := nums[i]
		count++
		for tempSum < target && secondIndex < arrayLen {
			tempSum += nums[secondIndex]
			secondIndex++
			count++
		}
		if tempSum >= target && count < min {
			min = count
			found = true
		} else if !found {
			min = 0
		}
	}
	return min
}
