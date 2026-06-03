/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head==nil || head.Next==nil{
		return
	}

	s:=head
	f:=head

	for f!=nil && f.Next!=nil{
		s=s.Next
		f=f.Next.Next
	}

	Second:=s.Next
	s.Next=nil

	curr:=Second
	var prev *ListNode
	for curr!=nil{
		next:=curr.Next
		curr.Next=prev
		prev=curr
		curr=next
	}
	first:=head
	second:=prev

	for second!=nil{
		t1:=first.Next
		t2:=second.Next
		first.Next=second
		second.Next=t1
		first=t1
		second=t2
	}







    
}
