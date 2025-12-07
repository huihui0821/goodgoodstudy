package main

import (
	"fmt"
	"sync"
)

const count = 1000
const goroutines = 10

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	counter := &Counter{}
	wg := sync.WaitGroup{}
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			for j := 0; j < count; j++ {
				counter.Increment()
			}
			fmt.Printf("协程 %d 完成 1000 次递增\n", id)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	println("Final Counter Value:", counter.Value())
}
