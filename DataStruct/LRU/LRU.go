package LRU

import "fmt"

/**
 * @Description

LRU:最近最久未使用
 哈希表 + 双向链表
 哈希表指向双向链表结点的一个结点
 每次删除的是链表尾部的结点
 使用过的结点都要放到链表头
 哈希表能快速找到要删除的结点

 * @Author ZzzWw
 * @Date 2022-05-23 13:42
 **/

type DLinkedNode struct {
	key,value int
	//思考一下，为什么链表结点也要有Key?
	//因为删除时候，方便cache找到key，从而删除结点
	prev,next*DLinkedNode
}

func initDLinkedNode(key,value int)*DLinkedNode{
	return &DLinkedNode{
		key: key,
		value: value,
	}
}

type LruCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head,tail *DLinkedNode
}

func Constructor(cap int) LruCache {
	l:= LruCache{
		size: 0,
		capacity: cap,
		cache: map[int]*DLinkedNode{},
		head: initDLinkedNode(0,0),
		tail: initDLinkedNode(0,0),
	}
	l.head.next=l.tail
	l.tail.prev=l.head
	return l
}

//查找

func(this *LruCache)Get(key int)int{
	if node,ok:=this.cache[key];!ok{
		return -1
	}else{
		this.moveToHead(node)
		return node.value
	}
}

//修改

func(this*LruCache)Put(key,val int){
	if node,ok:=this.cache[key];ok{
		node.value=val
		this.moveToHead(node)
	}else{
		//不存在
		//创造一个结点
		newNode:=initDLinkedNode(key,val)
		//先判断是否超过容量
		if this.size==this.capacity{
			removeNode:=this.removeTail()
			delete(this.cache,removeNode.key)
			this.size--
		}
		this.cache[key]=newNode
		this.addToHead(newNode)
		this.size++
	}
}

func(this*LruCache)moveToHead(node*DLinkedNode){
	this.removeNode(node)
	this.addToHead(node)
}

func(this*LruCache)removeNode(node *DLinkedNode){
	node.next.prev=node.prev
	node.prev.next=node.next
	node.next=nil
	node.prev=nil
}
func(this*LruCache)addToHead(node*DLinkedNode){
	node.next=this.head.next
	this.head.next.prev=node
	node.prev=this.head
	this.head.next=node
}

func(this*LruCache)removeTail()*DLinkedNode{
	rm:=this.tail.prev
	this.removeNode(rm)
	return rm
}

func LRUTest(){
	lruchache := Constructor(2)
	lruchache.Put(1, 1)
	lruchache.Put(2, 2)
	fmt.Println(lruchache.Get(1))
	lruchache.Put(3, 3)
	fmt.Println(lruchache.Get(2))
	lruchache.Put(4, 4)
	fmt.Println(lruchache.Get(1))
	fmt.Println(lruchache.Get(3))
	fmt.Println(lruchache.Get(4))
}