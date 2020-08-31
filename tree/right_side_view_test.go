package tree

import (
	"testing"
)

// 199. 二叉树的右视图
// https://leetcode-cn.com/problems/binary-tree-right-side-view/
func Test_RightSideView(t *testing.T) {
	root := NewTree()
	t.Log(rightSideView(root))
}

// 从右侧层序遍历
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	pre := []*TreeNode{root}
	for len(pre) > 0 {
		tmp := make([]*TreeNode, 0)
		for i := 0; i < len(pre); i++ {
			if i == len(pre)-1 {
				ret = append(ret, pre[i].Val)
			}
			if pre[i].Left != nil {
				tmp = append(tmp, pre[i].Left)
			}
			if pre[i].Right != nil {
				tmp = append(tmp, pre[i].Right)
			}
		}
		pre = tmp
	}
	return ret
}
