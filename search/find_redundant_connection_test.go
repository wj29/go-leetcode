package search

import (
	"testing"
)

// 684. 冗余连接
// https://leetcode-cn.com/problems/redundant-connection/
// 685. 冗余连接 II
// https://leetcode-cn.com/problems/redundant-connection-ii/
func Test_FindRedundantConnection(t *testing.T) {
	edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
	t.Log(findRedundantConnection(edges))
	t.Log(findRedundantDirectedConnection(edges))
}

// 并查集
// 一般用于处理一些不相交集合问题，例如连通子图、最近公共祖先、最小生成树
// 并查集的三个操作
// make  新建一个并查集，其中包含n个单元素的集合
// find  寻找元素x的代表，判断两个元素是否在统一集合只需要判断其代表是否相同即可
// union 将元素x与元素y所在的并查集合并，如果相交则不合并

// 路径压缩 在每次查找时，将路径上的每个节点指向根节点(代表)
// 按秩合并 将较小秩向较大秩的合并

// 无向图中有附加边时一定是形成了一个环
func findRedundantConnection(edges [][]int) []int {
	dsu := make([]int, len(edges)+1)
	for i := range dsu {
		dsu[i] = i
	}
	for i := 0; i < len(edges); i++ {
		x, y := edges[i][0], edges[i][1]
		if findBehalf(dsu, x) == findBehalf(dsu, y) {
			return edges[i]
		}
		union(dsu, x, y)
	}
	return nil
}

func findBehalf(dsu []int, x int) int {
	if dsu[x] != x {
		dsu[x] = findBehalf(dsu, dsu[x])
	}
	return dsu[x]
}

func union(dsu []int, x, y int) {
	xBehalf, yBehalf := findBehalf(dsu, x), findBehalf(dsu, y)
	if xBehalf == yBehalf {
		return
	}
	dsu[xBehalf] = yBehalf
	return
}

// 有向图中，题目中有一个节点会存在两个父节点
// 或者形成了一个环
// 官方题解
// https://leetcode-cn.com/problems/redundant-connection-ii/solution/rong-yu-lian-jie-ii-by-leetcode-solution/
func findRedundantDirectedConnection(edges [][]int) []int {
	dsu := make([]int, len(edges)+1)
	for i := range dsu {
		dsu[i] = i
	}
	parent := make([]int, len(edges)+1)
	copy(parent, dsu)

	conflictEdge, cycleEdge := make([]int, 0), make([]int, 0)
	for i := 0; i < len(edges); i++ {
		x, y := edges[i][0], edges[i][1]
		if parent[y] != y {
			conflictEdge = edges[i]
		} else {
			parent[y] = x
			if findBehalf(dsu, x) == findBehalf(dsu, y) {
				cycleEdge = edges[i]
			} else {
				union(dsu, x, y)
			}
		}
	}
	// 若不存在一个节点有两个父节点的情况，则附加的边一定导致环路出现
	if len(conflictEdge) == 0 {
		return cycleEdge
	}
	// conflictEdge[1] 有两个父节点，其中之一与其构成附加的边
	// 由于我们是按照 edges 的顺序连接的，若在访问到 conflictEdge 之前已经形成了环路，则附加的边在环上
	// 否则附加的边就是 conflictEdge
	if len(cycleEdge) != 0 {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge
}