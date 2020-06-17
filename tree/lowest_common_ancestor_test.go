package tree

import "testing"

// 给定一个二叉树，找到指定两个结点的最近公共祖先
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出: 3 节点 5 和节点 1 的最近公共祖先是节点 3。
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree medium
func Test_LowestCommonAncestor(t *testing.T) {
	root := NewTree()
	p := root.Left.Left
	q := root.Right
	t.Log(lowestCommonAncestor(root, p, q))
}

// 二叉树不可由子结点找到父节点，那么需要构建一个从子结点到父节点的路径
// 找到p和q的祖先，一一遍历即可
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	m := make(map[int]*TreeNode) // 子结点值为key，父节点为value
	traversal2(root, m)
	visited := make(map[*TreeNode]bool)
	for p != nil {
		visited[p] = true
		p = m[p.Val]
	}
	for q != nil {
		if visited[q] == true {
			return q
		}
		q = m[q.Val]
	}
	return nil
}

func traversal2(root *TreeNode, m map[int]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		m[root.Left.Val] = root
		traversal2(root.Left, m)
	}
	if root.Right != nil {
		m[root.Right.Val] = root
		traversal2(root.Right, m)
	}
}
