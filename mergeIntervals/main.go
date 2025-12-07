package main

import (
	"fmt"
	"sort"
)

// 区间合并
func main() {
	arrs := [][]int{{2, 4}, {1, 3}, {5, 7}, {8, 10}}
	fmt.Printf("arrs: %v\n", arrs)
	ret := mergeIntervals(arrs)
	fmt.Printf("ret: %v\n", ret)
}

func mergeIntervals(arrs [][]int) [][]int {
	lenArr := len(arrs)
	if lenArr <= 1 {
		return arrs
	}

	retArrs := make([][]int, 0, lenArr)
	// 使用sort库函数对二维数字进行排序
	sort.Slice(arrs, func(i, j int) bool {
		return arrs[i][0] < arrs[j][0]
	})
	fmt.Printf("arrs: %v\n", arrs)

	retArrs = append(retArrs, arrs[0])
	for i := 1; i < lenArr; i++ {
		arr := arrs[i]
		ret := retArrs[(len(retArrs))-1]
		// 出现区间重叠
		if ret[1] >= arr[0] {
			retArrs[len(retArrs)-1][1] = arr[1]
		} else {
			// m没有重叠就将arrs[i] 追加到retArrs,添加一个新的区间
			retArrs = append(retArrs, arrs[i])
		}
	}

	return retArrs
}
