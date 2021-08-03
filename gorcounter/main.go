package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func counWorker(wg *sync.WaitGroup, id int, ctx context.Context, counter chan int) {

	defer func() {
		wg.Done()
		fmt.Println("Stop")
	}()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			cach, ok := <-counter
			if cach >= 1000 {
				close(counter)
				return
			}
			if ok {
				cach++
				counter <- cach
				fmt.Println("Worker ", id, " push value: ", cach)
			} else {
				return
			}
		}
	}
}

func main() {

	wg := &sync.WaitGroup{}

	counter := make(chan int, 1)
	counter <- 0

	ctx := context.Background()
	cancelableCtx, cancelfunc := context.WithTimeout(ctx, 1*time.Second)

	sigchecker := make(chan os.Signal, 2)
	signal.Notify(sigchecker, syscall.SIGTERM)

	go func() {
		gotsign := <-sigchecker
		fmt.Println(gotsign)
		func(context.Context) {

		}(cancelableCtx)
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counWorker(wg, i, ctx, counter)
	}

	wg.Wait()
	fmt.Println("All gorutinse was stoped")
	cancelfunc()
}
