package tree

import (
	"github.com/wujiang1994/go-leetcode/list"
	"testing"
)

// 给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。
// 如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。
// 一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径
// https://leetcode-cn.com/problems/linked-list-in-binary-tree
func Test_IsSubPath(t *testing.T) {
	root := NewTree()
	head := &list.ListNode{
		Val: 1,
	}
	p2 := &list.ListNode{
		Val: 2,
	}
	p4 := &list.ListNode{
		Val: 4,
	}
	head.Next = p2
	p2.Next = p4
	t.Log(isSubPath(head, root))
}

// 找到所有的开始结点，然后每个开始结点向下遍历即可
func isSubPath(head *list.ListNode, root *TreeNode) bool {
	if head == nil || root == nil {
		return false
	}
	ret := make([]*TreeNode, 0)
	traversalForIsSubPath(root, head.Val, &ret)
	for _, v := range ret {
		if isSub(v, head) {
			return true
		}
	}
	return false
}

func traversalForIsSubPath(root *TreeNode, target int, ret *[]*TreeNode) {
	if root != nil {
		if root.Val == target {
			*ret = append(*ret, root)
		}
		traversalForIsSubPath(root.Left, target, ret)
		traversalForIsSubPath(root.Right, target, ret)
	}
}

func isSub(root *TreeNode, head *list.ListNode) bool {
	if root == nil {
		return false
	}
	if root.Val != head.Val {
		return false
	}
	if head.Next == nil {
		return true
	}
	return isSub(root.Left, head.Next) || isSub(root.Right, head.Next)
}
