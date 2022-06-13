package SkipList

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-10 21:08
 **/


// 跳表
func init() {
	rand.Seed(time.Now().UnixNano())
}

type SkipList2 struct {
	head     *SkipListNode // 头节点
	maxLevel int           // 最大层数
	length   int           // 长度
}

type SkipListNode struct {
	key   int
	value int
	next  []*SkipListNode // 每个level的下一个节点
}

func NewSkipList2() *SkipList {
	return &SkipList{
		head: &SkipListNode{
			next: make([]*SkipListNode, 1),
		},
		maxLevel: 1,
		length:   0,
	}
}

func (s *SkipList) Insert(key int, value int) {
	update := make([]*SkipListNode, s.maxLevel)
	current := s.head
	for i := s.maxLevel - 1; i >= 0; i-- {
		// 找到当前level最后一个小于key的节点
		for current.next[i] != nil && current.next[i].key < key {
			current = current.next[i]
		}
		update[i] = current
	}
	// 如果插入的key已经存在，则更新value，并返回
	current = current.next[0]
	if current != nil && current.key == key {
		current.value = value
		return
	}
	// 随机插入节点的level
	level := randomLevel(s.maxLevel)
	// 如果level大于当前跳表的最大level，则更新最大level
	if level > s.maxLevel {
		for i := s.maxLevel; i < level; i++ {
			update = append(update, s.head)
		}
		s.maxLevel++
		s.head.next = append(s.head.next, nil)
	}
	// 新建节点
	node := &SkipListNode{
		key:   key,
		value: value,
		next:  make([]*SkipListNode, level),
	}
	// 插入节点
	for i := 0; i < level; i++ {
		if i < len(update[i].next) {
			node.next[i] = update[i].next[i]
			update[i].next[i] = node
		}
	}
	s.length++
}

func randomLevel(l int) int {
	level := 1
	for rand.Intn(2) == 0 {
		level++
		if level >= l+1 {
			break
		}
	}
	return level
}

func (s *SkipList) Delete(key int) {
	update := make([]*SkipListNode, s.maxLevel)
	current := s.head
	// 找到要删除节点每个level下的前一个节点，
	for i := s.maxLevel - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key < key {
			current = current.next[i]
		}
		update[i] = current
	}
	current = current.next[0]
	// 如果找到了就删除
	if current != nil && current.key == key {
		for i := 0; i < s.maxLevel; i++ {
			if update[i].next[i] != current {
				break
			}
			// 直接将当前level的前一个节点指向下一个节点
			update[i].next[i] = current.next[i]
		}
		// 更新最大level
		for s.maxLevel > 1 && s.head.next[s.maxLevel-1] == nil {
			s.maxLevel--
		}
		s.length--
	}
}

func (s *SkipList) Search(key int) int {
	current := s.head // 从头节点开始
	for i := s.maxLevel - 1; i >= 0; i-- {
		// 找到当前level最后一个小于key的节点
		for current.next[i] != nil && current.next[i].key < key {
			current = current.next[i]
		}
		// 进入下一个level寻找
	}
	// 如果找到了就返回value
	current = current.next[0] // 大于等于key的第一个节点，只会在level为0的时候找到
	if current != nil && current.key == key {
		return current.value
	}
	return -1
}

func (s *SkipList) Print() {
	for i := s.maxLevel - 1; i >= 0; i-- {
		current := s.head
		fmt.Printf("level %d: ", i)
		for current.next[i] != nil {
			current = current.next[i]
			fmt.Printf("%d ", current.key)
		}
		fmt.Println()
	}
}

func main() {
	sl := NewSkipList()
	for i := 0; i < 20; i++ {
		sl.Insert(i, i+100)
	}
	sl.Print()
	fmt.Printf("search: %d\n", sl.Search(5))
}

