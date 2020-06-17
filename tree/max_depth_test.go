package tree

import (
	"math"
	"testing"
)

// 给定一个二叉树，找出其最大深度。
// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/ easy
func Test_MaxDepth(t *testing.T)  {
	root := NewTree()
	t.Log(maxDepth(root))
}

// 树的高度肯定是左子树或者右子树的较高的+1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}