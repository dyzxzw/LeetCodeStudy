package 链表相交与成环问题

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * @Description
          1. 是否有环
          2. 找出环的入口
          3. 环中节点个数
          4.打印首尾结点
 * @Author ZzzWw
 * @Date 2022-06-15 14:54
 **/
func detectCycle(head *ListNode) (*ListNode, *ListNode) {
	if head==nil ||head.Next==nil{
		fmt.Println("环中结点个数为：",0)
		return nil,nil
	}
	nodesize:=1//环中结点个数
	//1.快慢指针判断是否有环
	fast,slow:=head,head
	for fast!=nil && fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
		//2.如果相遇，说明有环
		if slow==fast{
			//3.设置新指针从头出发
			cur:=head
			for cur!=fast{
				cur=cur.Next
				fast=fast.Next
			}
			//4.相遇的地方就是环的入口
			//return cur 或 return fast
			//5.找环中结点个数
			for fast.Next!=cur{
				fast=fast.Next
				nodesize++
			}
			fmt.Println("环中结点个数为：",nodesize)
			//返回首尾结点
			return cur,fast
		}
	}
	fmt.Println("环中的结点个数为：", 0)
	return nil, nil
}

func getListNode(nums []int, pos int) *ListNode {
	if nums == nil {
		return nil
	}
	head := &ListNode{nums[0], nil}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{nums[i], nil}
		cur = cur.Next
	}
	//找到Pos位置的结点

	if pos > 0 {
		tmp := head
		for i := 0; i < pos; i++ {
			tmp = tmp.Next
		}
		//成环
		cur.Next = tmp
	}

	return head
}

func detectCycleTest(){
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	nums1 := []int{}
	str := strings.Split(input.Text(), ",")
	for _, v := range str {
		value, _ := strconv.Atoi(v)
		nums1 = append(nums1, value)
	}
	input.Scan()
	pos, _ := strconv.Atoi(input.Text())
	list := getListNode(nums1, pos)
	start, end := detectCycle(list)
	if start == nil || end == nil {
		fmt.Println("链表中没有环")
	} else {
		fmt.Println("首结点：", start.Val, " ", "尾部结点:", end.Val)
	}
}