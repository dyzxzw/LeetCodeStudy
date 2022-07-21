package Set

import "sync"

/**
 * @Description
   并发安全访问的集合实现
 * @Author ZzzWw
 * @Date 2022-07-21 20:36
 **/

type Set struct {
	m map[interface{}]struct{}
	sync.RWMutex
}
func NewSet()*Set{
	return &Set{
		m: map[interface{}]struct{}{},
	}
}
func(s *Set)Add(item interface{}){
	//写锁
	s.Lock()
	defer s.Unlock()
	s.m[item]= struct{}{}
}
func(s *Set)Remove(item interface{}){
	s.Lock()
	defer s.Unlock()
	delete(s.m,item)
}
func(s *Set)Has(item interface{})bool{
	//读锁
	s.RLock()
	defer s.RUnlock()
	_,ok:=s.m[item]
	return ok
}
func(s *Set)SetPrint() []interface{}{
	s.RLock()
	defer s.RUnlock()
	var res []interface{}
	for value:=range s.m{
		res=append(res,value)
	}
	return res
}

func (s *Set) Len() int {
	return len(s.SetPrint())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[interface{}]struct{}{}
}

func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}