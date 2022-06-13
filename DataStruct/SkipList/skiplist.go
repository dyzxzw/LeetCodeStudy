package SkipList

/**
 * @Description

跃表是用于有序元素序列快速搜索查找的一个数据结构。
跳表是一个随机化的数据结构，实质就是一种可以进行二分查找的有序链表。

跳表的时间复杂度计算：
原始的链表的结点个数为n,称为第0层
第1层的结点个数:
假设我们建立索引是以每两个结点建立一个索引
第1层含有结点个数n/2

第2层的结点个数：
n/4 :n/2^2

第3层的结点个数:
n/8 n/2^3

第k层的结点个数：
n/2^k

而第k层的结点个数为n/2^k==2--->k=log2n-1

如果算上原始的第0层，k=long2n-1+1=long2n

假设每一层遍历的结点个数为m，则一共有k层，所以时间复杂度是m*k=m*long2n

简写为 O(log n)

空间复杂度为所有结点个数相加：
n/2+n/4+.....+2=n-2
简写为 O(n)


 * @Author ZzzWw
 * @Date 2022-06-10 14:38
 **/
/*
const(
	P = 0.25
	MaxLevel =16 //最大层数，视情况而定
)
func randomLevel()int{
	i:=1
	rand.Seed(time.Now().UnixNano())
    for i<MaxLevel{
		if rand.Float64()<P{
			i++
		}else{
			break
		}
	}
	return i
}
type Node struct {
	level int
	next []*Node
	v int
}
type skipList struct {
	head *Node
}
func NewSkipList()*skipList{
	s:=new(skipList)
	s.head=new(Node)
    s.head.level=MaxLevel
	s.head.next=make([]*Node,MaxLevel)
	return s
}

func(s*skipList)insert(v int){
	l:=randomLevel()
	add:=new(Node)
	add.level=l
	add.next=make([]*Node,l)
	add.v=v
	i:=l

	p:=s.head
	for i>0{
		n:=p.next[i-1]
		for n!=nil && n.v<v{
			p=n
			n=n.next[i-1]
		}
	}
}
*/