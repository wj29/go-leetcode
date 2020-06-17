package tree

import "testing"

// 给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。
// s 也可以看做它自身的一棵子树。
// https://leetcode-cn.com/problems/subtree-of-another-tree
func Test_IsSubtree(t *testing.T) {
	s := NewTree()
	p := &TreeNode{
		Val:   2,
	}
	p.Left = &TreeNode{
		Val:   4,
	}
	p.Right = &TreeNode{
		Val:   5,
	}
	t.Log(isSubtree(s, p))
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}
	ret := make([]*TreeNode, 0)
	traversalForRoot(s, t.Val, &ret)
	for _, v := range ret {
		if traversalIsSubTree(v, t) {
			return true
		}
	}
	return false
}

// 找到s中所有与t.val相等的结点
func traversalForRoot(root *TreeNode, target int, ret *[]*TreeNode) {
	if root != nil {
		if root.Val == target {
			*ret = append(*ret, root)
		}
		traversalForRoot(root.Left, target, ret)
		traversalForRoot(root.Right, target, ret)
	}
}

// 判断两棵树是不是一致
func traversalIsSubTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil && t != nil {
		return false
	}
	if s != nil && t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return traversalIsSubTree(s.Left, t.Left) && traversalIsSubTree(s.Right, t.Right)
}
