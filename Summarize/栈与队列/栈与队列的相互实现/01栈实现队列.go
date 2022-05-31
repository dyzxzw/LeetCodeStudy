package 栈与队列的相互实现

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-05-24 13:48
 **/


type MyQueue struct {
       stackIn []int
	   stackOut []int
}


func Constructor() MyQueue {
    return MyQueue{
		stackIn: make([]int,0),
		stackOut: make([]int,0),
	}
}


func (this *MyQueue) Push(x int)  {
    this.stackIn=append(this.stackIn,x)
}


func (this *MyQueue) Pop() int {
	// 只有当stOut为空的时候，再从stIn里导入数据（导入stIn全部数据）
	if len(this.stackOut)==0{
		for len(this.stackIn)>0{
			this.stackOut=append(this.stackOut,this.stackIn[len(this.stackIn)-1])
			this.stackIn=this.stackIn[:len(this.stackIn)-1]
		}
	}
	res:=this.stackOut[len(this.stackOut)-1]
	this.stackOut=this.stackOut[:len(this.stackOut)-1]
	return res
}


func (this *MyQueue) Peek() int {
    res:=this.Pop()
	this.stackOut=append(this.stackOut,res)
	return res
}


func (this *MyQueue) Empty() bool {
     if len(this.stackOut)==0 && len(this.stackIn)==0{
		 return true
	 }
	 return false
}

func main() {

}
