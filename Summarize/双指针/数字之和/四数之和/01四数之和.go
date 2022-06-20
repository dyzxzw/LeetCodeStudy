package 四数之和

import "sort"

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-20 14:33
 **/
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	if n < 4 {
		return nil
	}
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			t := target - nums[i] - nums[j]
			left, right := j+1, n-1
			for left < right {
				if nums[left]+nums[right] == t {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if nums[left]+nums[right] > t {
					right--
				} else {
					left++
				}
			}
		}
	}
	return res
}