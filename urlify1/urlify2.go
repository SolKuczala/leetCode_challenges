package main

import (
	"fmt"
)

func main() {
	fmt.Printf(uRLify("D D"))
}
func uRLify(str string) (URLifyString string) {
	newArray := make([]rune, 0)
	for i := 0; i < len(str); i++ {
		spaceRune:= rune(' ')
		runa := rune(str[i])
		if runa == spaceRune {
			newArray = append(newArray, rune('%'))
			newArray = append(newArray, rune('2'))
			newArray = append(newArray, rune('0'))
			i = i + 2
		} else {
			newArray = append(newArray, runa)
		}
	}
	return string(newArray)
}
