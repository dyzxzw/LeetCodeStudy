package hotoffer

/**
 * @Description
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。
它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。


和两数相加1的区别是：这个是正序相加的

7--->2--->4--->3
    5---->6---->4

7807

思想：逆序，然后按照1的方式相加
 * @Author ZzzWw
 * @Date 2022-07-29 21:22
 **/

//非递归
func reverseList1(head *ListNode)*ListNode{
	cur:=head
	var pre *ListNode
	pre=nil
	for cur!=nil{
		curNext:=cur.Next
		cur.Next=pre
		pre=cur
		cur=curNext
	}
	return pre
}

//递归
func reverseList2(head*ListNode)*ListNode{
	if head==nil || head.Next==nil{
		return head
	}
	newHead:=reverseList2(head.Next)
	head.Next.Next=head
	head.Next=nil
	return newHead
}
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	if l1==nil && l2==nil{
		return nil
	}
	l1=reverseList1(l1)
	l2=reverseList1(l2)

	dummyHead:=&ListNode{-1,nil}
	p:=dummyHead
	carry:=0
	for l1!=nil || l2!=nil{
		a,b:=0,0
		if l1!=nil{
			a=l1.Val
		}else{
			a=0
		}
		if l2!=nil{
			b=l2.Val
		}else{
			b=0
		}

		newNode:=&ListNode{(a+b+carry)%10,nil}
		p.Next=newNode
		p=p.Next
		carry=(a+b+carry)/10
		if l1!=nil{
			l1=l1.Next
		}
		if l2!=nil{
			l2=l2.Next
		}
	}
	if carry>0{
		p.Next=&ListNode{carry,nil}
	}
	dummyHead=reverseList1(dummyHead.Next)
	return dummyHead
}