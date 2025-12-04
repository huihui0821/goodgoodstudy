package main

import (
	"fmt"
)

func main() {
	num := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 8, 7, 5, 4, 3, 2, 1}
	fmt.Printf("singleNumber(num[:]): %v\n", singleNumber(num[:]))

}

func singleNumber(nums []int) int {
	number := 1
	for i := 0; i < (len(nums) - 1); i++ {
		number = 1
		for j := i + 1; j < len(nums); j++ {
			// fmt.Printf("nums[i]: %v nums[j]: %v\n", nums[i], nums[j])
			if nums[i] == nums[j] {
				number++
			}
		}
		if number == 1 {
			return nums[i]
		}
	}
	return -1
}
