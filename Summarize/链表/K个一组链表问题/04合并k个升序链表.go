package K个一组链表问题

import (
	"container/heap"
	"math"
)

/**
 * @Description

给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。


 * @Author ZzzWw
 * @Date 2022-06-15 15:37
 **/
//1.归并合并
func mergeKLists(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{math.MaxInt32, nil}
	cur := dummyHead
	for _, list := range lists {
		cur = mergerTwoList(cur, list)
	}
	return dummyHead.Next
}

//合并两个有序链表
func mergerTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	l1 := list1
	l2 := list2
	dummyHead := &ListNode{-1, nil}
	cur := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummyHead.Next
}

//2.最小堆

//最小堆的模板

type IntHeap []*ListNode

func (h IntHeap) Len() int { return len(h) }

// 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆

func (h IntHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }

// 绑定pop方法，从最后拿出一个元素并返回 联系堆排序，从数组最后一个位置出去

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Time: O(n*log(k)), Space: O(k) n是总节点数量，k是链表个数
func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil // 链表为空或长度为0，直接返回空指针
	}

	h := &IntHeap{} // 最小堆用于维护当前k个节点
	heap.Init(h)    // 用于节点间的比较

	for _, list := range lists {
		// 数组中非空的链表加入到最小堆
		if list != nil {
			heap.Push(h, list)
		}
	}
	// 定义dummy节点用于统一处理
	dummy := &ListNode{}
	p := dummy // p初始指向dummy节点

	// 当最小堆不为空时，不断执行以下操作
	for h.Len() > 0 { // 取出堆顶元素，即取出最小值节点
		min := heap.Pop(h).(*ListNode)
		p.Next = min // 游标p指向最小值节点
		p = p.Next   // p向后移动一个位置
		// 这样就确定一个节点在合并链表中的位置
		if min.Next != nil { // 如果最小值节点后面的节点非空
			heap.Push(h, min.Next) // 则把最小值节点后面的节点加入到最小堆中
		}
	}
	return dummy.Next // 最后只要返回dummy.Next即可

}