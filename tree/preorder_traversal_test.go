package tree

import (
	"testing"
)

// 二叉树的前序遍历
// 给定一个二叉树，返回它的前序遍历
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/ medium
func Test_PreorderTraversal(t *testing.T) {
	root := NewTree()
	t.Log(preorderTraversal(root))
	t.Log(preorderTraversal1(root))
}

func preorderTraversal(root *TreeNode) []int {
	arr = make([]int, 0)
	preorder(root)
	return arr
}

// 递归
func preorder(root *TreeNode) {
	if root != nil {
		arr = append(arr, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
}

// 迭代
// 前序遍历为中左右 先遍历父结点再遍历左孩子结点再遍历右孩子结点
func preorderTraversal1(root *TreeNode) []int {
	var (
		res   []int
		stack []*TreeNode
	)
	for len(stack) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1   //栈顶
		root = stack[index].Right //出栈
		stack = stack[:index]
	}
	return res
}
