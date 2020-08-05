package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf(uRLify("katarina pepo"))
}
func uRLify(str string) (URLifyString string) {
	newArray := make([]rune, 0)
	percentageRune, _ := utf8.DecodeRuneInString("%")
	twoRune, _ := utf8.DecodeRuneInString("2")
	zeroRune, _ := utf8.DecodeRuneInString("0")
	//fmt.Printf("%d %d %d\n",percentageRune, twoRune,zeroRune) funciona
	for i, rune := range str {
		//fmt.Printf("%v %c \n", i, rune)funciona
		spaceRune, _ := utf8.DecodeRuneInString(" ")
		fmt.Printf("\n %d",spaceRune)
		if rune == spaceRune {
			fmt.Printf("ESPACIO AT INDEX %d\n", i)
			newArray = append(newArray, percentageRune)
			newArray = append(newArray, twoRune)
			newArray = append(newArray, zeroRune)
		} else {
			newArray = append(newArray, rune)
		}
	}
	fmt.Printf("%d \n",newArray)
	return string(newArray)
}
