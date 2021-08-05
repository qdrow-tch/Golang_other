package main

import "sync"

type Set struct {
	sync.Mutex
	mm map[int]float64
}

func NewSet() *Set {
	return &Set{
		mm: make(map[int]float64),
	}
}

func (s *Set) Add_val(index int, value float64) {
	s.Lock()
	defer s.Unlock()
	s.mm[index] = value

}

func (s *Set) Get_val(index int) float64 {
	s.Lock()
	defer s.Unlock()
	return s.mm[index]
}

func main() {

}
