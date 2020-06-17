package tree

import (
	"testing"
)

// 给定一个二叉搜索树，同时给定最小边界L 和最大边界 R。通过修剪二叉搜索树，使得所有节点的值在[L, R]中 (R>=L)
// 你可能需要改变树的根节点，所以结果应当返回修剪好的二叉搜索树的新的根节点
// https://leetcode-cn.com/problems/trim-a-binary-search-tree  easy
func Test_TrimBST(t *testing.T) {
	root := NewBSTree()
	t.Log(inorderTraversal(root))
	L, R := 1, 3
	t.Log(inorderTraversal(trimBST(root, L, R)))
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L { // 删除所有左子树
		return trimBST(root.Right, L, R)
	}
	if root.Val > R { // 删除所有右子树
		return trimBST(root.Left, L, R)
	}
	// 处理正常的结点
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}
