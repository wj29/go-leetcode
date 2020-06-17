package tree

import (
	"strconv"
	"testing"
)

// 给定一个二叉树，返回所有从根节点到叶子节点的路径
// https://leetcode-cn.com/problems/binary-tree-paths/  easy
func Test_BinaryTreePaths(t *testing.T) {
	root := NewTree()
	t.Log(binaryTreePaths(root))
}

func binaryTreePaths(root *TreeNode) []string {
	ret := make([]string, 0)
	traversal3(root, "", &ret)
	return ret
}

func traversal3(root *TreeNode, pre string, str *[]string) {
	if root != nil {
		if pre == "" {
			pre = strconv.Itoa(root.Val)
		} else {
			pre += "->" + strconv.Itoa(root.Val)
		}
		traversal3(root.Left, pre, str)
		traversal3(root.Right, pre, str)
		if root.Left == nil && root.Right == nil {
			*str = append(*str, pre)
		}
	}
}
