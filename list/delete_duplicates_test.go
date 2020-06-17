package list

import "testing"

// 给定一个排序列表，删除重复的元素，使得每个元素只出现一次
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/  easy
func Test_DeleteDuplicates(t *testing.T) {

}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
