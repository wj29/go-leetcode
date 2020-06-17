package tree

import (
	"testing"
)

// 给定一个二叉树，翻转它(左右孩子结点互换)
// https://leetcode-cn.com/problems/invert-binary-tree/description/ easy
func Test_InvertTree(t *testing.T) {
	root := NewTree()
	t.Log(preorderTraversal(root))
	ret := invertTree(root)
	t.Log(preorderTraversal(ret))
}

// 翻转左右子树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}