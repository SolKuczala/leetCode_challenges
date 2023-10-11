package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
You will not be given access to the first node of head.

All the values of the linked list are unique,

	and it is guaranteed that the given node node is not the last node in the linked list.
*/
func main() {
	//example
	//	head = [4,5,1,9], node = 1
	removeNthFromEnd(input)
	fmt.Print(input)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// encontrar el ultimo nodo
	dummy := &ListNode{Val: 0, Next: head}
	first := dummy
	second := dummy

	// Move the first pointer n nodes ahead
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// Move both pointers one node at a time
	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return dummy.Next

}
