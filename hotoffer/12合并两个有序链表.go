package hotoffer

import "container/heap"

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-07-30 10:57
 **/

//归并
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead:=&ListNode{-1,nil}
	p:=dummyHead
	for list1!=nil && list2!=nil{
		if list1.Val<list2.Val{
			p.Next=list1
			list1=list1.Next
		}else{
			p.Next=list2
			list2=list2.Next
		}
		p=p.Next
	}
	if list1!=nil{
		p.Next=list1
	}
	if list2!=nil{
		p.Next=list2
	}
	return dummyHead.Next
}

//堆
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode{
	dummyHead:=&ListNode{-1,nil}
	p:=dummyHead
	if list1==nil{
		return list2
	}
	if list2==nil{
		return list1
	}
	if list1==nil&&list2==nil{
		return nil
	}
	h:=&IntHeap{}
	heap.Init(h)
	heap.Push(h,list1)
	heap.Push(h,list2)
	for h.Len()>0{
		minList:=heap.Pop(h).(*ListNode)
		p.Next=minList
		p=p.Next
		if minList.Next!=nil{
			heap.Push(h,minList.Next)
		}
	}
	return dummyHead.Next
}

type IntHeap []*ListNode
func (h IntHeap)Len()int{ return len(h) }
func(h IntHeap)Less(i,j int)bool { return h[i].Val<h[j].Val }
func(h IntHeap)Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func(h*IntHeap)Push(x interface{}) {
	*h=append(*h,x.(*ListNode))
}
func(h *IntHeap)Pop() interface{}{
	old:=*h
	n:=len(old)
	x:=old[n-1]
	*h=old[:n-1]
	return x
}