package 链表相交与成环问题

/**
 * @Description
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 * @Author ZzzWw
 * @Date 2022-06-14 16:28
 **/

// 快慢指针方法
//如果要删除倒数第n个节点
//让fast移动n步，然后让fast和slow同时移动，直到fast指向链表末尾。
//删掉slow所指向的节点就可以了。

type ListNode struct {
	Val  int
	Next *ListNode
}
func removeNthFromEnd(head *ListNode, n int) *ListNode{
	dummyHead:=&ListNode{-1,head}
	slow,fast:=dummyHead,dummyHead
	//注意多了头节点，所以移动n
	for i:=0;i<=n;i++{
		fast=fast.Next
	}
	for fast!=nil{
		fast=fast.Next
		slow= slow.Next
	}
	slow.Next=slow.Next.Next
	return dummyHead.Next
}