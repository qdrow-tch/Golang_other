package main

import (
	"fmt"
	"time"
)

func main() {

	counter := make(chan int, 1)
	counter <- 1

	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				cach := <-counter
				if cach == 1000 {
					break
				}
				cach++
				counter <- cach
				fmt.Println("Worker ", id, " push value: ", cach)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
}
