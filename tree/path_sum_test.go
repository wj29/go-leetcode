package tree

import (
	"testing"
)

// 路径和1
// 给定一个二叉树和一个目标和，判断该树中是否有从根结点到叶子结点路径，使得其和等于目标和
// https://leetcode-cn.com/problems/path-sum/description/  easy

// 路径和2
// 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径
// https://leetcode-cn.com/problems/path-sum-ii/  medium

// 路径和3
// 给定一个二叉树，它的每个结点都存放着一个整数值，找出路径和等于给定数值的路径总数
// 路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的(只能从父节点到子节点)
// 二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数
// https://leetcode-cn.com/problems/path-sum-iii/description/ easy
func Test_PathSum(t *testing.T) {
	root := NewTree()
	sum := 7
	t.Log(hasPathSum(root, sum))
	t.Log(pathSum2(root, sum))
	sum = 6
	t.Log(pathSum3(root, sum))

}

// 路径和1
// 从根结点遍历，保存所有叶子结点的路径
// 对于左右子树，它的父节点到跟结点的路径是一定的 可以通过栈实现
// 从根结点遍历，对所有左右子结点判断值是不是sum-父节点，如果是判断子结点是不是叶子结点，不是往下遍历
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil {
		return sum == 0
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

// 路径和2
// 遍历所有叶子结点，看看其路径和是否为sum，保存所有孩子结点到父节点的映射
// 得到从叶子结点到根结点路径，翻转即可得到结果
func pathSum2(root *TreeNode, sum int) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			ret = append(ret, []int{root.Val})
		}
		return ret
	}
	res := make([]*TreeNode, 0) // 满足条件的叶子结点
	m := make(map[*TreeNode]*TreeNode)
	path2(root, sum, &res, m)
	if len(m) == 0 {
		return ret
	}
	for i := range res {
		current := res[i]
		tmp := make([]int, 0)
		for current != nil {
			tmp = append(tmp, current.Val)
			current = m[current]
		}
		for j := 0; j < len(tmp)/2; j++ { // 翻转
			tmp[j], tmp[len(tmp)-j-1] = tmp[len(tmp)-j-1], tmp[j]
		}
		ret = append(ret, tmp)
	}
	return ret
}

func path2(root *TreeNode, sum int, res *[]*TreeNode, m map[*TreeNode]*TreeNode) {
	if root == nil {
		return
	}
	sum -= root.Val
	if root.Left == nil && root.Right == nil { // 判断是否为叶子结点
		if sum == 0 { // 满足
			*res = append(*res, root)
		} else {
			return
		}
	}
	if root.Left != nil { // 维护映射
		m[root.Left] = root
	}
	if root.Right != nil {
		m[root.Right] = root
	}
	path2(root.Left, sum, res, m)
	path2(root.Right, sum, res, m)
}

// 路径和3
// 路径只能从父结点到子结点
// 路径总数，所有结点必须全部遍历到
func pathSum3(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	count := 0
	p := &count // 维护一个全局变量count 发现一个路径和等于p则p++
	traversal(root, p, sum)
	return count
}

// 遍历所有的结点
func traversal(root *TreeNode, p *int, sum int) {
	if root == nil {
		return
	}
	path3(root, p, sum)
	traversal(root.Left, p, sum)
	traversal(root.Right, p, sum)
}

// 单结点到叶子结点的值是不是等于sum
// 受启发于hash_path_sum
func path3(root *TreeNode, p *int, sum int) {
	if root == nil {
		return
	}
	if root.Val == sum {
		*p++
	}
	sum -= root.Val
	path3(root.Left, p, sum)
	path3(root.Right, p, sum)
}
