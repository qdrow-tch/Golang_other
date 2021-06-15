package main

import (
	"fmt"
	"math"
)

func main() {
	var circle_place float64
	var circle_dia float64

	fmt.Printf("Введите площадь круга: ")
	fmt.Scanln(&circle_place)
	circle_dia = math.Sqrt(circle_place/math.Pi) * 2

	fmt.Println("Диаметр круга:", circle_dia, "\nДлина окружности:", circle_dia*math.Pi)

}
