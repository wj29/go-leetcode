package tree

import (
	"testing"
)

// 根据一棵树的前序遍历与中序遍历构造二叉树
// 你可以假设树中没有重复的元素
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/ medium
// 根据一棵树的中序遍历与后序遍历构造二叉树
// 你可以假设树中没有重复的元素
// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/ medium
func Test_BuildTree(t *testing.T) {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	post := []int{9, 15, 7, 20, 3}
	root1 := buildTree1(pre, in)
	t.Log(root1)
	root2 := buildTree2(in, post)
	t.Log(root2)
}

// 二叉树的前序遍历 [根 [左子树的遍历结果] [右子树的遍历结果] ]
// 二叉树的中序遍历 [[左子树的遍历结果] 根 [右子树的遍历结果] ]
// 我们找到根结点，将左右子树分别接上，而左右子树又可以通过递归知道
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}
	root.Left = buildTree1(preorder[1:index+1], inorder[:index])
	root.Right = buildTree1(preorder[index+1:], inorder[index+1:])
	return root
}

// 二叉树的中序遍历 [[左子树的遍历结果] 根 [右子树的遍历结果] ]
// 二叉树的后序遍历 [[左子树的遍历结果] [右子树的遍历结果] 根 ]
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}
	root.Left = buildTree2(inorder[:index], postorder[:index])
	root.Right = buildTree2(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
