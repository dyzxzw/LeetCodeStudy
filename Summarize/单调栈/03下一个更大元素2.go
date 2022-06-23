package 单调栈

import "container/list"

/**
 * @Description
503
给定一个循环数组nums（nums[nums.length - 1]的下一个元素是nums[0]），
返回nums中每个元素的 下一个更大元素 。

输入: nums = [1,2,1]
输出: [2,-1,2]

和每日温度完全类似，只不过多了循环的写法
将Nums： 1 2 1  1 2 1
 * @Author ZzzWw
 * @Date 2022-06-23 16:23
 **/

func nextGreaterElements(nums []int) []int {
	st := list.New()
	st.PushBack(0)
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}
    for i:=1;i<2*len(nums);i++{
		for st.Len()>0 && nums[i%len(nums)]>nums[st.Back().Value.(int)]{
			res[st.Back().Value.(int)]=nums[i%len(nums)]
			st.Remove(st.Back())
		}
		st.PushBack(i%len(nums))
	}
	return res
}