package main

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	aux := &ListNode{}
	result := aux

	// add a check for 0 or one nodes
	if head == nil || head.Next == nil {
		return head
	}

	for head.Next != nil {
		println("about to set first node")

		aux = head.Next //el segundo
		aux.Next = head
		println("aux:", aux.Val, "aux.next.val", aux.Next.Val)
		//result.val es un 2
		// mover aux hasta el proximo nil(3r posicion)
		aux = aux.Next.Next.Next
		// movemos el head hasta la 3era posicion
		println()
		head = head.Next
		head = head.Next
		head = head.Next
		head = head.Next

		println("head tercera posicion(3)", head.Val)

		println("tiene que ser un 3", head.Val)
		break
	}

	return result
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}

	println(swapPairs(head))
}
