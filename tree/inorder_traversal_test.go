package tree

import (
	"testing"
)

// 二叉树的中序遍历
// 给定一个二叉树，返回它的中序遍历
// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/ medium
func Test_InorderTraversal(t *testing.T) {
	root := NewTree()
	t.Log(inorderTraversal(root))
	t.Log(inorderTraversal1(root))
}

// 递归
func inorderTraversal(root *TreeNode) []int {
	arr = make([]int, 0)
	inorder(root)
	return arr
}

func inorder(root *TreeNode) {
	if root != nil {
		inorder(root.Left)
		arr = append(arr, root.Val)
		inorder(root.Right)
	}
}

// 迭代
// 中序遍历为左中右 先遍历左孩子结点再遍历父结点再遍历右孩子结点
//
// 结点不为nil则入栈，继续遍历左孩子结点，直到左孩子结点为nil
// 出栈，相当于后退一步，将出栈元素放入slice
// 遍历其右孩子结点，又进入访问右子树的左孩子结点中
// 直到stack==0且root==nil
func inorderTraversal1(root *TreeNode) []int {
	var (
		res   []int
		stack []*TreeNode // 栈
	)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root) // 入栈 栈中依次存的是所有的父节点及左孩子结点
			root = root.Left // 叶子结点的左孩子结点一定是空，那么它本身就是中序遍历的中
		}
		index := len(stack) - 1 // 栈顶
		res = append(res, stack[index].Val) // 中序输出
		root = stack[index].Right
		stack = stack[:index] // 出栈
	}
	return res
}
