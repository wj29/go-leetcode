package tree

import (
	"math"
	"testing"
)

// 给定一个二叉树，找出其最小深度。
// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/ easy
func Test_MinDepth(t *testing.T) {
	root := NewTree()
	t.Log(minDepth(root))
}

// 递归
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	min := math.MaxInt64
	if root.Left != nil {
		min = int(math.Min(float64(minDepth(root.Left)), float64(min)))
	}
	if root.Right != nil {
		min = int(math.Min(float64(minDepth(root.Right)), float64(min)))
	}
	return min + 1
}
