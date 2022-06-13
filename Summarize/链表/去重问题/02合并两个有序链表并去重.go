package 去重问题

import "container/heap"

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-13 15:44
 **/

//归并
func mergeTwoListsRemoveDup(list1 *ListNode, list2 *ListNode) *ListNode{
	l1,l2:=list1,list2
	dummyHead:=&ListNode{-1,nil}
	cur:=dummyHead
	for l1!=nil && l2!=nil{
		if l1.Val<l2.Val{
			cur.Next=l1
			l1=l1.Next
		}else if l1.Val>l2.Val{
			cur.Next=l2
			l2=l2.Next
		}else{
			cur.Next=l1
			l1=l1.Next
			l2=l2.Next
		}
		cur=cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummyHead.Next
}
func mergeTwoListsHeapRemoveDup(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1==nil{
		return list2
	}
	if list2==nil{
		return list1
	}
	if list1==nil && list2==nil{
		return nil
	}
	h:=&IntHeap{}
	heap.Init(h)
	dummyHead:=&ListNode{-1,nil}
	p:=dummyHead
	heap.Push(h,list1)
	heap.Push(h,list2)
	for h.Len()>0{
		min:=heap.Pop(h).(*ListNode)
		//去重
		if min.Val!=p.Val{
			p.Next=min
			p=p.Next
		}
		if min.Next!=nil{
			heap.Push(h,min.Next)
		}
	}
	return dummyHead.Next
}