package K个一组链表问题

/**
 * @Description
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）
 * @Author ZzzWw
 * @Date 2022-06-15 15:10
 **/

type ListNode struct {
	Val  int
	Next *ListNode
}
func swapPairs(head *ListNode) *ListNode {
	dummyHead:=&ListNode{-1,head}
	cur:=dummyHead
	for cur.Next!=nil && cur.Next.Next!=nil{
		s1:=cur.Next
		s2:=cur.Next.Next
		//交换
		s1.Next=s2.Next
		s2.Next=s1
		cur.Next=s2
		//cur前进
		cur=s1
	}
	return dummyHead.Next
}