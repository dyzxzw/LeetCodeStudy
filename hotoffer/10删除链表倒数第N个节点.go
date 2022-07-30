package hotoffer

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-07-30 10:49
 **/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
   dummyHead:=&ListNode{-1,head}
   fast:=dummyHead
   for i:=0;i<=n;i++{
	   fast=fast.Next
   }
   slow:=dummyHead
   for fast!=nil{
	   slow=slow.Next
	   fast=fast.Next
   }
   slow.Next=slow.Next.Next
   return dummyHead.Next
}