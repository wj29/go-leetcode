package dynamicprogramming

import (
	"testing"
	"github.com/wujiang1994/go-leetcode/common"
)

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 问总共有多少条不同的路径？
// 这个问题又是一个排列组合
// https://leetcode-cn.com/problems/unique-paths/  medium

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径
// https://leetcode-cn.com/problems/unique-paths-ii medium

// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小
// https://leetcode-cn.com/problems/minimum-path-sum/description/  medium
func Test_Chess(t *testing.T) {
	//t.Log("递归:", uniquePaths(51, 9)) // 算不出来，超时
	//t.Log("动态规划:", process1(51, 9))
	//obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	//t.Log(uniquePathsWithObstacles(obstacleGrid))
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	t.Log(minPathSum(grid))
	t.Log(maxValue(grid))
}

// 每一步达到的路径只有它的下方或者左方，所以它的值是它左方的走法加上下方的走法
func uniquePaths(m, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	return process(m, n)
}

func process(m, n int) int {
	if m == 1 && n == 1 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	return process(m-1, n) + process(m, n-1)
}

// 动态规划我们知道m*n的格子存储的大小应该为m*n的二维数组
func process1(m, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				arr[i][j] = 0
			} else if i == 0 || j == 0 {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i][j-1] + arr[i-1][j]
			}
		}
	}
	return arr[m-1][n-1]
}

// 到达右下角的路径只能由其上方或者左方到达，如果其上方或者左方为1即有障碍，那么右下角永远无法到达
// 那么 dp[m][n] = dp[m-1][n] + dp[m][n-1]
// 如果一个点是障碍物，那么改点的dp[i][j]=0
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	if len(obstacleGrid) == 1 && len(obstacleGrid[0]) == 1 {
		if obstacleGrid[0][0] == 0 {
			return 1
		}
		return 0
	}
	dp := make([][]int, len(obstacleGrid))
	for i := range dp {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	// 处理base
	dp[0][0] = 0
	flag := false
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			flag = true
		}
		if flag {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}
	}
	flag = false
	for j := 0; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] == 1 {
			flag = true
		}
		if flag {
			dp[0][j] = 0
		} else {
			dp[0][j] = 1
		}
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

// 同样的做法dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
// base为横竖第一行
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	if len(grid) == 1 && len(grid[0]) == 1 {
		return grid[0][0]
	}
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	// 处理base
	for i := range dp {
		if i == 0 {
			dp[0][0] = grid[0][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for i := range dp[0] {
		if i == 0 {
			continue
		}
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = common.Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

// 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0
// 你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角
// 给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物
// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof  medium
func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	if len(grid) == 1 && len(grid[0]) == 1 {
		return grid[0][0]
	}
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	// 处理base
	dp[0][0] = grid[0][0]
	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = common.Max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
