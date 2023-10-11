package main

import (
	"fmt"
)

func main() {
	valid := "()[]{}"
	invalid := ")()]"
	invalid2 := "(]{})"
	fmt.Println(isValid(valid))    // return true
	fmt.Println(isValid(invalid))  // return false
	fmt.Println(isValid(invalid2)) // return false
}

var (
	startingParenthesis = 40
	startingBrackets    = 123
	startingSqBrackets  = 91

	closingParenthesis        = 41
	closingStartingBrackets   = 93
	closingStartingSqBrackets = 125
)

func isValid(input string) bool {
	// ( - { - [
	firstChar := input[0]
	fmt.Println(firstChar, byte(startingParenthesis))
	if firstChar != byte(closingParenthesis) || byte(firstChar) != byte(closingStartingBrackets) || firstChar != byte(closingStartingSqBrackets) {
		return false
	}
	// if we reach here, means that the first char is an open/opening char
	for i := 0; i < len(input); i++ {
		fmt.Println()
		openingSign := input[i]
		switch openingSign {
		case byte(startingParenthesis):
			//next char needs to be closingPar
			if input[i+1] != byte(closingParenthesis) {
				return false
			}
		case byte(startingBrackets):
			if input[i+1] != byte(closingStartingBrackets) {
				return false
			}
		case byte(startingSqBrackets):
			if input[i+1] != byte(closingStartingSqBrackets) {
				return false
			}
		}
		// if openingSign != rune(startingParenthesis) || openingSign != rune(startingBrackets) || openingSign != rune(startingSqBrackets) {
		// 	return false
		// }
		// count succesfull open and closing chars.
	}

	return true
	// for _, char := range input {
	// 	// if I can say

	// 	switch char {
	// 	case rune(startingParenthesis):
	// 		//next char needs to be closingPar
	// 		continue
	// 	case rune(startingBrackets):
	// 		//
	// 	case rune(startingSqBrackets):

	// 	}
	// 	if char != rune(startingParenthesis) || char != rune(startingBrackets) || char != rune(startingSqBrackets) {
	// 		return false
	// 	}
	// 	// count succesfull open and closing chars.
	// }

}
