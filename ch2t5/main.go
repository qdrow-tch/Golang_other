package main

//Залить в ветку для размещения на сайте

import "sync"

type Set struct {
	sync.Mutex
	mm map[int]float64
}

func NewSet(fullness int) *Set {
	mp := make(map[int]float64)
	for i := 0; i < fullness; i++ {
		mp[i] = 1.1
	}
	return &Set{
		mm: mp,
	}
}

func (s *Set) Add_Mtx(index int, value float64) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.mm[index]; ok {
		return
	}
	s.mm[index] = value
}

type Set_RW struct {
	sync.RWMutex
	mm map[int]float64
}

func NewSet_RW(fullness int) *Set_RW {
	mp := make(map[int]float64)
	for i := 0; i < fullness; i++ {
		mp[i] = 1.1
	}
	return &Set_RW{
		mm: mp,
	}
}

func (s *Set_RW) Add_RWMtx(index int, value float64) {
	s.RLock()
	if _, ok := s.mm[index]; ok {
		return
	}
	s.RUnlock()
	s.Lock()
	s.mm[index] = value
	s.Unlock()

}

func main() {

}
