package main
import(
	"fmt"
)
func main() {
	a := []int{0,1,1,1,2,2,3,4,5,6}
	
	fmt.Println(removeDuplicates(a))
	fmt.Println(removeDuplicates2(a))
}

func removeDuplicates(nums []int) int{
	counter:= 1
    numberBefore:= nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != numberBefore{
			counter++
			numberBefore = nums[i]
		}
	}
	return counter
}

func removeDuplicates2(nums []int) int{
	if len(nums) == 0{
		return 0
	}
	numberBefore:= nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == numberBefore{
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}else{
			numberBefore = nums[i]
		}
	}
	return len(nums)
}