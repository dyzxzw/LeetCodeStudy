package 交集

import "sort"

/**
 * @Description

给定两个数组 nums1 和 nums2 ，返回 它们的交集 。
输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。

输入：nums1 = [1,2,2,1], nums2 = [2,2]
//交集 去掉重复的数字
输出：[2]


 * @Author ZzzWw
 * @Date 2022-06-29 15:30
 **/

//方法1 利用集合 时间空间复杂度 O(min(m,n))=O(m+n)
func intersection1(nums1 []int, nums2 []int) []int {
	set:=make(map[int]struct{})
	for _,v:=range nums1{
		set[v]= struct{}{}
	}
	res:=[]int{}
	for _,v:=range nums2{
		if _,ok:=set[v];ok{
			res=append(res,v)
			delete(set,v)
		}
	}

	return res
}

//方法2 排序+双指针
func intersection2(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Slice(nums2,func(i ,j int) bool {
		return nums2[i]<nums2[j]
	})
	i,j:=0,0
	for i<len(nums1) && j<len(nums2){
		if nums1[i]<nums2[j]{
			i++
		}else if nums1[i]>nums2[j]{
			j++
		}else{
			if res == nil || nums1[i]!=res[len(res)-1]{
				res=append(res,nums1[i])
			}
			i++
			j++
		}
	}
	return res
}