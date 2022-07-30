package hotoffer

/**
 * @Description
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
 * @Author ZzzWw
 * @Date 2022-07-13 20:01
 **/
func twoSum(nums []int, target int) []int {
	h:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		if v,ok:=h[nums[i]];ok{
			return []int{v,i}
		}
		h[nums[i]]=i
	}
	return nil
}