package main

import (
	"fmt"
	"math"
)

// #公众号：Golang语言开发栈
// 给定一个无重复元素的有序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置
func main() {
	arr := []int{1, 3, 4, 5, 7, 9}
	target := 4
	v1 := searchV1(arr, target)
	fmt.Println(v1)
	v2 := searchV2(arr, target)
	fmt.Println(v2)
}

// 时间复杂度 O(n)
// 空间复杂度 O(1)
func searchV1(arr []int, target int) int {
	var i = 0
	for i = 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		} else if arr[i] > target {
			return i
		}
	}
	return i
}

// 时间复杂度 O(logN)
// 空间复杂度 O(1)
func searchV2(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := int(math.Floor(float64((low + high) / 2)))
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		}
	}
	return high + 1
}
