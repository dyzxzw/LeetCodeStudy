package 前序遍历

import "container/list"

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-13 19:50
 **/

/*
利用栈，每次先加入右孩子，然后再加入左孩子，保证
出栈的时候，是先出去左孩子 （根 左 右：前序）

*/
func preorderTraversal2(root *TreeNode) []int {
     ans:=[]int{}
	 if root==nil{
		 return ans
	 }
	 st:=list.New()
	 st.PushBack(root)

	 for st.Len()>0{
		 node:=st.Remove(st.Back()).(*TreeNode)

		 ans=append(ans,node.Val)
		 //先加入右
		 if node.Right!=nil{
			 st.PushBack(node.Right)
		 }
		 if node.Left!=nil{
			 st.PushBack(node.Left)
		 }
	 }
	 return ans
}