package 中序遍历

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-20 14:47
 **/

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}
func preorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal=func(node *TreeNode){
		if node==nil{
			return
		}
		traversal(node.Left)
		res=append(res,node.Val)
		traversal(node.Right)
	}
	return res
}