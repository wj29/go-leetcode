package tree

import (
	"testing"
)

// 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素
// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/description/ medium
func Test_KthSmallest(t *testing.T) {
	root := NewBSTree()
	t.Log(inorderTraversal(root))
	t.Log(kthSmallest(root, 2))
}

// BST的中序遍历是有序的
func kthSmallest(root *TreeNode, k int) int {
	arr := inorderTraversal(root)
	return arr[k-1]
}
