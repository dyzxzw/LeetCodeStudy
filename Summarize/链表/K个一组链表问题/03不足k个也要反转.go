package K个一组链表问题

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-15 15:36
 **/
func reverse2(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	cur := head
	for prev != tail {
		curNext := cur
		cur.Next = prev
		prev = cur
		cur = curNext
	}
	return tail, head
}
func reverseKGroup2(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{-1, head}
	tmpNode := dummyHead
	for head != nil {
		tail := tmpNode
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail.Next == nil {
				break
			}
		}
		//1.
		tailNextNode := tail.Next
		//2.
		head, tail = reverse2(head, tail)
		//3.
		tmpNode.Next = head
		tail.Next = tailNextNode
		//4.
		tmpNode = tail
		head = tail.Next
	}
	return dummyHead.Next
}