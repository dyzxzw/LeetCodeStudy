package 设计计算器相关

import (
	"unicode"
)

/**
 * @Description
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。

注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。

输入：s = " 2-1 + 2 "
输出：3

 * @Author ZzzWw
 * @Date 2022-05-31 16:00
 **/

func calculate(s string) int {
	index:=0
	var solution func(s string,index int) int
	solution =func(s string,index int) int{
		var stk []int
		num:=0
		sign:='+'
		for i,c:=range s{
			if unicode.IsDigit(c){
				num=10*num+ int(c-'0')
			}
			if c=='('{
				index++
				num=solution(s,index)
			}
			if (!unicode.IsDigit(c)&&c!=' ') || i==len(s)-1{
				switch sign {
				case '+':stk=append(stk,num)
				case '-':stk=append(stk,-num)
				}
				sign= c
				num=0
			}
			if c==')'{
				index++
				break
			}
		}
		res:=0
		for len(stk)>0{
			res+=stk[len(stk)-1]
			stk=stk[:len(stk)-1]
		}
		return res
	}
  return solution(s,index)
}