package counterMTX

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

type Counter struct {
	sync.Mutex
	value   int64
	max_val int64
	stop_F  context.CancelFunc
}

func NewCounter(size int64, exit_F context.CancelFunc) *Counter {
	return &Counter{
		value:   0,
		max_val: size,
		stop_F:  exit_F,
	}
}

func (c *Counter) Add() {
	c.Lock()
	defer c.Unlock()
	if c.value >= c.max_val {
		c.stop_F()
		return
	}
	if c.value%10 == 0 {
		fmt.Println("GoGOFOR")
		runtime.Gosched()
	}
	c.value++
	fmt.Println("Значение счетчика: ", c.value)

}
