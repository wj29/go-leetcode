package search

import "testing"

// 841. 钥匙和房间
// https://leetcode-cn.com/problems/keys-and-rooms/
func Test_CanVisitAllRooms(t *testing.T) {
	rooms := [][]int{
		{1}, {2}, {3}, {},
	}
	t.Log(canVisitAllRooms(rooms))
}

func canVisitAllRooms(rooms [][]int) bool {
	stack := make([]int, 0)     // 栈实现深度优先搜索
	visit := make(map[int]bool) // 记录访问过的房间
	count := make(map[int]bool) // 总房间数,相同钥匙去重
	for _, v := range rooms[0] {
		stack = append(stack, v)
	}
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			if count[rooms[i][j]] {
				continue
			}
			count[rooms[i][j]] = true
		}
	}
	for len(stack) != 0 {
		index := len(stack) - 1
		current := stack[index]
		stack = stack[:index]
		if visit[current] == true {
			continue
		}
		visit[current] = true
		for _, v := range rooms[current] {
			stack = append(stack, v)
		}
	}
	return len(count) == 0 || len(visit) == len(count)
}
