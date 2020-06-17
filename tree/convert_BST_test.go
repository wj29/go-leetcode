package tree

import (
	"testing"
)

// 给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，
// 使得每个节点的值是原来的节点值加上所有大于它的节点值之和
// https://leetcode-cn.com/problems/convert-bst-to-greater-tree easy
func Test_ConvertBST(t *testing.T) {
	root := NewSmallBSTree()
	t.Log(inorderTraversal(root))
	t.Log(inorderTraversal(convertBST(root)))
}

//
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var sum int
	add(root, &sum)
	return root
}

func add(root *TreeNode, sum *int) {
	if root != nil {
		add(root.Right, sum)
		*sum += root.Val
		root.Val = *sum
		add(root.Left, sum)
	}
}
