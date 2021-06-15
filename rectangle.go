package main

import (
	"fmt"
)

func main() {
	var size_side_a int
	var size_side_b int

	fmt.Printf("Введите значение стороны A прямоугольника: ")
	fmt.Scanln(&size_side_a)
	fmt.Printf("Введите значение стороны B прямоугольника: ")
	fmt.Scanln(&size_side_b)

	fmt.Println("Площадь прямоугольника:", size_side_a*size_side_b)
}
