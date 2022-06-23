package 单调栈

import "container/list"

/**
 * @Description

给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

但是和以往的不同，
这道题是小于栈顶元素，才出栈
 * @Author ZzzWw
 * @Date 2022-06-23 17:08
 **/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func largestRectangleArea(heights []int) int {
	result := 0
	heights = append(heights, 0)
	heights = append([]int{0}, heights[:]...)

	st := list.New()
	st.PushBack(0)
	for i := 1; i < len(heights); i++ {
		if heights[i] > heights[st.Back().Value.(int)] {
			st.PushBack(i)
		} else if heights[i] == heights[st.Back().Value.(int)] {
			st.PushBack(i)
		} else {
			for st.Len() > 0 && heights[i] < heights[st.Back().Value.(int)] {
				mid := st.Remove(st.Back()).(int)
				if st.Len() > 0 {
					h := heights[mid]
					w := i - st.Back().Value.(int) - 1
					result = max(result, h*w)
				}
			}
			st.PushBack(i)
		}
	}
	return result
}