package 滑动窗口与高频元素

import "container/heap"

/**
 * @Description
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
 * @Author ZzzWw
 * @Date 2022-06-07 16:33
 **/

//构建小顶堆
//Len,Less,Swap,Push,Pop

type IHeap [][2]int
func(h IHeap)Len()int{return len(h)}
func(h IHeap)Less(i,j int)bool{return h[i][1]<h[j][1]}
func(h IHeap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func(h*IHeap)Push(x interface{}){ *h=append(*h,x.([2]int))}
func(h*IHeap)Pop()interface{}{
	old:=*h
	n:=len(old)
	x:=old[n-1]
	*h=old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
      mapNum :=map[int]int{}
	  //记录每个元素出现的次数
	  for _,item:=range nums{
		  mapNum[item]++
	  }
	  h:=&IHeap{}
	  heap.Init(h)
	  //入堆
	  for key,value:=range mapNum {
		  heap.Push(h,[2]int{key,value})
		  if h.Len()>k{
            heap.Pop(h)
		  }
	  }
	  res:=make([]int,k)
	  for i:=0;i<k;i++{
		  res[k-i-1]=heap.Pop(h).([2]int)[0]
	  }
	  return res
}