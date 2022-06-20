package 三数之和

import "sort"

/**
 * @Description
nums[i]+nums[j]+nums[k]<target 成立的三元组的个数

nums =-2,0,1,3 target=2
输出： 2
-2，0，1
-2，0，3
 * @Author ZzzWw
 * @Date 2022-06-20 14:26
 **/
func threeSumSmaller(nums []int, target int) int {
	n:=len(nums)
	if n<3{
		return 0
	}
	res:=0
	sort.Ints(nums)
	for i:=0;i<n-2;i++{
		left,right:=i+1,n-1
		for left<right{
			if nums[i]+nums[left]+nums[right]<target{
				//right之前的所有三元组都满足条件，所以res+=right-left
				res+=right-left
				left++
			}else{
				right--
			}
		}
	}
	return res
}

//如果要求去掉重复项，返回个数
func threeSumSmaller2(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	ans := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			if nums[i]+nums[left]+nums[right] < target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				right--
				left++

			} else {
				right--
			}
		}
	}

	return len(ans)
}