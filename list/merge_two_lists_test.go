package list

import (
	"testing"
)

// 给定两个升序链表，按照升序排列合并成一个链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists/description/ easy

func Test_MergeTwoLists(t *testing.T) {
	l1 := NewList()
	l2 := NewList()
	l3 := mergeTwoLists(l1, l2)
	for l3 != nil {
		t.Log(l3.Val)
		l3 = l3.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l3 := new(ListNode)
	newHead := l3
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l3.Next = l1
			l3 = l3.Next
			l1 = l1.Next
		} else {
			l3.Next = l2
			l3 = l3.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		l3.Next = l1
	}
	if l2 != nil {
		l3.Next = l2
	}
	return newHead.Next
}
