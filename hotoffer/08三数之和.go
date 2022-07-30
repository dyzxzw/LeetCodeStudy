package hotoffer

import "sort"

/**
 * @Description
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，
使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。


 * @Author ZzzWw
 * @Date 2022-07-30 10:34
 **/
func threeSum(nums []int) [][]int {
	if len(nums)<3{
		return nil
	}
	res:=make([][]int,0)
	sort.Ints(nums)
	if nums[0]>0{
		return nil
	}
	for i:=0;i<len(nums)-2;i++{
		if nums[i]>0{
			return res
		}
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		left,right:=i+1,len(nums)-1
		target:=0-nums[i]
		for left<right {
			if nums[left]+nums[right] == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				//去重
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				right--
				left++
			}else if nums[left]+nums[right]>target{
				right--
			}else{
				left++
			}

		}
	}
	return res
}