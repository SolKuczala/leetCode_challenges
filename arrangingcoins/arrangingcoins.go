package main

import "fmt"

func main() {
	for i := 0; i < 30; i++ {
		fmt.Printf("%d:%d \n", arrangeCoins(i), i)
	}
}

func arrangeCoins(n int) int {
	if n < 1 {
		return 0
	}
	step := 1
	for n-step > step {
		n = n - step
		step++
	}
	return step
}
