package main

import (
	"fmt"
)

func main() {
	str := "({}("
	fmt.Printf("str: %v\n", str)
	fmt.Printf("isValid(str): %v\n", isValid(str))

}

func isValid(str string) bool {
	lenTemp := len(str)
	if lenTemp%2 == 1 {
		return false
	}
	myMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	for i := 0; i < lenTemp; i++ {
		c := str[i]
		if c1, isRight := myMap[c]; isRight {
			if len(stack) == 0 || stack[len(stack)-1] != c1 {
				return false
			} else {
				stack = stack[:len(stack)-1] // 切片
			}
		} else {
			// 如果是左括号则入栈
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
