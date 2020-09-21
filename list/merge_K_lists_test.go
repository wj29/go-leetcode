package list

import "testing"

// 23. 合并K个升序链表
// https://leetcode-cn.com/problems/merge-k-sorted-lists/
func Test_MergeKLists(t *testing.T) {
	list1 := NewList()
	list2 := NewList()
	list3 := NewList()
	lists := []*ListNode{list1, list2, list3}
	t.Log(mergeKLists(lists))
}

// 分治合并
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	first := mergeKLists(lists[:len(lists)/2])
	second := mergeKLists(lists[len(lists)/2:])
	return merge(first, second)
}

func merge(list1, list2 *ListNode) *ListNode {
	head := new(ListNode)
	current := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = &ListNode{
				Val: list1.Val,
			}
			current = current.Next
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			current.Next = &ListNode{
				Val: list2.Val,
			}
			current = current.Next
			list2 = list2.Next
		} else {
			current.Next = &ListNode{
				Val: list1.Val,
			}
			current = current.Next
			list1 = list1.Next
			current.Next = &ListNode{
				Val: list2.Val,
			}
			current = current.Next
			list2 = list2.Next
		}
	}
	for list1 != nil {
		current.Next = &ListNode{
			Val: list1.Val,
		}
		current = current.Next
		list1 = list1.Next
	}
	for list2 != nil {
		current.Next = &ListNode{
			Val: list2.Val,
		}
		current = current.Next
		list2 = list2.Next
	}
	return head.Next
}
