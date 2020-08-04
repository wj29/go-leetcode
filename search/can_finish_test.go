package search

import (
	"testing"
)

// 你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
// 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
// https://leetcode-cn.com/problems/course-schedule

func Test_CanFinish(t *testing.T) {
	prerequisites := [][]int{{0, 1}, {0, 2}, {1, 2}}
	numCourses := 3
	t.Log(canFinish(numCourses, prerequisites))
	t.Log(canFinish1(numCourses, prerequisites))
	t.Log(canFinish2(numCourses, prerequisites))
}

// 二维数组表示前置学习条件，深度搜索优先，如果有环则无法完成
func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 {
		return true
	}
	if len(prerequisites) == 0 || len(prerequisites[0]) == 0 {
		return true
	}

	var
	(
		visit = make([]bool, numCourses)
		edges = make([][]int, numCourses) // 图的对应边
		dfs   = func(int, *bool) {}
		flag  bool // 是否有环
	)

	for _, p := range prerequisites { // 构建图，找到所有节点的下一个节点列表
		edges[p[1]] = append(edges[p[1]], p[0])
	}
	dfs = func(start int, flag *bool) {
		if *flag {
			return
		}
		for i := 0; i < len(edges[start]); i++ {
			if visit[edges[start][i]] { // 它的子节点已经被访问过，说明有环
				*flag = true
				return
			}
			visit[edges[start][i]] = true
			dfs(edges[start][i], flag) // 访问所有下一节点
			visit[edges[start][i]] = false
		}
	}
	for i := 0; i < numCourses && !flag; i++ { // 此处对不对visit[i]标记不影响，如果有环，在dfs中会标记到
		dfs(i, &flag)
	}
	return !flag
}

// 使用
func canFinish1(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 {
		return true
	}
	if len(prerequisites) == 0 || len(prerequisites[0]) == 0 {
		return true
	}

	var
	(
		stack  = make([]int, 0)
		edges  = make([][]int, numCourses) // 图的对应边
		preCnt = make(map[int]int)         // 前置节点数
		visit  = make(map[int]bool)
	)

	for _, p := range prerequisites { // 构建图，找到所有节点的下一个节点列表
		edges[p[1]] = append(edges[p[1]], p[0])
		preCnt[p[0]] ++ // prerequisites中不能出现{0,1}重复的数组，否则要做去重
	}
	for i := 0; i < numCourses; i++ {
		if preCnt[i] == 0 {
			stack = append(stack, i)
			visit[i] = true
		}
	}
	for len(stack) != 0 {
		index := len(stack) - 1
		current := stack[index]
		stack = stack[:index] // 出栈
		visit[current] = true
		numCourses--
		for _, v := range edges[current] {
			if visit[v] { // 搜索时发现下一个节点被访问过，说明出现了环
				return false
			}
			preCnt[v]--
			if preCnt[v] == 0 { //
				stack = append(stack, v)
			}
		}
		visit[current] = false
	}
	return numCourses == 0
}

// 广度优先搜索
func canFinish2(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 {
		return true
	}
	if len(prerequisites) == 0 || len(prerequisites[0]) == 0 {
		return true
	}

	var
	(
		queue  = make([]int, 0)
		edges  = make([][]int, numCourses) // 图的对应边
		preCnt = make(map[int]int)         // 前置节点数
	)

	for _, p := range prerequisites { // 构建图，找到所有节点的下一个节点列表
		edges[p[1]] = append(edges[p[1]], p[0])
		preCnt[p[0]] ++ // prerequisites中不能出现{0,1}重复的数组，否则要做去重
	}
	for i := 0; i < numCourses; i++ {
		if preCnt[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 若图没有环，则所有节点一定都入队并出队过
	// 若有环，则当queue空时仍有节点未遍历到，此时numCourses一定不是0
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:] // 出队列
		numCourses--      // 访问了一个节点
		for _, v := range edges[current] {
			preCnt[v]--
			if preCnt[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return numCourses == 0
}
