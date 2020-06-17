package list

import (
	"testing"
)

// 给定一个单链表，把所有的奇数节点和偶数节点分别排在一起
// 算法的空间复杂度应为1，时间复杂度应为N
// 输入: 1->2->3->4->5->NULL
// 输出: 1->3->5->2->4->NULL
// https://leetcode-cn.com/problems/odd-even-linked-list/  medium
func Test_OddEvenList(t *testing.T) {
	head := NewList()

	r := oddEvenList(head)
	for r != nil {
		t.Log(r.Val)
		r = r.Next
	}
}

// 通过双指针遍历整个链表即可得到结果
// 一次遍历时间复杂度N，使用常数变量空间复杂度复杂度1
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	even := odd.Next
	for odd != nil && even != nil && even.Next != nil {
		oddTmp := odd.Next        // 记录当前奇数结点的下一个结点，后续中间会插入另一个奇数结点
		evenTmp := even.Next.Next // 记录下一个偶数结点位置，偶数后的奇数结点会插入上一行中
		odd.Next = even.Next      // 将奇数结点next指向新的奇数结点
		odd.Next.Next = oddTmp    // 连接奇数结点与偶数第一个结点
		even.Next = evenTmp       // 连接偶数末尾结点至未完成排序的链表
		odd = odd.Next            // 移动奇数结点至已排序的末尾
		even = even.Next          // 移动偶数结点至已排序的末尾
	}
	return head
}
