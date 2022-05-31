package 栈与队列的相互实现

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-05-24 14:13
 **/

type MyStack struct {
  queue []int
}


func Constructor2() MyStack {
  return MyStack{
	  queue:make([]int,0),
  }
}


func (this *MyStack) Push(x int)  {
   this.queue=append(this.queue,x)
}


func (this *MyStack) Pop() int {
 n:=len(this.queue)-1
 for n!=0{
	 //除了最后一个，其余的都重新添加到队列里
	 val:=this.queue[0]
	 this.queue=this.queue[1:]
	 this.queue=append(this.queue,val)
	 n--
 }
 //弹出元素
 val:=this.queue[0]
 this.queue=this.queue[1:]
 return val
}


func (this *MyStack) Top() int {
    val:=this.Pop()
	this.queue=append(this.queue,val)
	return val
}


func (this *MyStack) Empty() bool {
     return len(this.queue)==0
}

func main() {

}
