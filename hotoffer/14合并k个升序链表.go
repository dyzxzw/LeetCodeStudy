package hotoffer

import "container/heap"

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-07-30 11:11
 **/

type ListHeap []*ListNode
func (h ListHeap)Len() int {
	return len(h)
}
func(h ListHeap)Less(i,j int)bool{
	return h[i].Val<h[j].Val
}
func(h ListHeap)Swap(i,j int){
	h[i],h[j]=h[j],h[i]
}
func(h *ListHeap)Push(x interface{}){
	*h=append(*h,x.(*ListNode))
}
func(h *ListHeap)Pop()interface{}{
	old:=*h
	n:=len(old)
	x:=old[n-1]
	*h=old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
    if lists==nil || len(lists)==0{
		return nil
	}
	dummyHead:=&ListNode{-1,nil}
	p:=dummyHead

	h:=&ListHeap{}
	heap.Init(h)
	for _,l:=range lists{
		if l!=nil{
			heap.Push(h,l)
		}
	}
	for h.Len()>0{
		m:=heap.Pop(h).(*ListNode)
		p.Next=m
		p=p.Next
		if m.Next!=nil{
			heap.Push(h,m.Next)
		}
	}
	return dummyHead.Next
}