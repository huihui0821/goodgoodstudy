package main

import "fmt"

func main() {
	str := []string{"string", "straaa", "strqqq"}
	strRet := longestPrefix(str)
	fmt.Printf("strRet: %v\n", strRet)

}

//  选出第一个字符串，从第一个字符串的第一个字符开始，遍历每个字符串的第一个字符
//  string
//  straaa
//  strqqq

func longestPrefix(str []string) string {
	s1 := str[0]

	for i, c := range s1 {
		for _, s := range str {
			if i == len(s) || s[i] != byte(c) {
				return s1[0:i]
			}
		}
	}

	return s1
}
