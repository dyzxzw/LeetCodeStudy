package 链表相交与成环问题

/**
 * @Description
判断链表
1. 是否有环
2. 找出环的入口
3. 环中节点个数
4.打印首尾结点
 * @Author ZzzWw
 * @Date 2022-06-15 14:48
 **/

func hasCycle(head *ListNode) bool{
	fast:=head
	slow:=head
	for fast!=nil && fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
		if slow==fast{
			return true
		}
	}
	return false
}