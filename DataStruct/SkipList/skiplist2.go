package SkipList

import (
	"bytes"
	"math"
	"math/rand"
	"time"
)

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-10 21:08
 **/

const(
	MaxLevel int = 16 //设置最大为18层，可以视情况调节
	probability float64 =1/math.E //0.36787944117144233
)

type Element struct {
	Node
	key []byte
	value interface{}
}
type Node struct {
	next []*Element
}

type SkipList struct {
	Node
	maxLevel int
	Len int
	randSource rand.Source
	probability    float64
	probTable      []float64
	prevNodesCache []*Node
}

// NewSkipList create a new skip list.
func NewSkipList()*SkipList{
	return &SkipList{
		Node:  Node{next: make([]*Element, MaxLevel)},
		prevNodesCache: make([]*Node,MaxLevel),
		maxLevel:       MaxLevel,
		randSource:     rand.New(rand.NewSource(time.Now().UnixNano())),
		probability:    probability,
		probTable:      probabilityTable(probability,  MaxLevel),

	}
}

// Key 获得element的key
func(e *Element)Key()[]byte{
	return e.key
}

// Value 获得element的value
func(e *Element)Value()interface{}{
	return e.value
}

// SetValue 设置element的value
func(e*Element)SetValue(val interface{}){
	e.value=val
}

//Next
func(e*Element)Next()*Element{
	return e.next[0]
}
//Front
func(t *SkipList)Front()*Element{
	return t.next[0]
}

func (t *SkipList) Get(key []byte) *Element{
	var prev=&t.Node
	var next *Element

	for i:=t.maxLevel-1;i>=0;i--{
		next=prev.next[i]

		for next!=nil &&bytes.Compare(key,next.key)>0{
			prev=&next.Node
			next=next.next[i]
		}
	}

	if next!=nil && bytes.Compare(next.key,key)<=0{
		return next
	}
	return nil
}

//所以调用get函数，只要返回不为nil即表明目标存在跳表中
// Exist check if exists the key in skl.

func (t *SkipList) Exist(key []byte) bool {
	return t.Get(key) != nil
}

