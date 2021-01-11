package main

import "fmt"

//Say you have an array prices for which the ith element is the price of a given stock on day i.
//Design an algorithm to find the maximum profit. You may complete as many transactions
//as you like (i.e., buy one and sell one share of the stock multiple times).
func main() {
	a := []int{7, 1, 5, 3, 6, 4}
	b := []int{1, 2, 3, 4, 5}
	c := []int{7, 6, 4, 3, 1}

	fmt.Println(maxProfit(a))
	fmt.Println(maxProfit(b))
	fmt.Println(maxProfit(c))
}

func maxProfit(prices []int) int {
	//comparo los primero dos para ver cual es mas bajo para hacer la primer compra
	var base int
	//secElement := prices[1]
	var profit int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			base = prices[i]
			//cuenta- buscar cual es el mayor de los que vienen,
			//y si aun hay menores despues del encontrado, seguir haciendo el loop
			profit += prices[i+1] - base
			//se corre el puntero
			i++
		} else {
			base = prices[i+1]
		}
	}

	return profit
}
