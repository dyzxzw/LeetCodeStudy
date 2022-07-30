package hotoffer

/**
 * @Description
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 * @Author ZzzWw

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
 * @Date 2022-07-30 11:07
 **/

func backtracking(n int,res*[]string,left,right int,str string){
	if right>left{
		return
	}
	if left==right&& left==n{
		*res=append(*res,str)
		return
	}
	if left<n{
		backtracking(n,res,left+1,right,str+"(")
	}
	if right<n{
		backtracking(n,res,left,right+1,str+")")
	}
}
func generateParenthesis(n int) []string {
	res:=[]string{}
	backtracking(n,&res,0,0,"")

	return res
}