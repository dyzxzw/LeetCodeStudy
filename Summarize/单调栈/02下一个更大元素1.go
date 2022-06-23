package 单调栈

import "container/list"

/**
 * @Description

nums1中数字x的 下一个更大元素 是指x在nums2 中对应位置 右侧 的 第一个 比x大的元素。
给你两个 没有重复元素 的数组nums1 和nums2 ，下标从 0 开始计数，其中nums1是nums2的子集。

输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
输出：[-1,3,-1]


要遍历Nums2,看每个数字的下一个最大元素是多少，这个地方和和每日温度类似。
然后还要限制 NUms2遍历的数字存在于nums1:哈希表的构造 关键字是nums1[i]，因为要看每个nums1[i]是否存在

 * @Author ZzzWw
 * @Date 2022-06-23 16:02
 **/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
      hashmap:=make(map[int]int,0)
	  for i:=0;i<len(nums1);i++{
		  hashmap[nums1[i]]=i
	  }
	  res:=make([]int,len(nums1))
	  for i := 0; i < len(nums1); i++ {
		res[i] = -1
	}
	st := list.New()
	st.PushBack(0)

	for i:=1;i<len(nums2);i++{
		for st.Len()>0&& nums2[i]>nums2[st.Back().Value.(int)]{
			if index,ok:=hashmap[nums2[st.Back().Value.(int)]];ok{
				res[index]=nums2[i]
			}
			st.Remove(st.Back())
		}
		st.PushBack(i)
	}
	return res
}