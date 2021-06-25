package main

import (
	"fmt"
)

func fibon(value int) int {
	if value == 0 {
		return 0
	}
	if value == 1 {
		return 1
	}
	return fibon(value-2) + fibon(value-1)
}
func main() {

	var count int
	fmt.Println("Введите количество чисел Фибоначчи: ")
	fmt.Scanln(&count)

	for i := 0; i < count; i++ {
		fmt.Println(fibon(i))
	}

	fmt.Println("End")
}
