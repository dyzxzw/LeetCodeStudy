package 括号问题

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * @Description
给定一个只包含三种字符的字符串：（，）和 *，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则：

1.任何左括号 (必须有相应的右括号 )
2.任何右括号 )必须有相应的左括号 (
3.左括号 ( 必须在对应的右括号之前 )
4.*可以被视为单个右括号 )，或单个左括号 (，或一个空字符串
5.一个空字符串也被视为有效字符串

输入: "(*)"
输出: True
示例 3:

输入: "(*))"
输出: True
 * @Author ZzzWw
 * @Date 2022-05-27 15:56
 **/

/*
两个栈解决
*/
func checkValidString(s string) bool{
       var leftStk []int     //存放左括号 下标 的栈
	   var asteriskStk []int //存放星号 下标 的栈
	   for i,ch:=range s{
		   if ch=='('{
			   leftStk=append(leftStk,i)
		   }else if ch=='*'{
			   asteriskStk =append(asteriskStk,i)
		   }else if len(leftStk)>0{ //遇到右括号，则保持左括号下标的栈 出栈
			   leftStk=leftStk[:len(leftStk)-1]
		   }else if len(asteriskStk)>0{ //如果没有左括号了，则到保持星号下标的栈中寻找
			   asteriskStk = asteriskStk[:len(asteriskStk)-1]
		   }else{
			   return false
		   }
	   }
	   //遍历结束之后，左括号栈和星号栈可能还有元素:要满足下列两个规则：
	   //为了将每个左括号匹配，需要将星号看成右括号
	   //且每个左括号必须出现在其匹配的星号之前
	   i:=len(leftStk)-1
	   for j:=len(asteriskStk)-1;i>=0&&j>=0;i,j=i-1,j-1{
		   if leftStk[i]>asteriskStk[j]{ //如果左括号出现在星号之后，则不能匹配 *（=）（
			   return false
		   }
	   }
        return i<0 //最后应该栈为空
}
func  checkValidStringTest(){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()
	str:=input.Text()
	fmt.Println(checkValidString(str))
}