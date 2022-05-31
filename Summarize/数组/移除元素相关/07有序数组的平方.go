package 移除元素相关

/**
 * @Description
给你一个按 非递减顺序 排序的整数数组 nums，
返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

 * @Author ZzzWw
 * @Date 2022-05-31 15:17
 **/
//关键在于双指针，并且从后向前填入结果；
//因为负数的平方可能超过正数的平方，所以双指针一前一后进行比较，选择较大额

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	k := len(nums) - 1
	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			res[k] = nums[right] * nums[right]
			right--
			k--
		} else {
			res[k] = nums[left] * nums[left]
			k--
			left++
		}
	}
	return res
}