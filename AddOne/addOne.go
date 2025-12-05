package main

import "fmt"

func main() {
	array := [4]int{2, 3, 4, 9}
	// fmt.Printf("array: %v\n", array)
	// fmt.Printf("arrayProcess(array[:]): %v\n", arrayProcess(array[:]))
	fmt.Printf("arrayProcess(array[:]): %v\n", arrayProcess(array[:]))
}

func arrayProcess(arr []int) []int {
	n := len(arr)
	fmt.Printf("n: %v\n", n)
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("i: %v\n", i)
		if arr[i] != 9 {
			arr[i]++
			return arr
		}
		arr[i] = 0
	}

	arr = make([]int, n+1)
	arr[0] = 1
	return arr
}
