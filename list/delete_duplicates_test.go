package list

import "testing"

// 给定一个排序列表，删除重复的元素，使得每个元素只出现一次
// 83. 删除排序链表中的重复元素
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/

// 82. 删除排序链表中的重复元素 II
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
func Test_DeleteDuplicates(t *testing.T) {
    head := NewList()
    t.Log(deleteDuplicates(head))
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

func deleteDuplicates2(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    hair := &ListNode{
        Next: head,
    }
    pre, current := hair, head
    count := 0
    for current != nil {
        if current.Val == pre.Next.Val { // 寻找重复元素
            current = current.Next
            count++
            // 末尾判断
            if current == nil && count > 1 {
                pre.Next = current
            }
        } else if count > 1 { // 删除所有重复元素
            count = 0
            pre.Next = current
        } else { // 加载不重复元素
            pre = pre.Next
            current = current.Next
        }
    }
    return hair.Next
}
