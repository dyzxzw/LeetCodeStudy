package 设计题

import "math"

/**
 * @Description
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。


 * @Author ZzzWw
 * @Date 2022-07-13 19:30
 **/

//要求算法的时间复杂度为O(1)

type MinStack struct {
  stack[] int
  MinStack []int
}


func Constructor() MinStack {
 return MinStack{
	 stack: []int{},
	 MinStack: []int{math.MaxInt64},
 }
}


func (this *MinStack) Push(val int)  {
        this.stack=append(this.stack,val)
		top:=this.MinStack[len(this.MinStack)-1]
		this.MinStack=append(this.MinStack,min(top,val))
}


func (this *MinStack) Pop()  {
     this.stack=this.stack[:len(this.stack)-1]
	 this.MinStack=this.MinStack[:len(this.MinStack)-1]
}


func (this *MinStack) Top() int {
 return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.MinStack[len(this.MinStack)-1]
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}