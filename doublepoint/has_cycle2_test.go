package doublepoint

import (
	"testing"
	"github.com/wujiang1994/go-leetcode/list"
)

// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null
// 不改变原链表
// https://leetcode-cn.com/problems/linked-list-cycle-ii/ medium
func Test_DetectCycle(t *testing.T) {
	head := list.NewCycleList()
	t.Log(detectCycle(head))
	t.Log(detectCycle1(head))
}

// hash遍历到最后一个结点为nil则不相交，hash出现相同的第一个值就是第一个相交的结点
func detectCycle(head *list.ListNode) *list.ListNode {
	m := make(map[*list.ListNode]bool)
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		} else {
			m[head] = true
		}
		head = head.Next
	}
	return nil
}

// 题解见 arrayandmatrix/find_duplicate_test.go
// 相似的题目，判断两个链表是否相交 list/get_intersection_listNode.go
func detectCycle1(head *list.ListNode) *list.ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head.Next
	flag := false
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			flag = true
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	if flag {
		slow = new(list.ListNode)
		slow.Next = head
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		return fast
	}
	return nil
}
