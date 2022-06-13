package 前序遍历

import "container/list"

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-13 19:50
 **/

func preorderTraversal2(root *TreeNode) []int {
	ans:=[]int{}
	if root==nil{
		return nil
	}
	st:=list.New()
	st.PushBack(root)

	for st.Len()>0{
		node:=st.Remove(st.Back()).(*TreeNode)
		ans=append(ans,node.Val)
		if node.Right!=nil{
			st.PushBack(node.Right)
		}
		if node.Left!=nil{
			st.PushBack(node.Left)
		}
	}
	return ans
}