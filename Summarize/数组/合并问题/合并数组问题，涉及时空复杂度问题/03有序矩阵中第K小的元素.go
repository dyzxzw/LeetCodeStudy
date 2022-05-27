package 合并数组问题_涉及时空复杂度问题

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * @Description
给你一个n x n矩阵matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。

输入：matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
输出：13
解释：矩阵中的元素为 [1,5,9,10,11,12,13,13,15]，第 8 小元素是 13

 * @Author 赵稳
 * @Date 2022-05-27 13:46
 **/

//堆模板

type IHeap [][3]int

//Len Less Swap Push Pop
func(h IHeap)Len()int{ return len(h) }
func(h IHeap)Less(i,j int)bool { return h[i][0]<h[j][0] }
func(h IHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func(h *IHeap)Push(x interface{}){*h=append(*h,x.([3]int))}
func(h *IHeap)Pop()interface{}{
	old:=*h
	n:=len(old)
	x:=old[n-1]
	*h=old[:n-1]
	return x
}
func kthSmallest(matrix [][]int, k int) int{
	// 归并排序后用小顶堆取出k - 1次值, 然后返回堆顶元素
	//1.初始化堆
	h:=&IHeap{}
	heap.Init(h)
	//2.建堆
	for i:=0;i<len(matrix);i++{
		heap.Push(h,[3]int{matrix[i][0],i,0})
	}
	for j:=0;j<k-1;j++{
		now:=heap.Pop(h).([3]int)
		if now[2]!=len(matrix)-1{
			heap.Push(h,[3]int{ matrix[now[1]][now[2]+1],now[1] ,now[2]+1 }     )
		}
	}
	return (*h)[0][0]
}
func  kthSmallestTest(){
    input:=bufio.NewScanner(os.Stdin)
	input.Scan()//读取一行NxN
	n,_:=strconv.Atoi(input.Text())
	nums:=make([][]int,n)
	for i:=0;i<n;i++{
		nums[i]=make([]int,n)
	}
	for i:=0;i<n;i++{
		input.Scan()//读取一行数据
		for j:=0;j<n;j++{
			nums[i][j],_=strconv.Atoi(strings.Split(input.Text(),",")[j])
		}
	}
	input.Scan()//读取k
	k,_:=strconv.Atoi(input.Text())
	//输出结果
	res:=kthSmallest(nums,k)
	fmt.Println(res)
}