package 链表设计与移除

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * @Description
给你一个链表的头节点 head 和一个整数 val
请你删除链表中所有满足 Node.val == val 的节点
并返回 新的头节点
 * @Author ZzzWw
 * @Date 2022-06-07 13:41
 **/

type ListNode struct {
	Val  int
	Next *ListNode
}
func removeElements(head *ListNode, val int) *ListNode{
       //1.设置虚拟头结点
	   dummyHead:=&ListNode{-1,head}
	   cur:=dummyHead
	   //2.找到要删除结点的前一个结点
	   for cur!=nil &&cur.Next!=nil{
		   if cur.Next.Val==val{
			   cur.Next=cur.Next.Next
		   }else{
			   cur=cur.Next
		   }
	   }
	   return dummyHead.Next
}

//构造测试用例，把输入的数组转化为链表的形式
func getListNode(nums[]int)*ListNode{
	if nums==nil{
		return nil
	}
	head:=&ListNode{nums[0],nil}
	cur:=head
	for i:=1;i<len(nums);i++{
		cur.Next=&ListNode{nums[i],nil}
		cur=cur.Next
	}
	return head
}
func removeElementsTest(){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()//读取一行
	var nums []int
	str:=strings.Split(input.Text(),",")
	for _,v:=range str{
		value,_:=strconv.Atoi(v)
		nums=append(nums,value)
	}
	//数组转化为链表
	head:=getListNode(nums)
	input.Scan()//读取val
	val,_:=strconv.Atoi(input.Text())
	head=removeElements(head,val)
	//输出结果
	cur:=head
	for cur!=nil{
		fmt.Println(cur.Val," ")
		cur=cur.Next
	}
}