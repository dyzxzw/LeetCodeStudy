package Map

import "sync"

/**
 * @Description
 * @Author ZzzWw
 * @Date 2022-06-14 14:16
 **/


type RWMap struct {
	m map[interface{}]interface{}
	sync.RWMutex
}

// NewRWMap 新建一个RWMap
func NewRWMap(n int)*RWMap{
	return &RWMap{
		m:make(map[interface{}]interface{},n),
	}
}
func(this *RWMap)Get(k interface{})(interface{},bool){
	//加锁
	this.RLocker()
	defer this.RUnlock()
	v,existed:=this.m[k]
	return v,existed
}
func(this*RWMap)Set(k,v interface{}){
	this.Lock()
	defer this.Unlock()
	this.m[k]=v
}
func(this*RWMap)Delete(k interface{}){
	this.Lock()
	defer this.Unlock()
	delete(this.m,k)
}