package list

import (
    "testing"
)

// 86. 分隔链表
// https://leetcode-cn.com/problems/partition-list/
func Test_partition(t *testing.T) {

}

func partition(head *ListNode, x int) *ListNode {
    hair := &ListNode{
        Next: head,
    }
    pre, current := hair, head
    for current.Next != nil {
        if current.Next.Val >= x {
            current = current.Next
        } else {
            if pre == current { // pre和current重合时，下面的逻辑会形成自己指向自己的环链表，只需要将pre和current向后移动即可
                pre = pre.Next
                current = current.Next
            } else {
                tmp := current.Next
                current.Next = current.Next.Next
                tmp.Next = pre.Next
                pre.Next = tmp
                pre = pre.Next
            }
        }
    }
    return hair.Next
}
