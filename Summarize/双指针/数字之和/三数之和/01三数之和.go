package 三数之和

import "sort"

/**
 * @Description
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，
使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
 * @Author ZzzWw
 * @Date 2022-06-20 14:18
 **/
func threeSum(nums []int) [][]int{
	n:=len(nums)
	if n<3{
		return nil
	}
	res:=make([][]int,0)
	sort.Ints(nums) //排序
	if nums[0]>0{
		return nil
	}
	for i:=0;i<n-2;i++{
		if nums[i]>0{
			return res
		}
		//去重
		if i>0 && nums[i]==nums[i-1]{
			continue
		}
		left,right:=i+1,n-1
		target:=0-nums[i]
		for left<right{
          if nums[left]+nums[right]==target{
			  res=append(res,[]int{nums[i],nums[left],nums[right]})
			  //去重
			  for left<right && nums[right]==nums[right-1]{
				  right--
			  }
			  for left<right&& nums[left]==nums[left+1]{
				  left++
			  }
			  right--
			  left++
		  }else if nums[left]+nums[right]>target{
			  right--
		  } else{
			left++
			}
		}
	}
	return res
}