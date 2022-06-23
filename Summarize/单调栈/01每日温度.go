package 单调栈

import "container/list"

/**
 * @Description
给定一个整数数组temperatures，表示每天的温度，
返回一个数组answer，其中answer[i]是指在第 i 天之后，
才会有更高的温度。如果气温在这之后都不会升高，请在该位置用0 来代替。

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]

 * @Author ZzzWw
 * @Date 2022-06-21 14:18
 **/

func dailyTemperatures(temperatures []int) []int {
      //链表模拟栈
	 st:=list.New()
	 st.PushBack(0)
	 //存放结果，初始化为0
	 res:=make([]int,len(temperatures))
	 for i:=1;i<len(temperatures);i++{
		 if temperatures[i]<temperatures[st.Back().Value.(int)]{
			 st.PushBack(i)
		 } else if temperatures[i] == temperatures[st.Back().Value.(int)] {
			 st.PushBack(i)
		 }else{
			 for st.Len()>0 && temperatures[i]>temperatures[st.Back().Value.(int)]{
				 res[st.Back().Value.(int)]=i-st.Back().Value.(int)
				 st.Remove(st.Back())
			 }
			 st.PushBack(i)
		 }
	 }
	return res
}