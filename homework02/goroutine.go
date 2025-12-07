package main

import (
	"fmt"
	"time"
)

func printOdds() {
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数协程: %d\n", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func printEvens() {
	for i := 0; i < 10; i += 2 {
		fmt.Printf("偶数协程: %d\n", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go printOdds()
	go printEvens()

	time.Sleep(2 * time.Second)

	fmt.Println("主协程结束，程序退出")
}
