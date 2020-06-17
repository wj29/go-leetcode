package list

import (
	"testing"
)

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换
// https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/ medium
func Test_SwapPairs(t *testing.T) {
	head := NewList() // 1 2 6 7 8
	r := swapPairs(head)
	for r != nil {
		t.Log(r.Val)
		r = r.Next
	}
}

// 遍历调换
// 在链表元素进行调换时注意防止形成环链
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := new(ListNode)
	newHead.Next = head
	current := newHead
	for head != nil && head.Next != nil {
		current.Next = head.Next
		tmp := head // 记录当前head位置
		head = head.Next.Next
		current.Next.Next = tmp // 防止形成环链
		current.Next.Next.Next = head
		current = current.Next.Next
	}
	return newHead.Next
}

// 结点分为奇数结点和偶数结点，遍历得到两个链表组合成新的链表