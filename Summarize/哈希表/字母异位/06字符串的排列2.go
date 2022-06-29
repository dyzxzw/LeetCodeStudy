package 字母异位

import "sort"

/**
 * @Description
输入一个字符串，打印出该字符串中字符的所有排列。
你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]

 * @Author ZzzWw
 * @Date 2022-06-29 16:12
 **/


func permutation(s string) (ans []string) {
   t:=[]byte(s)
   sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
   for{
	   ans = append(ans,string(t))
		   if !nextPermutation(t){
			   break
		   }
	   }
   return ans
}


func nextPermutation(nums []byte)bool {
	if nums==nil || len(nums)<=1{
		return false
	}
	i:=len(nums)-1
	for ;;{
		j:=i
		i--
		if nums[i]<nums[j]{
			k:=len(nums)//注意这个地方，不要写成len(nums)-1
			for ;;{
				k--
				if nums[k]>nums[i]{
					nums[i],nums[k]=nums[k],nums[i]
					//Reverse
					end := len(nums) - 1
					for j < end {
						nums[j], nums[end] = nums[end], nums[j]
						j++
						end--
					}
					return true
				}
			}
		}
		if  i==0{
			return false
		}
	}
	//返回
	return false
}