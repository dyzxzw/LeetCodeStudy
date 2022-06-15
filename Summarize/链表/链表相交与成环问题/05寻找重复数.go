package 链表相交与成环问题

/**
 * @Description

给定一个包含n + 1 个整数的数组nums ，其数字都在[1, n]范围内（包括 1 和 n）
可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
 * @Author ZzzWw
 * @Date 2022-06-15 15:01
 **/

/*
1,3,4,2,2
映射关系：
0->1
1->3
2->4
3->2
4->2
如果有环：
0->1->3->2->4->2->4->2->……
快慢指针就相当于
一个从下标0出发，一个多走一步就相当于进一步映射：从下标的值出发
简单来说就是
slow = nums[index]
fast = nums[ nums[index]]

*/
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	//设置新的指针从头出发，找到相遇点
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}