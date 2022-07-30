package hotoffer

import "sort"

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-07-30 15:05
 **/
func nextPermutation(nums []int) {
	if nums==nil || len(nums)<=1{
		return
	}
	i:=len(nums)-1
	for ;;{
		j:=i
		i--
		if nums[i]<nums[j]{
			k:=len(nums)
			for ;;{
				k--
				if nums[k]>nums[i]{
					nums[i],nums[k]=nums[k],nums[i]
					//Reverse
					end:=len(nums)-1
					for j<end{
						nums[j],nums[end]=nums[end],nums[j]
						end--
						j++
					}
					return
				}
			}
		}
		if i==0{
			sort.Ints(nums)
			return
		}
	}
	return
}