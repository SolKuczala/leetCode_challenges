package main

//Given a non-empty array of integers nums, every element appears twice except for one.
//Find that single one.
//Follow up: Could you implement a solution with a linear runtime complexity and without
//using extra memory?

func main() {
	a := []int{4, 1, 2, 1, 2} //4
	print(singleNumber(a))
}

func singleNumber(nums []int) int {
	numbers := map[int]bool{} //key:value
	for _, n := range nums {
		//esta duplicado?
		if _, dup := numbers[n]; !dup {
			//no
			numbers[n] = true
		} else {
			//si esta duplicado
			//delete mapname:key
			delete(numbers, n)
		}
	}
	var result int
	for key := range numbers {
		result = key
	}
	return result
}
