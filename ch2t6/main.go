package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkg/profile"

	"github.com/qdrow-tch/Golang/tree/main/ch2t6/counterMTX"
)

func main() {

	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	count_ctx, cancelFunc := context.WithCancel(context.Background())
	counter := counterMTX.NewCounter(1000, cancelFunc)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-count_ctx.Done():
					fmt.Println("Горутина № ", id, " завершила работу")
					return
				default:
					counter.Add()
				}
			}
		}(i)
	}
	wg.Wait()

}
