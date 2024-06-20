package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	// check needle str is < than haystack str
	if len(needle) > len(haystack) {
		return -1
	}
	// keep track of the needle'
	index := 0
	if strings.Contains(haystack, needle) {
		for index, h := range haystack {
			for _, n := range needle {
				if n != h {
					continue
				}
			}
			return index
		}
	} else {
		return -1
	}

	return index
}

func main() {
	//strStr()

	fmt.Print(len("abvdf"))
}
