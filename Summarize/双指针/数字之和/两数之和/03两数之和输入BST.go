package 两数之和

/**
 * @Description
给定一个二叉搜索树 root 和一个目标结果 k
如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

 * @Author ZzzWw
 * @Date 2022-06-20 14:37
 **/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func findTarget(root *TreeNode, k int) bool {
	set := map[int]struct{}{}
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := set[k-node.Val]; ok {
			return true
		}
		set[node.Val] = struct{}{}
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}