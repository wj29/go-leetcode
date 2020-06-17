package tree

import (
	"testing"
)

// 给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL
// https://leetcode-cn.com/problems/search-in-a-binary-search-tree/  easy

// 给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 保证原始二叉搜索树中不存在新值。
//注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回任意有效的结果。
// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree medium
func Test_InsertIntoBST(t *testing.T) {
	root := NewBSTree()
	val := 2
	t.Log(searchBST(root, val))

	root1 := NewSmallBSTree()
	newTree := insertIntoBST(root1, 8)
	t.Log(newTree)
}

// 二叉搜索树左<中<右
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

func searchBST1(root *TreeNode, val int) *TreeNode {
	var (
		res   []int
		stack []*TreeNode
	)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1
		res = append(res, stack[index].Val)
		if stack[index].Val == val {
			return stack[index]
		} else if stack[index].Val > val {
			return nil
		}
		root = stack[index].Right
		stack = stack[:index]
	}
	return nil
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
