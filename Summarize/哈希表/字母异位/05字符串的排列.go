package 字母异位

/**
 * @Description

给你两个字符串s1和s2
写一个函数来判断 s2 是否包含 s1的排列。
如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。

输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").

输入：s1= "ab" s2 = "eidboaoo"
输出：false
解释：s1的排列ab,ba，s2都不包含

 * @Author ZzzWw
 * @Date 2022-06-29 16:06
 **/

//和04找到字符串中所有字母异位词思想一模一样
func checkInclusion(s1 string, s2 string) bool {
	m,n:=len(s1),len(s2)
	if m>n{
		return false
	}
	cntS1:=[26]int{}
	cntS2:=[26]int{}

	left:=0
	right:=m-1
	//先判断一次前m个是否相等
	if cntS1 == cntS2 {
		return true
	}

	for i:=range s1{
		cntS1[s1[i]-'a']++
	}
	for i:=0;i<m;i++{
		cntS2[s2[i]-'a']++
	}

	for ;right<n;right++{
		if left>0{
			cntS2[s2[left-1]-'a']--
			cntS2[s2[right]-'a']++
		}
		if cntS2==cntS1{
			return true
		}
		left++
	}
	return false
}