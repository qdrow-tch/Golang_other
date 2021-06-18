package main

import (
	"fmt"
)

func main() {

	data := []int{7, 3, 4, 5, 12, 8, 10, 3, 5, 2}

	fmt.Println(sorting(data))

}

func sorting(data []int) []int {
	cach := 0
	slice_cach := make([]int, len(data))
	copy(slice_cach, data)
	for i := 0; i < len(slice_cach); i++ {
		for j := 0; j < i; j++ {
			if slice_cach[i] < slice_cach[j] {
				cach = slice_cach[i]
				for mv := i; mv > j; mv-- {
					slice_cach[mv] = slice_cach[mv-1]
				}
				slice_cach[j] = cach
			}
		}
	}
	return slice_cach
}
