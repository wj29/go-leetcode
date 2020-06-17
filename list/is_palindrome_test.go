package list

import (
	"testing"
)

// 判断一个链表是否为回文链表
// https://leetcode-cn.com/problems/palindrome-linked-list/description/ easy
// 进阶: O(n) 时间复杂度和 O(1) 空间复杂度
func Test_IsPalindrome(t *testing.T) {
	head := NewList()
	isPalindromeByArr(head)
}

// 将链表放入数组，判断数组即可
func isPalindromeByArr(head *ListNode) bool {
	if head == nil {
		return true
	}
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	length := len(arr)
	for i := 0; i < length/2; i++ {
		if arr[i] != arr[length-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	return false
}
