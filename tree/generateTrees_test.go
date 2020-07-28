package tree

import (
	"testing"
)

// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树
// https://leetcode-cn.com/problems/unique-binary-search-trees/ medium
// https://leetcode-cn.com/problems/unique-binary-search-trees-ii/  medium
func Test_GenerateTrees(t *testing.T) {
	n := 3
	t.Log(generateTreesNum(n))
	r := generateTrees(3)
	for _, v := range r {
		t.Log(v.Val, v.Left, v.Right)
	}
	ret := generateTreesByBackTrack(3)
	for _, v := range ret {
		t.Log(v.Val, v.Left, v.Right)
	}
}

// 设f(n)为以n为根结点的二叉树个数，则dp[n] = f(1)+f(2)+...+f(n)
// f(i)的以i为根结点，那么左子树个数是[1,i-1]，它的dp为dp[i-1]，右子树为[i,n-i]，dp为dp[n-i]
// 动态方程为 dp[n] = dp[0]*dp[n-1]+dp[1]*dp[n-2]*...*dp[n-1]*dp[0]
// 可以知道共有多少种生成的二叉搜索树
// 符合上述方程式的数为卡特兰数
func generateTreesNum(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < len(dp); i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// 根据上面的结论，我们知道f(i)是以i为根结点的二叉树，那么左右子树的列表分别是1到i-1和i+1到n构成的子树列表
// 循环两个列表即可以得到f(i)的所有情况
// 在i+1到n的子树列表和对应到0-n-i即可，需要加上偏差
// 那么n个数组成的二叉树，即分别以1-n为根结点的二叉树列表相加
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	dp := make([][]*TreeNode, n+1)
	dp[0], dp[1] = []*TreeNode{nil}, []*TreeNode{{Val: 1}}
	for i := 2; i < len(dp); i++ {
		for j := 1; j <= i; j++ {
			for _, l := range dp[j-1] {
				for _, r := range dp[i-j] {
					root := &TreeNode{Val: j}
					root.Left = l
					root.Right = addDiff(r, j)
					dp[i] = append(dp[i], root)
				}
			}
		}
	}
	return dp[n]
}

func addDiff(root *TreeNode, diff int) *TreeNode {
	if root == nil {
		return nil
	}
	ret := &TreeNode{
		Val: root.Val + diff,
	}
	ret.Left = addDiff(root.Left, diff)
	ret.Right = addDiff(root.Right, diff)
	return ret
}

// 回溯法
func generateTreesByBackTrack(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return backTrackForGenerateTree(1, n)
}

func backTrackForGenerateTree(start, end int) []*TreeNode {
	ret := make([]*TreeNode, 0)
	if start > end {
		ret = append(ret, nil)
		return ret
	}
	if start == end {
		ret = append(ret, &TreeNode{Val: start})
		return ret
	}
	for i := start; i <= end; i++ {
		//root := &TreeNode{Val: i}                    // 以i为跟节点 放在这里在下面循环时会重复
		left := backTrackForGenerateTree(start, i-1) // 左子树的列表
		right := backTrackForGenerateTree(i+1, end)  // 右子树的列表
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i}
				root.Left = l
				root.Right = r
				ret = append(ret, root)
			}
		}
	}
	return ret
}
