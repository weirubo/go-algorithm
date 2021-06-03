package main

import "fmt"

// #公众号：Golang语言开发栈
// 力扣 203
// 给你一个链表的头节点head和一个整数val，请你删除链表中所有满足Node.Val == val的节点，并返回新链表。

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
// 移除链表中的元素
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	prev := dummy
	for head != nil {
		if head.Val == val {
			prev.Next = head.Next
		} else {
			prev = prev.Next
		}
		head = head.Next
	}
	return dummy.Next
}

func main() {
	s := []int{2, 3, 5, 4, 6, 5, 7, 8}
	l := convertListNode(s)
	val := 5
	result := removeElements(l, val)
	// 输出结果链表
	result.printListNode()
}
