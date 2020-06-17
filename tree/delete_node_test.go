package tree

import (
	"testing"
)

// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变
// 返回二叉搜索树（有可能被更新）的根节点的引用
// https://leetcode-cn.com/problems/delete-node-in-a-bst  medium
// 要求算法时间复杂度为 O(h)，h 为树的高度
func Test_DeleteNode(t *testing.T) {
	root := NewBSTree()
	//root = &TreeNode{
	//	Val:   2,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root.Left = &TreeNode{
	//	Val:   1,
	//	Left:  nil,
	//	Right: nil,
	//}
	//root.Right = &TreeNode{
	//	Val:   3,
	//	Left:  nil,
	//	Right: nil,
	//}
	head := deleteNode(root, 2)
	t.Log(head)
}

// BST的对应右子树的值比根结点大，删除某个结点后，将其左子树移到其右子树的最左左子树的左结点即可
// 左 < 根 < 右 删除根后，要找到比根大的最小的结点才不会破坏BST
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		min := root.Right
		for min.Left != nil {
			min = min.Left
		}
		min.Left = root.Left
		root = root.Right
	}
	return root
}
