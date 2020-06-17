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
// 那么n个数组成的二叉树，即分别以1-n为根结点的二叉树列表相加
// ... how to implement?
func generateTrees(n int) []*TreeNode {
	return []*TreeNode{}
}
