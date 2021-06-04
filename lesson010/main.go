package main

import "fmt"

// #公众号：Golang语言开发栈
// 力扣 876
// 给定一个头结点为head的非空单链表，返回链表的中间结点。如果有两个中间结点，则返回第二个中间结点

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

// 时间复杂度 O(n)
// 空间复杂度 O(1)
// 中间节点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	l := convertListNode(s)
	result := middleNode(l)
	result.printListNode()
}
