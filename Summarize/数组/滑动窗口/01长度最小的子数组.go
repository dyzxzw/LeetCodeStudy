package 滑动窗口

/**
 * @Description
给定一个含有n个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组
[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。
如果不存在符合条件的子数组，返回 0 。

 * @Author ZzzWw
 * @Date 2022-05-31 15:22
 **/

func minSubArrayLen(target int, nums []int) int {
     left:=0 //滑动窗口的左边界
	 sum:=0
	 res:=len(nums)+1 ///不能超过数组最大长度
	 for i:=0;i<len(nums);i++{
		 sum+=nums[i]
		 for sum>=target{
			 res=min(i-left+1,res)
			 sum-=nums[left]
			 left++ //窗口移动
		 }
	 }
	 if res==len(nums)+1{
		 return 0 // 不存在
	 }
	 return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}