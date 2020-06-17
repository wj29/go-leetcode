package tree

import (
	"testing"
)

// 二叉树的后序遍历
// 给定一个二叉树，返回它的后序遍历
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/  hard
func Test_PostorderTraversal(t *testing.T) {
	root := NewTree()
	t.Log(postorderTraversal(root))
	t.Log(postorderTraversal1(root))
}

func postorderTraversal(root *TreeNode) []int {
	arr = make([]int, 0)
	postorder(root)
	return arr
}

// 递归
func postorder(root *TreeNode) {
	if root != nil {
		postorder(root.Left)
		postorder(root.Right)
		arr = append(arr, root.Val)
	}
}

// 迭代
// 后序遍历为左右中 先遍历左孩子结点再遍历右孩子结点再遍历父结点
// 前序遍历为中左右，那么我们可以将前序遍历改成中右左，那么就是后序遍历的倒序
func postorderTraversal1(root *TreeNode) []int {
	var (
		res   []int
		stack []*TreeNode
	)
	for len(stack) != 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Right
		}
		index := len(stack) - 1
		root = stack[index].Left
		stack = stack[:index]
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
