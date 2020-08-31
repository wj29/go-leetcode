package tree

import (
	"sort"
	"testing"
)

// 173. 二叉搜索树迭代器
// https://leetcode-cn.com/problems/binary-search-tree-iterator/
func Test_BSTIterator(t *testing.T) {
	root := NewTree()
	ret := ConstructorBSTIterator(root)
	t.Log(ret.Next())
	t.Log(ret.HasNext())
}

type BSTIterator struct {
	nums  []int // 顺序值
	index int   // 当前
}

func ConstructorBSTIterator(root *TreeNode) BSTIterator {
	r := BSTIterator{
		nums:  make([]int, 0),
		index: -1,
	}
	traversal5(root, &r.nums)
	sort.Ints(r.nums)
	return r
}

func traversal5(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	*nums = append(*nums, root.Val)
	traversal5(root.Left, nums)
	traversal5(root.Right, nums)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.index++
	return this.nums[this.index]
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.nums)-1
}
