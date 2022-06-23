package 单调栈

import (
	"math"
	"strconv"
)

/**
 * @Description
556
给你一个正整数n ，请你找出符合条件的最小整数，其由重新排列 n中存在的每位数字组成，并且其值大于 n 。
如果不存在这样的正整数，则返回 -1 。

注意 ，返回的整数应当是一个 32 位整数 ，如果存在满足题意的答案，但不是 32 位整数 ，同样返回 -1 。

和下一个最大排列类似，基本思路：
从后向前找第一个升序对（i，j)
然后从后向前找第一个K ：k所指的元素的值大于i所指的元素的值
交换i,k的值
将j到end的元素逆序，Reverse操作

但这道题，还需要将整数先转化为字符串

 * @Author ZzzWw
 * @Date 2022-06-23 17:02
 **/

func nextGreaterElement556(n int) int {
	//定义答案
	res := -1
	//将数字转化为字符串
	str := strconv.Itoa(n)
	//将字符串转化为切片
	strslice := []byte(str)
	if nextMax(&strslice) == true {
		//将下一个最大排序转化为整数
		res, _ = strconv.Atoi(string(strslice))
		//如果转化后的整数超过整数最大上限，则不存在
		if res > math.MaxInt32 {
			return -1
		}
	}
	return res
}

//下一个更大元素
func nextMax(n *[]byte) bool {
	if len(*n) == 0 {
		return false
	}
	if len(*n) == 1 {
		return false
	}
	i := len(*n) - 1
	for {
		j := i
		i--
		if (*n)[i] < (*n)[j] {
			k := len(*n)
			for {
				k--
				if (*n)[k] > (*n)[i] {
					(*n)[k], (*n)[i] = (*n)[i], (*n)[k]
					//将j到end的元素全部逆序，Reverse操作
					//reverse(j,end)
					end := len(*n) - 1
					for j < end {
						//双指针依次交换每个元素的值
						(*n)[j], (*n)[end] = (*n)[end], (*n)[j]
						j++
						end--
					}
					return true
				}
			}
		}
		if i == 0 {
			return false
		}
	}
	return false
}