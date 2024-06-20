package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	copyHead := head
	fmt.Println("first, ", copyHead)

	count := 0
	for head != nil {
		head = head.Next
		count++
	}
	i := 0
	for i < count/2+1 {
		fmt.Println("while loop: ", copyHead)
		copyHead = copyHead.Next
		i++
	}

	return copyHead
}

func main() {
	fmt.Println(middleNode([]{1,2,3,4}))
}
