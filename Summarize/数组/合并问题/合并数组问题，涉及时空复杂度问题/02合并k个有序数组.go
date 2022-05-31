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
  用归并排合并 或者优先队列（堆）解决
 * @Author ZzzWw
 * @Date 2022-05-25 15:03
 **/


//记住堆的模板

type IntHeap [][3]int

//Len Less Swap Push Pop
func(h IntHeap)Len() int    { return len(h)}
func(h IntHeap)Less(i,j int)bool { return h[i][0]<h[j][0] }
func(h IntHeap)Swap(i,j int) { h[i],h[j]=h[j],h[i]}
func(h *IntHeap)Push(x interface{}) { *h=append(*h,x.([3]int) )  }
func(h*IntHeap)Pop()interface{}{
	old:=*h
	n:=len(old)
	x:=old[n-1]
	*h=old[:n-1]
	return x
}

//归并排序+优先队列（堆）
// 堆的 元素 :0存放值, 1存行下标, 2存列下标
func mergeKNums(lists [][]int) []int{
	//初始化堆
	h:=&IntHeap{}
	heap.Init(h)
	//如果lists为Nil
	if lists==nil{
		return nil
	}
	//定义排序后的结构
	res:=make([]int,0,len(lists)*len(lists[0]))
	//构造堆
	for i:=0;i<len(lists);i++{
		heap.Push(h,[3]int{ lists[i][0],i,0} )
	}
	//
	for j:=0;j<len(lists)*len(lists[0]);j++{
		now:=heap.Pop(h).([3]int)
		res=append(res,now[0])
		if now[2]!=len(lists[0])-1{
			heap.Push(h,[3]int{ lists[now[1]][now[2]+1],now[1],now[2] +1 } )
		}
	}
	return res
}

func  mergeKNumsTest(){
	input := bufio.NewScanner(os.Stdin)
	input.Scan() //读取一行m x n
	m,_:=strconv.Atoi(strings.Split(input.Text(),",")[0])
	n,_:=strconv.Atoi(strings.Split(input.Text(),",")[1])

	lists := make([][]int, m)
	for i := 0; i < m; i++ {
		lists[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		input.Scan() //读取一行
		for j := 0; j < n; j++ {
			lists[i][j], _ = strconv.Atoi(strings.Split(input.Text(), ",")[j])
		}
	}
	res :=mergeKNums(lists)
	fmt.Println(res)
}