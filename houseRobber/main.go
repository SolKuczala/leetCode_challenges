package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(rob([]int{2, 7, 9, 3, 1, 8, 6, 8, 4, 1, 38, 6, 8, 6, 86, 4, 354, 35}))

}

// vienen del 1 al 100 en largo y valores del 0 al 400
func rob(houses []int) int {
	// split the array
	sumOfPathOne := 0
	sumOfPathTwo := 0

	for i := 0; i < len(houses); i++ {
		// identify if the index is odd or even
		if i%2 == 0 {
			sumOfPathOne += houses[i]
			fmt.Println(sumOfPathOne)
		} else {
			sumOfPathTwo += houses[i]
			//fmt.Println(sumOfPathTwo)
		}
	}

	return int(math.Max(float64(sumOfPathOne), float64(sumOfPathTwo)))
}
