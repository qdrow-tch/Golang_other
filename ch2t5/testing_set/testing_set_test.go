package testing_set

import "testing"

func BenchmarkSetAdd(b *testing.B) {
	var set = NewSet()

	for i := 0; i < 100; i++ {
		go set.Add_val(1, 7.333)
	}

	for i := 0; i < 100; i++ {
		go set.Get_val(1)
	}
}
