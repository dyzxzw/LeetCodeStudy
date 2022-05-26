package main

import "fmt"

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
拓展1：如果不要求按顺序，比如([)]也可以怎么办：分别用n1，n2，n3记录，
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
func isValid2(s string )bool{
	//设置三个优先级以及临时优先级
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
					if x1<x{
						return false
					}
					//更新优先级
					x=x1
				}else if hash[s[i]]=='['{
					if x2<x{
						return false
					}
					//更新优先级
					x=x2
				}else{
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
func main(){
	str:="{{[[(())]]}}"
	fmt.Println(isValid2(str))
}