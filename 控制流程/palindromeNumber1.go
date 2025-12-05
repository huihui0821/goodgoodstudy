package main

import "fmt"

func main() {
	num := 12321
	fmt.Printf("isSingleNum(num): %v\n", isSingleNum(num))
	fmt.Printf("num: %v\n", num)
}

func isSingleNum(num int) bool {
	if (num != 0 && num%10 == 0) || (num < 0) { // 如果数字是个数是双数返回false
		return false
	}

	numm := num
	numTmp := 0
	for numm > 0 {
		numTmp = numm%10 + numTmp*10
		fmt.Printf("numTmp: %v\n", numTmp)
		numm /= 10
	}
	return numTmp == num
}
