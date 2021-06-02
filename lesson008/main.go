package main

import "fmt"

// #公众号：Golang语言开发栈
// 力扣 206
// 给你单链表的头节点head，请你反转链表，并返回反转后的链表。

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

// 反转链表
// 时间复杂度 O(n)
// 空间复杂度 O(1)
// prev 上一个节点 current 当前节点 next 下一个节点
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	s := []int{1, 3, 5, 7, 9}
	l := convertListNode(s)
	result := reverseList(l)
	result.printListNode()
}
