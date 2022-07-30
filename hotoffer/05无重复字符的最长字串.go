package hotoffer

/**
 * @Description
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
abcabcbb
3

注意：
  子串 是连续的  abc
  子序列 不是连续的 abb
 * @Author ZzzWw
 * @Date 2022-07-29 21:56
 **/
// 滑动窗口+集合

func lengthOfLongestSubstring(s string) int {
    left,res:=0,0
	set:=map[byte]struct{}{}
	for i:=0;i<len(s);i++{
		for{
			if _,ok:=set[s[i]];ok{
				delete(set,s[left])
				left++
			}else{
				break
			}
		}
		res=max(res,i-left+1)
		set[s[i]]= struct{}{}
	}
	return res
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}