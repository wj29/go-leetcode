package tree

import "testing"

// 给定一个二叉树，检查它是否是镜像对称的
// https://leetcode-cn.com/problems/symmetric-tree/ easy
func Test_IsSymmetric(t *testing.T) {
	root := NewTree()
	t.Log(isSymmetric(root))
}

// 如果是一个对称二叉树，那么它的左子树和右子树是对称的，递归即可
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(left, right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	if left == nil && right == nil {
		return true
	}
	return left.Val == right.Val && symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
}