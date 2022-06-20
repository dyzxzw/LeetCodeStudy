package 三数之和

import "sort"

/**
 * @Description
给你一个长度为 n 的整数数组nums和 一个目标值target。
请你从 nums 中选出三个整数，使它们的和与target最接近。
返回这三个数的和。
假定每组输入只存在恰好一个解。
 * @Author ZzzWw
 * @Date 2022-06-20 14:22
 **/

func threeSumClosest(nums []int, target int) int {
	n:=len(nums)
	if n<3{
		return -1
	}
	sort.Ints(nums)
	ans :=nums[0]+nums[1]+nums[2]
	for i:=0;i<n-2;i++{
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		left,right:=i+1,n-1
		for left<right{
			tmp:=nums[i]+nums[left]+nums[right]
			if tmp==target{
				return tmp
			}else if tmp>target{
				right--
			}else{
				left++
			}
			if abs(tmp-target)<abs(ans-target){
				ans=tmp
			}
		}
	}
	return ans
}
func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}