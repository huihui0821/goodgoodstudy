package main

import (
	"fmt"
)

func main() {
	num := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 8, 7, 5, 4, 3, 2, 1}
	fmt.Printf("singleNumber(num[:]): %v\n", singleNumber(num[:]))

}

func singleNumber(nums []int) int {
	single := 0
	for _, nummber := range nums {
		single ^= nummber
	}
	return single
}
