package tree

import (
	"math"
	"testing"
)

// 给定一个二叉树，计算出它的直径长度
// 二叉树的直径长度等于任意两个结点长度路径的最大值(不一定穿过根结点)
// https://leetcode-cn.com/problems/diameter-of-binary-tree/description/ easy

func Test_DiameterOfBinaryTree(t *testing.T) {
	root := NewTree()
	t.Log(diameterOfBinaryTree(root))
}

// 二叉树的直径长度等于max(左子树的直径长度，右子树的直径长度，经过根结点的最长路径)
// 左右子树的直径长度可通过递归实现，经过根结点的最长路径及左子树的高度+右子树的高度
// 拆分成子问题及遍历所有结点，每一个结点的经过该结点最长路径(左子树+右子树高度)与当前直径长度(维护一个全局直径长度)比较
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	h := 0
	max := &h
	depth(root, max)
	return *max
}

func depth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left, max)
	r := depth(root.Right, max)
	*max = int(math.Max(float64(l+r), float64(*max))) // 当前结点与全局变量比较，维护全局变量
	return int(math.Max(float64(l), float64(r))) + 1 // 返回树的高度
}