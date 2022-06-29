package 交集

import "sort"

/**
 * @Description

给你两个整数数组nums1 和 nums2 ，
请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，
应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。
可以不考虑输出结果的顺序。


输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：2,2

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]

 * @Author ZzzWw
 * @Date 2022-06-29 15:35
 **/

func intersect(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			res = append(res, nums1[i])
			i++
			j++
		}
	}
	return res
}

