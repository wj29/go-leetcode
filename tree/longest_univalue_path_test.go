package tree

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。
// 注意：两个节点之间的路径长度由它们之间的边数表示。
// https://leetcode-cn.com/problems/longest-univalue-path/
func Test_LongestUnivaluePath(t *testing.T) {
	root := NewTree()
	t.Log(longestUnivaluePath(root))
}

// 对于任意一个节点，经过它的最长边数等于其左子树最长边数+右子树的最长边数+2
func longestUnivaluePath(root *TreeNode) int {
	ret := 0
	traversalForUnivaluePath(root, &ret)
	return ret
}

func traversalForUnivaluePath(root *TreeNode, ret *int) {
	if root == nil {
		return
	}
	univaluePath(root, ret)
	traversalForUnivaluePath(root.Left, ret)
	traversalForUnivaluePath(root.Right, ret)
}

func univaluePath(root *TreeNode, ret *int) int {
	if root == nil {
		return 0
	}
	left, right := 0, 0
	if root.Left != nil && root.Left.Val == root.Val {
		left = univaluePath(root.Left, ret) + 1

	}
	if root.Right != nil && root.Right.Val == root.Val {
		right = univaluePath(root.Right, ret) + 1
	}
	*ret = common.Max(*ret, left+right)
	return common.Max(left, right)
}
