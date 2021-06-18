package main

import (
	"fmt"
)

func main() {

	data := []int{7, 3, 4, 5, 12, 8, 10, 3, 5, 2}

	sorting(data)

	fmt.Println(data)

}

func sorting(data []int) {
	cach := 0
	for i := 0; i < len(data); i++ {
		fmt.Println("i: ", i)
		for j := 0; j < i; j++ {
			fmt.Println("j: ", j)
			if data[i] < data[j] {
				cach = data[i]
				for mv := i; mv > j; mv-- {
					data[mv] = data[mv-1]
				}
				data[j] = cach
			}
		}
	}
}
