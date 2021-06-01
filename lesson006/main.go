package main

import "fmt"

// #公众号：Golang语言开发栈
// 力扣 21
// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

type ListNode struct {
	Val  int
	Next *ListNode
}

// 打印
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

// 合并链表
// 时间复杂度 O(M+N)
// 空间复杂度 O(M+N)
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	temp := &ListNode{}
	if l1.Val <= l2.Val {
		temp = l1
		temp.Next = mergeTwoLists(l1.Next, l2)
	} else {
		temp = l2
		temp.Next = mergeTwoLists(l1, l2.Next)
	}
	return temp
}

func main() {
	s1 := []int{1, 2, 4}
	s2 := []int{1, 3, 4}
	l1 := convertListNode(s1)
	l2 := convertListNode(s2)
	listNode := mergeTwoLists(l1, l2)
	listNode.printListNode()
}
