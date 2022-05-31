package 设计计算器相关

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * @Description
根据 逆波兰表示法，求表达式的值。

有效的算符包括+、-、*、/。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

注意两个整数之间的除法只保留整数部分。

可以保证给定的逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。


 * @Author ZzzWw
 * @Date 2022-05-30 14:33
 **/

func evalRPN(tokens []string) int {
	//定义栈,遇到数字就存放，遇到+ - * /就弹出元素进行计算
	stack:=[]int{}
	for i:=0;i<len(tokens);i++{
		 val,err:=strconv.Atoi(tokens[i])
		 if err==nil {
			 //如果是数字，进栈
			 stack=append(stack,val)
		 }else{
			 num1,num2:=stack[len(stack)-2],stack[len(stack)-1]
			 stack=stack[:len(stack)-2]
			 switch tokens[i]{
			 case "+":
				 stack=append(stack,num1+num2)
			 case "-":
				 stack=append(stack,num1-num2)
			 case "*":
				 stack=append(stack,num1*num2)
			 case "/":
				 stack=append(stack,num1/num2)
			 }
		 }
	}
  return stack[0]
}

func evalRPNTest(){
	input := bufio.NewScanner(os.Stdin)
	input.Scan() //读取tokens
	var tokens []string
	str:=strings.Split(input.Text(),",")
	for _,v:=range str{
		tokens=append(tokens,v)
	}
	res := evalRPN(tokens)
	fmt.Println(res)
}