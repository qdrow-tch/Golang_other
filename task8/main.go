package main

import (
	"fmt"
	"foo"
)

func main() {

	value := foo.FileConfig()

	fmt.Println(value)
}
