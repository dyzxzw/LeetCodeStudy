package 括号问题

/**
 * @Description
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"

 * @Author 赵稳
 * @Date 2022-05-27 21:38
 **/

/*
利用标记数组mark 记录连续0的个数，这个数量就是最长有效括号的长度

s = ")()())"

mark 100001


*/
func longestValidParentheses(s string) int {
         st:=make([]int,0)  //记录每个括号的下标栈，如果是左括号则入栈，右括号则出栈
		 mark:=make([]int,len(s)) //标记数组，初始全为0
		 for i:=0;i<len(s);i++{
			 if s[i]=='('{
				 st=append(st,i)
			 }else{
				 if len(st)==0{
					 mark[i]=1 //如果栈为空，且为右括号，则直接在mark数组对应位置标记为1
				 }else{
					 st=st[:len(st)-1] //左右括号匹配上了，出栈即可
				 }
			 }
		 }
	//如果栈中还存有没有匹配完的，说明剩下的括号都是无法匹配的
	//直接在对应mark的位置标记为1
	for len(st)>0{
		mark[st[len(st)-1]]=1
		st=st[:len(st)-1] //出栈
	}
	//统计连续0的个数
	cnt,maxCnt:=0,0
	for _,v:=range mark{
		if v==1{
			cnt=0
			continue
		}
		cnt++
		maxCnt=max(maxCnt,cnt)
	}
	return maxCnt
}

func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}