package 两数之和

/**
 * @Description
给定一个整数数组 nums 和一个整数目标值 target
请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
 * @Author ZzzWw
 * @Date 2022-06-20 14:34
 **/
func twoSum(nums []int, target int) []int {
	hsmap:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		if v,ok:=hsmap[target-nums[i]];ok{
			return []int{i,v}
		}
		hsmap[nums[i]]=i
	}
	return nil
}