package 设计计算器相关

/**
 * @Description
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。

注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
s 由数字、'+'、'-'、'*', '/'、('、')'、和 ' ' 组成

输入：s = " 2-1 + 2 "
输出：3

 * @Author ZzzWw
 * @Date 2022-05-31 16:00
 **/

func calculate(str string)int{
	i:=0
	return solution(&str,&i)
}
func solution(s *string, index *int)  int {
	//定义栈
	var stack []int
	//记录算式中的数字
	var num int =0
	//记录num前的符号，初始化为+
	var sign byte ='+'
	n:=len(*s)
	for ;*index<n;*index++{
		c:=(*s)[*index]
		//如果读取到的是数字，就加到num中
		if isDigit(c){
			num=10*num+int(c-'0')
		}
		//如果是左括号
		if c == '(' {
			*index++
			num = solution(s, index)
		}
		//如果不是数字，就要把之前的数字存入栈中；并更新符号
		if (!isDigit(c) && c!=' ') || *index==n-1{
			var pre int =0
			switch sign {

			case '+': stack=append(stack,num)
			case '-': stack=append(stack,-num)
			//注意 * /的运算
			case '*':
				pre=stack[len(stack)-1]
				stack=stack[:len(stack)-1]
				stack=append(stack,pre*num)

			case '/':
				pre=stack[len(stack)-1]
				stack=stack[:len(stack)-1]
				stack=append(stack,pre/num)
			}
			//更新符号为当前符号，数字清零
			sign=c
			num=0
		}
		//如果遇到右括号
		if c==')'{
			break
		}
	}
	//将栈中所有结果求和就是答案
	var res int=0
	for _,v:=range stack{
		res+=v
	}
	return res
}
func isDigit( x byte)bool{
	if x>='0'&&x<='9'{
		return true
	}
	return false
}
