package tree

import (
	"math"
	"testing"
)

// 给定一个二叉树，判断它是否是高度平衡的二叉树
// https://leetcode-cn.com/problems/balanced-binary-tree/description/ easy
func Test_IsBalanced(t *testing.T)  {
	root := NewTree()
	t.Log(isBalanced(root))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && int(math.Abs(float64(maxDepth(root.Left)) - float64(maxDepth(root.Right)))) <= 1
}