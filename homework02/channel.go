package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	go func() { //生产携程
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Printf("生产了: %v\n", i)
		}
		close(ch)
		println("生产结束")
	}()

	go func() { //消费携程
		for v := range ch {
			fmt.Printf("消费了: %v\n", v)
			time.Sleep(1 * time.Second)
		}
		println("消费结束")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("主程序结束")

}
