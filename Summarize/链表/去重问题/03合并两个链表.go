package 去重问题

/**
 * @Description
给你两个链表list1 和list2，它们包含的元素分别为n 个和m 个。
请你将list1中下标从 a 到 b 的全部节点都删除，并将list2接在被删除节点的位置。

list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
输出：[0,1,2,1000000,1000001,1000002,5]
解释：我们删除 list1 中下标为 3 和 4 的两个节点，并将 list2 接在该位置。

 * @Author ZzzWw
 * @Date 2022-06-13 15:48
 **/

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	//1.找到要删除的第一个下标的前一个节点
	cur1 := list1
	for i := 0; i < a-1; i++ {
		cur1 = cur1.Next
	}
	//2.找到要删除的最后一个下标的后一个节点
	cur2 := list1
	for i := 0; i < b+1; i++ {
		cur2 = cur2.Next
	}
	//3.
	cur1.Next = list2
	//4.找到List2的尾节点
	cur3 := list2
	for cur3.Next != nil {
		cur3 = cur3.Next
	}
	//5.
	cur3.Next = cur2
	return list1
}