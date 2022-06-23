package 后序遍历

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-20 15:21
 **/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) (res []int){
	var traversal func(node *TreeNode)
	traversal =func(node*TreeNode){
		if node==nil{
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res=append(res,node.Val)
	}
	return res
}