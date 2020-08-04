//URLify: Write a method to replace all spaces in a string with '%20'.
//You may assume that the string
//has sufficient space at the end to hold the additional characters,
// and that you are given the "true"
//length of the string.
//package main
//
//import (
//	"fmt"
//	"strings"
//	"unicode/utf8"
//)
//
//func main() {
//	fmt.Printf(uRLify("katarina pepo"))
//}
//
//func uRLify(str string) (URLifyString string) {
//	//este string tiene espacios, por ende itero y donde encuentro espacio o
//	//non character, meto en un string nuevo.
//	newArray := make([]rune, 0)
//	percentageRune, _ := utf8.DecodeRuneInString("%")
//	twoRune, _ := utf8.DecodeRuneInString("2")
//	zeroRune, _ := utf8.DecodeRuneInString("0")
//
//	for i, rune := range str {
//
//		fmt.Printf("%v %c \n", i, rune)
//		spaceRune, _ := utf8.DecodeRuneInString(" ")
//		if rune == spaceRune {
//			fmt.Printf("ESPACIO AT INDEX %d\n", i)
//			newArray = append(newArray, percentageRune)
//			newArray = append(newArray, twoRune)
//			newArray = append(newArray, zeroRune)
//		} else {
//			newArray = append(newArray, rune)
//		}
//
//		//newArray = append(newArray, "%20")
//
//		//newArray = append(newArray, string(rune))
//	}
//	return strings.Join(newArray, "2")
//}
//