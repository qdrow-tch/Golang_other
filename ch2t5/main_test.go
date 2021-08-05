package main

import (
	"testing"
)

func BenchmarkMutex_10_90(b *testing.B) {
	set := NewSet(10)
	for i := 0; i < b.N; i++ {
		go set.Add_Mtx(i, 1.5)
	}
}

func BenchmarkMutex_90_10(b *testing.B) {

	set := NewSet(90)
	for i := 0; i < b.N; i++ {
		go set.Add_Mtx(i, 1.5)
	}
}
func BenchmarkMutex_50_50(b *testing.B) {
	set := NewSet(50)
	for i := 0; i < b.N; i++ {
		go set.Add_Mtx(i, 1.5)
	}

}

func BenchmarkRWMutex_10_90(b *testing.B) {

	set := NewSet_RW(10)
	for i := 0; i < b.N; i++ {
		go set.Add_RWMtx(i, 1.5)
	}

}
func BenchmarkRWMutex_90_10(b *testing.B) {

	set := NewSet_RW(90)
	for i := 0; i < b.N; i++ {
		go set.Add_RWMtx(i, 1.5)
	}

}
func BenchmarkRWMutex_50_50(b *testing.B) {

	set := NewSet_RW(50)
	for i := 0; i < b.N; i++ {
		go set.Add_RWMtx(i, 1.5)
	}

}
