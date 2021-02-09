package main

import "fmt"

//Say you have an array prices for which the ith element is the price of a given stock on day i.
//Design an algorithm to find the maximum profit. You may complete as many transactions
//as you like (i.e., buy one and sell one share of the stock multiple times).
func main() {
	//a := []int{7, 1, 5, 3, 6, 4} //7
	//b := []int{1, 2, 3, 4, 5}    //4
	//c := []int{7, 6, 4, 3, 1}    //0
	d := []int{6, 1, 3, 2, 4, 7} //7
	//fmt.Println(maxProfit(a))
	//fmt.Println(maxProfit(b))
	//fmt.Println(maxProfit(c))
	fmt.Println(maxProfit(d))
}

func maxProfit(prices []int) int {
	//con un while, contando peaks y valleys
}
