package K个一组链表问题

/**
 * @Description

此题是 25题，属于困难题
给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
K是一个正整数，它的值小于或等于链表的长度。
如果节点总数不k的整数倍，那么请将最后剩余的节点保持原有顺序。

 * @Author ZzzWw
 * @Date 2022-06-15 15:12


考查小知识点链表的反转+链表的断开与连接 其实链表的反转不同于一般的反转

 **/

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode){
	pre:=tail.Next
	cur:=head
	for pre!=tail{
		curNext:=cur.Next
		cur.Next=pre
		pre=cur
		cur=curNext
	}
	return tail,head
}

func reverseKGroup(head *ListNode, k int) *ListNode{
	dummyHead:=&ListNode{-1,head}
	tmpNode:=dummyHead
	for head!=nil{
		tail:=tmpNode
		//找到k个一组链表的尾
		for i:=0;i<k;i++{
			tail=tail.Next
			//不足K个直接返回，不用反转
			if tail==nil{
				return dummyHead.Next
			}
		}
		//1.保存尾部结点的下一个结点
		tailNextNode:=tail.Next
		//2.反转链表，返回新的头和尾
		head,tail=reverse(head,tail)
		//3.把反转后的链表装入原来链表
		tmpNode.Next=head
		tail.Next=tailNextNode
		//4.重复以上操作
		tmpNode=tail
		head=tail.Next
	}
	return dummyHead.Next
}