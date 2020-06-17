package tree

import (
	"testing"
)

// 给定两个二叉树，用其中一个覆盖另一个，新结点值为两个对应位置结点值相加(nil的值为0)
// https://leetcode-cn.com/problems/merge-two-binary-trees/description/ easy
func Test_MergeTrees(t *testing.T) {
	root1 := NewTree()
	root2 := NewTree()
	ret := mergeTrees(root1, root2)
	t.Log(preorderTraversal(ret))
}

// 使用两个指针分别遍历两棵树，对应位置相加并赋值给目标树
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
