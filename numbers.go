package main

import (
	"fmt"
)

func main() {
	var value int64

	fmt.Printf("Введите трехзначное число: ")
	fmt.Scanln(&value)

	fmt.Println("Число состоит из: \n", (value / 100), "сотен\n",
		((value % 100) / 10), "десятков\n",
		((value % 100) % 10), "единиц")
}
