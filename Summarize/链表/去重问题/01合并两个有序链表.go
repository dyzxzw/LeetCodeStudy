package 去重问题

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * @Description
将两个升序链表合并为一个新的 升序 链表并返回。
新链表是通过拼接给定的两个链表的所有节点组成的。
 * @Author ZzzWw
 * @Date 2022-06-13 14:41
 **/

type ListNode struct {
	Val  int
	Next *ListNode
}
//归并
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode{
	l1,l2:=list1,list2
	dummyHead:=&ListNode{-1,nil}
	cur:=dummyHead
	for l1!=nil && l2!=nil{
		if l1.Val<l2.Val{
			cur.Next=l1
			l1=l1.Next
		}else{
			cur.Next=l2
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


//最小堆

type IntHeap []*ListNode
//Len Less Swap Push Pop
func(h IntHeap)Len() int{
	return len(h)
}
func(h IntHeap)Less(i,j int)bool{
	return h[i].Val<h[j].Val
}
func(h IntHeap)Swap(i,j int){
	h[i],h[j]=h[j],h[i]
}
func(h*IntHeap)Push(x interface{}){
	*h=append(*h,x.(*ListNode))
}
func(h*IntHeap)Pop()interface{}{
	old:=*h
	n:=len(old)
	x:=old[n-1]
	*h=old[:n-1]
	return x
}

func mergeTwoListsHeap(list1 *ListNode, list2 *ListNode) *ListNode{
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
		p.Next=min
		p=p.Next
		if min.Next!=nil{
			heap.Push(h,min.Next)
		}
	}
	return dummyHead.Next
}







//数组转化为链表
func getListNode(nums[]int)*ListNode{
   if nums==nil{
	   return nil
   }
   head:=&ListNode{nums[0],nil}
   cur:=head
   for i:=1;i<len(nums);i++{
	   cur.Next=&ListNode{nums[i],nil}
	   cur=cur.Next
   }
   return head
}

func mergeTwoListsTest(){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()//读取list1
	var list1 []int
	str1:=strings.Split(input.Text(),",")
	for _,v:=range str1{
		value,_:=strconv.Atoi(v)
		list1=append(list1,value)
	}
	//数组转化为链表
	l1:=getListNode(list1)

	input.Scan()//读取list2
	var list2 []int
	str2:=strings.Split(input.Text(),",")
	for _,v:=range str2{
		value,_:=strconv.Atoi(v)
		list2=append(list2,value)
	}
	//数组转化为链表
	l2:=getListNode(list2)
	//输出结果
	head:=mergeTwoLists(l1,l2)
	cur:=head
	for cur!=nil{
		fmt.Print(cur.Val," ")
		cur=cur.Next
	}
}