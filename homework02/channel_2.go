package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
			fmt.Printf("生产者发送: %v (缓冲区剩余容量: %v)\n", i, len(ch))
			if len(ch) >= 8 {
				time.Sleep(50 * time.Millisecond)
			}
		}
		close(ch)
		println("协程1执行完毕")
		defer wg.Done()
	}()

	go func() {
		for v := range ch {
			fmt.Printf("消费者接收: %v (缓冲区剩余容量: %v)\n", v, len(ch))
		}
		println("协程2执行完毕")
		defer wg.Done()
	}()
	wg.Wait()
}
