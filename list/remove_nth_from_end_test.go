package list

import (
	"testing"
)

// 给定一个链表，删除倒数第n个结点并返回
// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/ medium
// 进阶: 一趟扫描实现
func Test_RemoveNthFromEnd(t *testing.T) {
	head := NewList() // 1 2 6 7 8
	r := removeNthFromEndByTwice(head, 1)
	for r != nil {
		t.Log(r.Val)
		r = r.Next
	}
}

// 两次遍历
func removeNthFromEndByTwice(head *ListNode, n int) *ListNode {
	current := head
	length := 0
	for current != nil {
		length++
		current = current.Next
	}
	length -= n
	newHead := new(ListNode) // 此处新建一个node并指向head 因为他可能要删除第一个结点，此时需要前一个结点指向第二个结点
	newHead.Next = head
	current = newHead
	for length > 0 {
		length--
		current = current.Next
	}
	current.Next = current.Next.Next
	return newHead.Next
}

// 单次遍历通过定长n实现
// 倒数第n个，即末尾向前数n个，通过双指针描述即第一个指针指向nil时，第二个指针指向倒数第n个结点
// 我们通过第一个指针向前走n+1步，第二个指针从头开始，它们相隔n，当第一个指针指向结尾时，第二个指针指向倒数第n个
func removeNthFromEndByOnce(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := new(ListNode) // newHead原因同上
	newHead.Next = head
	p1 := newHead
	p2 := newHead
	count := 0

	for {
		count++
		p2 = p2.Next
		if count < n {
			continue
		}
		if p2.Next == nil {
			p1.Next = p1.Next.Next
			break
		}
		p1 = p1.Next
	}
	return newHead.Next
}
