package main

import (
	"fmt"
	"sort"
)

// #公众号：Golang语言开发栈
// 力扣 26
// 删除有序数组中的重复项，返回删除后数组的新长度和新数组
func main() {
	data := []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4}
	len1, arr1 := removeV1(data)
	fmt.Println(len1, arr1)
	len2, arr2 := removeV2(data)
	fmt.Println(len2, arr2)
}

// 时间复杂度 O(n)
// 空间复杂度 O(n)
func removeV1(arr []int) (int, []int) {
	keyMap := make(map[int]struct{})
	for i := 0; i < len(arr); i++ {
		if _, ok := keyMap[arr[i]]; !ok {
			keyMap[arr[i]] = struct{}{}
		}
	}
	length := len(keyMap)
	arrNew := make([]int, 0, length)
	for k := range keyMap {
		arrNew = append(arrNew, k)
	}
	sort.Ints(arrNew)
	return length, arrNew
}

// 时间复杂度 O(n)
// 空间复杂度 O(1)
func removeV2(arr []int) (int, []int) {
	if len(arr) < 2 {
		return len(arr), arr
	}
	slow := 1
	for fast := 1; fast < len(arr); fast++ {
		if arr[fast] != arr[fast-1] {
			arr[slow] = arr[fast]
			slow++
		}
	}
	return slow, arr[0:slow]
}
