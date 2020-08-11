package main

import (
	"fmt"
	"math"
	"math/big"
)

//Given two non-negative integers num1 and num2
//represented as strings, return the product of num1 and num2, also represented as a string.
//The length of both num1 and num2 is < 110.
//Both num1 and num2 contain only digits 0-9.
//Both num1 and num2 do not contain any leading zero, except the number 0 itself.
//You must not use any built-in BigInteger library or convert the inputs to integer directly.

func main() {
	fmt.Println(multiply("498828660196", "840477629533"))
}

//48-57 = 0-9
func multiply(num1 string, num2 string) string {
	zeroDecEquivalent := int64("0"[0]) //48
	accumulatorA := big.NewInt(0)
	accumulatorB := big.NewInt(0)
	//tengo que manejar todo con big int
	lenthArray := len(num1) - 1

	for i := 0; i < len(num1); i++ {
		ascii := int64(num1[i])
		accumulatorA += (ascii - zeroDecEquivalent) * int64(math.Pow(10, float64(lenthArray)))
		lenthArray--
	}
	lenthArray = len(num2) - 1

	for i := 0; i < len(num2); i++ {
		ascii := int64(num2[i])
		accumulatorB += (ascii - zeroDecEquivalent) * int64(math.Pow(10, float64(lenthArray)))
		lenthArray--
	}

	c := big.NewInt(0).Mul(big.NewInt(accumulatorA), big.NewInt(accumulatorB))

	return fmt.Sprintf("%d", c)
}
