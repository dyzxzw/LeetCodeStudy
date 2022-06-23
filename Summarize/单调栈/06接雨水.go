package 单调栈

import "container/list"

/**
 * @Description
42 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水

什么情况形成凹槽，然后才能有雨水？
遇到比栈顶元素大的，才能形成凹槽
且至少要有三个柱子

 * @Author ZzzWw
 * @Date 2022-06-23 17:07
 **/
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	sum := 0
	st := list.New()
	st.PushBack(0)
	for i := 1; i < len(height); i++ {
		if height[i] < height[st.Back().Value.(int)] {
			st.PushBack(i)
		} else if height[i] == height[st.Back().Value.(int)] {
			st.PushBack(i)
		} else {
			for st.Len() > 0 && height[i] > height[st.Back().Value.(int)] {
				mid := st.Remove(st.Back()).(int)
				if st.Len() > 0 {
					h := min(height[i], height[st.Back().Value.(int)]) - height[mid]
					w := i - st.Back().Value.(int) - 1
					sum += h * w
				}
			}
			st.PushBack(i)
		}
	}
	return sum
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}