package 翻转链表

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-07 14:41
 **/

func reverseList2(head *ListNode) *ListNode {
	cur:=head
	var prev *ListNode
	prev=nil
	for cur!=nil{
		curNext:=cur.Next
		cur.Next=prev
		prev=cur
		cur=curNext
	}
	return prev
}