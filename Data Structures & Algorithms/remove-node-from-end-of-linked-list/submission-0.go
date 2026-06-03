/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy:=&ListNode{}
    dummy.Next=head
	ptr1:=dummy
	ptr2:=dummy

	for i:=0;i<n;i++{
		ptr2=ptr2.Next
	}

	for ptr2.Next!=nil{
		ptr1=ptr1.Next
		ptr2=ptr2.Next
	}

	ptr1.Next=ptr1.Next.Next

	return dummy.Next    
}
