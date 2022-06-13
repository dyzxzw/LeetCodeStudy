package SkipList

import (
	"math/rand"
	"strings"
	"time"
)

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

func init(){
	rand.Seed(time.Now().UnixNano())
}

type skipListNode struct {
	key string
	value interface{}
	next []*skipListNode
}

type SkipList struct {
	head *skipListNode //头结点
	maxLevel int //最大层数
	length int //长度
}

func NewSkipList()*SkipList{
	return &SkipList{
		head:&skipListNode{
			next: make([]*skipListNode,1),
		},
		maxLevel: 1,
		length: 0,
	}
}

//Search is
func (s *SkipList)Search(key string)interface{}{
	current := s.head // 从头节点开始
	for i := s.maxLevel - 1; i >= 0; i-- {
		// 找到当前level最后一个小于key的节点
		for current.next[i] != nil && strings.Compare(current.next[i].key ,key)>0{
			current = current.next[i]
		}
		// 进入下一个level寻找
	}
	// 如果找到了就返回value
	current = current.next[0] // 大于等于key的第一个节点，只会在level为0的时候找到
	if current != nil && strings.Compare(current.key , key)==0 {
		return current.value
	}
	return -1
}