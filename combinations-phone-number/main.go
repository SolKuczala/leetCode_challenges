package main

import (
	"fmt"
	"strings"
)

var phoneNumbersMap = map[string]string{
	"2": "a,b,c",
	"3": "d,e,f",
	"4": "g,h,i",
	"5": "j,l,k",
	"6": "m,n,o",
	"7": "p,q,r,s",
	"8": "t,u,v",
	"9": "w,x,y,z",
}

func main() {
	input := "2"

	fmt.Println(letterCombination(input))
}

// receives a lenght of 4 numbers
func letterCombination(digits string) []string {
	var result []string
	if digits == "" {
		return result
	}

	var stringValues []string
	// extract values from digits
	for _, n := range digits {
		num := string(n)

		stringValues = append(stringValues, phoneNumbersMap[num])
	}

	if len(stringValues) == 1 {
		valuesInSlice := strings.Split(stringValues[0], ",")

		for _, v := range valuesInSlice {
			result = append(result, v)
		}
		return result

	}
	result = allCombinations(stringValues)
	return result
}

func allCombinations(numberValues []string) []string {
	var combination []string

	generateCombination(numberValues, 0, "", &combination)

	return combination
}

// example de 3
func generateCombination(input []string, index int, current string, result *[]string) {
	if index == len(input) {
		fmt.Println(current)
		*result = append(*result, current)
		return
	}

	chars := strings.Split(input[index], ",")
	for _, char := range chars { //first "a,b,c"
		newCurrent := current + char                            // apenda primero "a" // como primer new current, segundo "ad", tercer "adg"
		generateCombination(input, index+1, newCurrent, result) // "todo", 1, "a", array
	}
}

// abc, def
// abc - a
// def - d
//final: "ad"
//

// pqrs, wxyz, tuv
/*
pwt
qwt
rwt
swt

*/
/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

example:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

2 : abc
3 : def
4 : ghi
5 : jlk
6 : mno
7 : pqrs
8 : tuv
9 : wxyz



*/
