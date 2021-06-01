package main

import (
	"fmt"
)

// #公众号：Golang语言开发栈
// 给定一个整数数组和一个整数，找出数组中两数之和等于给定整数的两个整数。
func main() {
	nums := [9]int{3, 6, 9, 2, 4, 8, 1, 5, 7}
	target := 17
	integer1, integer2 := getIntegerV1(nums, target)
	fmt.Println(integer1, integer2)
	integer3, integer4 := getIntegerV2(nums, target)
	fmt.Println(integer3, integer4)
}

// 方式 1
// 时间复杂度 O(n²)
// 空间复杂度 O(1)
func getIntegerV1(nums [9]int, target int) (integer1, integer2 int) {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]+nums[j] == target {
				return nums[i], nums[j]
			}
		}
	}
	return 0, 0
}

// 方式 2
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func getIntegerV2(nums [9]int, target int) (integer1, integer2 int) {
	keyMap := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		difference := target - nums[i]
		if _, ok := keyMap[difference]; ok {
			return nums[i], difference
		}
		keyMap[nums[i]] = struct{}{}
	}
	return 0, 0
}
