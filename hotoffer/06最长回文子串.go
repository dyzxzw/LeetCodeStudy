package hotoffer

import "fmt"

/**
 * @Description
给你一个字符串 s，找到 s 中最长的回文子串。或输出回文子串的个数
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
 * @Author ZzzWw
 * @Date 2022-07-29 22:03
 **/
func longestPalindrome(s string) string {
	maxStr,left,right,res:=0,0,0,0
	dp:=make([][]bool,len(s))
	for i:=range dp{
		dp[i]=make([]bool,len(s))
	}
	for i:=len(s)-1;i>=0;i--{
		for j:=i;j<len(s);j++{
			if s[i]==s[j]{
				if j-i<=1{
					dp[i][j]=true
					res++
				}else if dp[i+1][j-1]{
					dp[i][j]=true
					res++
				}
			}
			if dp[i][j] && j-i+1>maxStr{
				maxStr,left,right=j-i+1,i,j
			}
		}
	}
	fmt.Println("回文字串的个数为：",res)
	return s[left:right+1]
}