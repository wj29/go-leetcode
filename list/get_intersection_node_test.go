package list

import (
	"testing"
)

// 编写一个程序找到链表相交的起点，链表没有环
// 程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存
// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/description/ easy

func Test_GetIntersectionListNode(t *testing.T) {
	p1 := &ListNode{
		Val: 1,
	}
	p2 := &ListNode{
		Val: 2,
	}
	p3 := &ListNode{
		Val: 3,
	}
	p4 := &ListNode{
		Val: 4,
	}
	p5 := &ListNode{
		Val: 5,
	}
	p6 := &ListNode{
		Val: 6,
	}
	p7 := &ListNode{
		Val: 7,
	}
	p8 := &ListNode{
		Val: 8,
	}
	p1.Next = p2
	p2.Next = p6
	p6.Next = p7
	p7.Next = p8
	p3.Next = p4
	p4.Next = p5
	p5.Next = p6
	t.Log(getIntersectionListNode(p1, p1))
}

// 暴力遍历两个链表看看是不是存在相同的结点，时间复杂度MN，空间复杂度1
// M+N的时间复杂度可以用map实现，空间复杂度M或者N，该结点作为map的key，遍历时存在即相交
// 1的空间复杂度代表不能用map，不能开辟新的空间或者仅开辟一个空间

// 双指针法(错的人迟早会走散，对的人迟早会相逢)
// 链表如果相交，那么AB两个指针相交之后的结点必定是重合的
// 链表的长度不一样，但是A+B和B+A的长度相同，由拥有相同长度的相交，AB第一次相遇的结点就是相交的结点、
// 链表如果不相交，他们末尾的结点必定不相同

// 如果只判断两链表是否相交
// 将其中一个链表头指向两外一个链表尾，判断是否有环
// 判断两个链表结尾是否相同
func getIntersectionListNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	if headA == headB {
		return headA
	}
	currentA := headA
	currentB := headB
	var lastListNode *ListNode
	for {
		if currentA.Next != nil {
			currentA = currentA.Next
		} else {
			if lastListNode == nil {
				lastListNode = currentA
			} else {
				if lastListNode != currentA {
					return nil
				}
			}
			currentA = headB
		}
		if currentB.Next != nil {
			currentB = currentB.Next
		} else {
			if lastListNode == nil {
				lastListNode = currentB
			} else {
				if lastListNode != currentB {
					return nil
				}
			}
			currentB = headA
		}
		if currentB == currentA {
			return currentA
		}
	}
}
