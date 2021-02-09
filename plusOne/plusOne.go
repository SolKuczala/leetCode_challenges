package main

func main() {

	a := []int{1, 2, 3}
	//b := []int{4, 3, 2, 1}
	//c := []int{0}

	print(plusOne(a))
	//print(plusOne(b))
	//print(plusOne(c))
}

//return []int
func plusOne(digits []int) []int {
	decena := len(digits) //3
	mynum := 0
	var finalNum []int
	for i := 0; i < len(digits); i++ {
		mynum += digits[i] * (10 ^ decena - 1) //300
		decena--
		//agarrar el numero y pushearlo sin los ceros
	}
	//finalNum:=mynum+1

	return []int{}
}
