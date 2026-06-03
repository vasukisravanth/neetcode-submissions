/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func MergeTwoLists(l1 *ListNode,l2 *ListNode)*ListNode{
	dummy:=&ListNode{}
	curr:=dummy
	for l1!=nil&&l2!=nil{
       if l1.Val<=l2.Val{
		curr.Next=l1
		l1=l1.Next
	   }else{
		curr.Next=l2
		l2=l2.Next

	   }
	   curr=curr.Next
	}

	if l1!=nil{
		curr.Next=l1
	}
	if l2!=nil{
		curr.Next=l2
	}

	return dummy.Next
 }

 func MergeLists(lists []*ListNode,start int,end int)*ListNode{
	if start==end{
		return lists[start]
	}

	mid:=start+(end-start)/2
    l1:=MergeLists(lists,start,mid)
	l2:=MergeLists(lists,mid+1,end)

	return MergeTwoLists(l1,l2)
 }

func mergeKLists(lists []*ListNode) *ListNode {

	if len(lists)==0{
		return nil
	}

	if len(lists)==1{
		return lists[0]
	}

	if len(lists)==2{
		return MergeTwoLists(lists[0],lists[1])
	}

	return MergeLists(lists,0,len(lists)-1)
    
}
