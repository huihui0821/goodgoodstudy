package main

import "fmt"

func multiplySliceByTwo(slice *[]int) {
	for i := range *slice {
		(*slice)[i] *= 2
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 10, 20}

	fmt.Println("修改前:", numbers)

	multiplySliceByTwo(&numbers)

	fmt.Println("修改后:", numbers)
}
