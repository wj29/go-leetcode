package list

import "testing"

// 147. 对链表进行插入排序
// https://leetcode-cn.com/problems/insertion-sort-list/
func Test_InsertionSortList(t *testing.T) {
    head := NewList()
    ret := insertionSortList(head)
    for ret != nil {
        t.Log(ret.Val)
        ret = ret.Next
    }
}

// 使用pre和head作为已排序的链表的前后端，进行比较，将新的节点插入即可
func insertionSortList(head *ListNode) *ListNode {
    hair := &ListNode{Next: head} // 使用前置节点方便对最小的数操作
    pre := hair
    for head != nil && head.Next != nil {
        if head.Next.Val >= head.Val {
            head = head.Next
            continue
        }
        for pre.Next.Val < head.Next.Val {
            pre = pre.Next
        }
        tmp := head.Next
        head.Next = head.Next.Next
        tmp.Next = pre.Next
        pre.Next = tmp
        pre = hair
    }
    return hair.Next
}
