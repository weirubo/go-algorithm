package main

import (
	"fmt"
)

// #公众号：Golang语言开发栈
// 给定一个整数数组和一个整数，将数组中和给定整数相等的元素移除，返回移除元素后的新数组的新长度和新数组，不需要考虑新数组中超出新长度后面的元素，新长度个元素可为任意顺序

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	target := 2
	len1, arr1 := removeTargetV1(arr, target)
	fmt.Println(len1, arr1)
	len2, arr2 := removeTargetV2(arr, target)
	fmt.Println(len2, arr2)
}

// 时间复杂度 O(n)
// 空间复杂度 O(n)
func removeTargetV1(arr []int, target int) (int, []int) {
	keyMap := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if arr[i] != target {
			keyMap[i] = arr[i]
		}
	}
	newArr := make([]int, 0, len(keyMap))
	for _, v := range keyMap {
		newArr = append(newArr, v)
	}
	return len(newArr), newArr
}

// 时间复杂度 O(n)
// 空间复杂度 O(1)
func removeTargetV2(arr []int, target int) (int, []int) {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != target {
			arr[j] = arr[i]
			j++
		}
	}
	return j, arr
}
