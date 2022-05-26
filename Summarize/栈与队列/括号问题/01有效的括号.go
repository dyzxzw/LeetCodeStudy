package 括号问题

/**
 * @Description
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

输入：s = "()[]{}"
输出：true

输入：s = "(]"
输出：false
 * @Author 赵稳
 * @Date 2022-05-25 15:43
 **/
func isValid(s string) bool {
	hash := map[byte]byte{')':'(', ']':'[', '}':'{'}
	stack := make([]byte, 0)
	if s == "" {
		return true
	}

	for i := 0; i < len(s); i++ {
		//如果全是左符号，则入栈
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			//如果栈不为空，且栈顶元素==入栈元素所应对的哈希表，则出栈
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
/*
拓展1：如果不要求按顺序，比如([)]也可以怎么办：分别用n1，n2，n3记录
左边括号的数量时刻》=右边括号的数量
*/
func isValid1(s string )bool{
       n1,n2,n3:=0,0,0
	   for i:=0;i<len(s);i++{
		   if s[i]=='('{
			   n1++
		   }else if s[i]==')'{
			   n1--
		   }
		   if s[i]=='['{
			   n2++
		   }else if s[i]==']'{
			   n2--
		   }
		   if s[i]=='{'{
			   n3++
		   }else if s[i]=='}'{
			   n3--
		   }
		   if n1==-1||n2==-1||n3==-1{
			   return false
		   }
	   }
	   return true
}

/*
拓展2：要求必须按 { [ ( 的顺序进行关闭，可以有连续的左括号或者右括号。
意思就是说大括号里面必须含有小括号，小括号里面不能含有大括号

思想：设置优先级，假设三个优先级
出栈的时刻记录一下当前优先级，下次出栈的时候的优先级必须大于上次的优先级
更新优先级
*/
func isValid2(s string )bool{
	//设置三个优先级以及当前临时优先级
	x1,x2,x3,x:=1,2,3,0
	hash := map[byte]byte{')':'(', ']':'[', '}':'{'}
	stack := make([]byte, 0)
	if s == "" {
		return true
	}
	for i := 0; i < len(s); i++ {
		//如果全是左符号，则入栈
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			//如果栈不为空，且栈顶元素==入栈元素所应对的哈希表，则出栈,记录当前优先级
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
			if x!=0{
				if hash[s[i]]=='('{
					//如果当前优先级小于上次优先级
					if x1<x{
						return false
					}
					//更新优先级
					x=x1
				}else if hash[s[i]]=='['{
					//如果当前优先级小于上次优先级
					if x2<x{
						return false
					}
					//更新优先级
					x=x2
				}else{
					//如果当前优先级小于上次优先级
					if x3<x{
						return false
					}
					//更新优先级
					x=x3
				}
			}else{
				if hash[s[i]]=='('{
					x=x1
				}else if hash[s[i]]=='['{
					x=x2
				}else{
					x=x3
				}
			}
			stack = stack[:len(stack)-1]

		} else {
			return false
		}
	}
	return len(stack) == 0
}