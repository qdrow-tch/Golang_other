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
				fmt.Println("Счетчик: %d", cach)
				if cach == 1000 {
					break
				}
				cach++
				counter <- cach
			}
		}(i)
	}

	time.Sleep(6 * time.Second)
}
