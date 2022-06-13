package 二叉树

import "fmt"

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-09 16:43
 **/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//数组构造二 叉树
func constructBinaryTree(nums []int)*TreeNode{
	if nums==nil{
		return nil
	}
	var root *TreeNode//根结点
	nodes:=make([]*TreeNode,len(nums))

	//初始化二 叉树结点
	for i:=0;i<len(nodes);i++{
		var node *TreeNode
		if nums[i]!=-1{
			node=&TreeNode{Val: nums[i]}
		}
		nodes[i]=node
		if i==0{
			root=node
		}
	}
	//串联结点
	for i:=0;i*2+2<len(nums);i++{
		if nodes[i]!=nil{
			nodes[i].Left=nodes[i*2+1]
			nodes[i].Right=nodes[i*2+2]
		}
	}
	return root
}

//层序打印二 叉树
func printBinaryTree(root *TreeNode,n int){
	var queue []*TreeNode
	if root!=nil{
		queue=append(queue,root)
	}
	res:=[]int{}
	for len(queue)>0{
		for j:=0;j<len(queue);j++{
			node:=queue[j]
			if node!=nil{
				res=append(res,node.Val)
				queue=append(queue,node.Left)
				queue=append(queue,node.Right)
			}else{
				res=append(res,-1)
			}
		}
		queue=queue[len(queue):]
	}
	fmt.Println(res[:n])
}