package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	counter := make(chan int, 1)
	counter <- 1

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

	for i := 0; i < 10; i++ {
		go func(id int, ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					cach, ok := <-counter
					if cach == 1000 {
						close(counter)
						return
					}
					if ok {
						cach++
						fmt.Println("Worker ", id, " push value: ", cach)
						counter <- cach
					} else {
						return
					}

				}
			}
		}(i, ctx)
	}

	time.Sleep(3 * time.Second)
	cancelfunc()
}
