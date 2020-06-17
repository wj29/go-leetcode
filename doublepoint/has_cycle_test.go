package doublepoint

import (
	"testing"
	"github.com/wujiang1994/go-leetcode/list"
)

// 给定一个链表，判断是否有环 easy
// https://leetcode-cn.com/problems/linked-list-cycle/description/

func Test_HasCycle(t *testing.T) {
	head := list.NewCycleList()
	t.Log(hasCycle(head))
}

// 快慢指针，如果有环，那么它们一定会相遇在交点
// map实现
func hasCycle(head *list.ListNode) bool {
	if head == nil {
		return false
	}
	p1 := head
	p2 := head.Next
	for p1 != nil && p2 != nil && p2.Next != nil {
		if p1 == p2 {
			return true
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return false
}
