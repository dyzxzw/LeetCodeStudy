package LFU

/**
 * @Description

LFU 最不经常使用
实现 LFUCache 类：

LFUCache(int capacity) - 用数据结构的容量capacity 初始化对象
int get(int key)- 如果键key 存在于缓存中，则获取键的值，否则返回 -1 。
void put(int key, int value)- 如果键key 已存在，则变更其值；
如果键不存在，请插入键值对。当缓存达到其容量capacity 时，则应该在插入新项之前，移除最不经常使用的项。
在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。
对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

 * @Author ZzzWw
 * @Date 2022-05-23 13:42
 **/

/*
双哈希表 +双向链表
一个哈希表指向每个结点
一个哈希表指向每个链表 ，关键字为频次
删除最小访问的链表  的 尾部结点 （频次相同，删除最近未使用）
 */

type Node struct {
	key,val ,freq int
	prev,next *Node
}

func initNode(key,val int)*Node{
	return &Node{
		key: key,
		val: val,
		freq: 1, //初始频次为1
	}
}
//双向链表

type DList struct {
	head,tail *Node
}
func initDList()*DList{
	l:=&DList{
		head:&Node{},
		tail:&Node{},
	}
	l.head.next=l.tail
	l.tail.prev=l.head
	return l
}

type LFUCache struct {
    size,cap,minFreq int//每次从最小访问频次的链表删除，每次删除链表尾结点
	cacheNode map[int]*Node
	cacheDList map[int]*DList
}

func Constructor(cap int)LFUCache{
	l:=LFUCache{
		size: 0,
		cap: cap,
		minFreq: 0,
		cacheNode: map[int]*Node{},
		cacheDList: map[int]*DList{},
	}
	return l
}

//查找
func (this *LFUCache) Get(key int) int{
	if node,ok:=this.cacheNode[key];ok{
		this.IncFreq(node)
		// IncFreq 增加访问频率:
		//1.首先，把结点从原来的map和链表删除
		//2.然后加入下一个频率的map和链表,同时更新minFreq
		return node.val
	}
	return -1
}

//修改
func (this *LFUCache) Put(key, val int) {
	//如果不存在链表
	if this.cap==0{
		return
	}
	//如果存在结点
	if node,ok:=this.cacheNode[key];ok{
		node.val=val
		this.IncFreq(node)
	}else{
		//如果不存在
		//1.先判断容量是否超标
		if this.size==this.cap{
			//找到频次最小的链表，删除尾部结点
			rn:=this.cacheDList[this.minFreq].RemoveTail()//从链表map中删除
			delete(this.cacheNode,rn.key)//从结点map中删除
			this.size--
		}
		//2.创造一个新结点
		newNode:=initNode(key,val)
		//2.1添加到结点map
		this.cacheNode[key]=newNode
		//2.2添加到链表map
		//如果访问频次为1的链表为空了，创造一个
		if this.cacheDList[1]==nil{
			this.cacheDList[1]=initDList()
		}
		//把这个结点添加到头部
		this.cacheDList[1].AddFirst(newNode)
		this.minFreq=1
		this.size++
	}
}

// IncFreq 增加访问频率，
//1.首先，把结点从原来的map和链表删除
//2.然后加入下一个频率的map和链表,同时更新minFreq
func (this *LFUCache) IncFreq(node *Node){
	//从当前访问频次的链表删除结点
	this.cacheDList[node.freq].removeNode(node)
	//如果当前的结点已经是最后一个结点，且是最小频次，删除这个结点后，才要将最小访问频次+1更新
	//同时，删除这个链表
	if node.freq==this.minFreq && this.cacheDList[this.minFreq].IsEmpty(){
		this.minFreq++
		delete(this.cacheDList,node.freq)
	}
	//访问频次+1
	node.freq++
	if this.cacheDList[node.freq]==nil{
		this.cacheDList[node.freq]=initDList()
	}
	this.cacheDList[node.freq].AddFirst(node)
}

// 移除节点

func (this *DList) removeNode(node *Node){
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}

// AddFirst 将节点添加到头部表示最近访问
func (this *DList) AddFirst(node *Node) {
	node.next = this.head.next
	node.prev = this.head

	this.head.next.prev = node
	this.head.next = node
}

// 移除最后的节点，同频率下末尾节点为最久未访问
func (this *DList) RemoveTail() *Node {
	if this.IsEmpty() {
		return nil
	}
	last := this.tail.prev
	this.removeNode(last)
	return last
}

func (this *DList) IsEmpty() bool {
	return this.head.next == this.tail
}