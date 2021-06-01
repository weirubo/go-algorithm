package main

import (
	"fmt"
)

// #公众号：Golang语言开发栈
// 给定一组整数，其中有两个整数重复，找出重复的整数
func main() {
	nums := [9]int{3, 6, 9, 2, 4, 8, 3, 5}
	val1 := getRepeatNumV1(nums)
	fmt.Println(val1)

	val2 := getRepeatNumV2(nums)
	fmt.Println(val2)
}

// 方式 1
// 时间复杂度 O(n²)
// 空间复杂度 O(1)
func getRepeatNumV1(nums [9]int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] == nums[i] {
				return nums[i]
			}
		}
	}
	return 0
}

// 方式 2
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func getRepeatNumV2(nums [9]int) int {
	keyMap := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := keyMap[nums[i]]; ok {
			return nums[i]
		}
		keyMap[nums[i]] = struct{}{}
	}
	return 0
}
