package 字母异位

/**
 * @Description

给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。
不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

 * @Author ZzzWw
 * @Date 2022-06-29 15:58
 **/

//暴力会超时，使用滑动窗口
func findAnagrams(s string, p string) (res []int) {
	m,n:=len(s),len(p)
	//如果p的长度都大于s，说明p无法构成异位词
	if n>m{
		return nil
	}
	cntS,cntP:=[26]int{},[26]int{}
	left:=0
	right:=n-1
	//统计p的字符出现的次数
	for i:=range p{
		cntP[p[i]-'a']++
	}
	// 统计cntS中前n个字符出现的次数
	for i:=0;i<n;i++{
		cntS[s[i]-'a']++
	}
	//滑动窗口
	for ;right<m;right++{
		if left>0{
			//关键
			cntS[s[left-1]-'a']-- //窗口右移，原来左边的减去
			cntS[s[right]-'a']++  //窗口右移，新增的要加上
		}
		//判断一次，是否是异位词
		if cntS==cntP{
			res=append(res,left)
		}
		left++
	}
	return res
}