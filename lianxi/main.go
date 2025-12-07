package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "结果来了"
	}()

	fmt.Println("等待中...")
	msg := <-ch
	fmt.Println(msg)
}
