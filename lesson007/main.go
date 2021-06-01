package main

import "fmt"

// #公众号：Golang语言开发栈
// 力扣 83
// 存在一个按升序排列的链表，给你这个链表的头节点head，请你删除所有重复的元素，使每个元素只出现一次。返回同样按升序排列的结果链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 打印链表
func (l *ListNode) printListNode() {
	var s []int
	pointer := l
	for pointer != nil {
		s = append(s, pointer.Val)
		pointer = pointer.Next
	}
	for _, v := range s {
		fmt.Println(v)
	}
}

// 转换链表
func convertListNode(s []int) *ListNode {
	sLen := len(s)
	if sLen == 0 {
		return nil
	}
	listNode := &ListNode{
		Val: s[0],
	}
	temp := listNode
	for i := 1; i < sLen; i++ {
		temp.Next = &ListNode{
			Val: s[i],
		}
		temp = temp.Next
	}
	return listNode
}

// 删除重复元素
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func main() {
	s := []int{1, 1, 2, 2, 3, 4, 4, 5, 5, 6, 6}
	l := convertListNode(s)
	result := deleteDuplicates(l)
	result.printListNode()
}
