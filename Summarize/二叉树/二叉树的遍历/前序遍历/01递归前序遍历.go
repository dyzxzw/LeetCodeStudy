package 前序遍历

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-13 16:10
 **/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func preorderTraversal(root *TreeNode) (res []int) {
    var traversal func(node*TreeNode)
	traversal=func(node *TreeNode){
		if node==nil{
			return
		}
		res=append(res,node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}