package main

import "fmt"

func main() {
	arr := [10]int{1, 1, 2, 3, 3, 4, 4, 7, 7, 9}
	fmt.Printf("deleRepeat(arr[:]): %v\n", deleRepeat(arr[:]))

}

func deleRepeat(arr []int) int {
	len := len(arr)

	j := 0
	for i := 1; i < len; i++ {
		if arr[j] != arr[i] {
			arr[j+1] = arr[i]
			j++
		}
	}

	return j + 1
}
