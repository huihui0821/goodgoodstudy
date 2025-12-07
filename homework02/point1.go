package main

import "fmt"

func addTen(num *int) {
	*num += 10
}

func main() {
	value := 30

	fmt.Printf("修改前: %v\n", value)

	addTen(&value)

	fmt.Printf("修改后: %v\n", value)
}
