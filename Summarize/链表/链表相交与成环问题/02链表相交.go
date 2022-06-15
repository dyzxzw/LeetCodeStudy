package 链表相交与成环问题

/**
 * @Description

给你两个单链表的头节点 headA 和 headB
请你找出并返回两个单链表相交的起始节点。
如果两个链表没有交点，返回 null 。

下面提供双指针方法的正确性证明。考虑两种情况，第一种情况是两个链表相交，第二种情况是两个链表不相交。
情况一：两个链表相交
链表 headA 和 headB 的长度分别是 m 和 n。
假设链表 headA 的不相交部分有 a 个节点
链表headB 的不相交部分有 b 个节点，两个链表相交的部分有 c 个节点
则有 a+c=m，b+c=n

1.如果a=b 则两个指针会同时到达两个链表相交的节点，此时返回相交的节点
2.如果a!=b,则 pa指针移动a+c+b,pb移动b+c+a.
两个指针会最终同时到达两个链表相交的节点，该节点也是两个指针第一次同时指向的节点，此时返回相交的节点。

情况二：两个链表不相交
1.如果m==n,则指针最终会到达nil
2.如果m！=n,由于没用交点，也会返回nil
 * @Author ZzzWw
 * @Date 2022-06-15 14:44
 **/

func getIntersectionNode(headA, headB *ListNode) *ListNode{
     if headA==nil || headB==nil{
		 return nil
	 }
	 pa,pb:=headA,headB
	 for pa!=pb{
		 if pa==nil{
			 pa=headB
		 }else{
			 pa=pa.Next
		 }
		 if pb==nil{
			 pb=headA
		 }else{
			 pb=pb.Next
		 }
	 }
	 return pa
}