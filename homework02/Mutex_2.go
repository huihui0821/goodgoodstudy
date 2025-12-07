package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const count = 1000
const goroutines = 10

type Counter struct {
	count int64
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Value() int64 {
	return c.count
}

func main() {
	counter := &Counter{}
	wg := sync.WaitGroup{}
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			for j := 0; j < count; j++ {
				atomic.AddInt64((*int64)(&counter.count), 1)
			}
			fmt.Printf("协程 %d 完成 1000 次递增\n", id)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	println("Final Counter Value:", counter.Value())
}
