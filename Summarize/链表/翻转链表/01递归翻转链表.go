package 翻转链表

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-07 14:31
 **/

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

/*
1->2->3->4->nil

递归：
栈空间：
reverseList2(4) head=3 newHead=4
reverseList2(3) head=2 newHead=3
reverseList2(2) head=1 newHead=2

出栈：
1.
reverseList2(4) head=3 newHead=4
head.Next.Next=head:
head.Next=nil
3->4->nil变成了 nil<--3<--4
2.
reverseList2(3) head=2 newHead=3
head.Next.Next=head:
head.Next=nil
2->3->nil变成了 nil<--2<--3

3. 其余同理
 */