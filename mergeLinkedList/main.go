package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
   if(list1!=null && list2!=null){
        if(list1.val<list2.val){
            list1.next=mergeTwoLists(list1.next,list2);
            return list1;
            }
            else{
                list2.next=mergeTwoLists(list1,list2.next);
                return list2;
        }
        }
        if(list1==null)
            return list2;
        return list1;
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			list1.Next = mergeTwoLists(list1.Next, list2)
			return list1
		} else {
			list2.Next = mergeTwoLists(list1, list2.Next)
			return list2
		}
	}

	if list1 == nil {
		return list2
	}
	return list1
}
