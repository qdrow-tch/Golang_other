package main

import (
	"fmt"

	"github.com/qdrow-tch/errors_workout/logging"
)

func main() {
	var value int = 100
	mass := []int{5, 15, 0, 8, 10}

	defer func() {
		if v := recover(); v != nil {
			logging.Push_to_logfile("Got panic plz chech your code")

		}
	}()

	for i, vl := range mass {
		fmt.Printf("%d: %d / %d = %f \n", i+1, value, vl, value/vl)
	}
}
