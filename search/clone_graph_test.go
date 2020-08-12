package search

import (
	"fmt"
	"testing"
)

// 给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）
// 图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）
// https://leetcode-cn.com/problems/clone-graph/

func Test_CloneGraph(t *testing.T) {
	node1 := &Node{
		Val:       1,
		Neighbors: make([]*Node, 0),
	}
	node2 := &Node{
		Val:       2,
		Neighbors: make([]*Node, 0),
	}
	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node2.Neighbors, node1)
	ret := cloneGraph(node1)
	fmt.Println(ret, ret.Neighbors[0].Neighbors)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

// use queue for BFS
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visit := make(map[int]bool)
	m := make(map[int]*Node)
	queue := make([]*Node, 0)
	queue = append(queue, node)
	m[node.Val] = &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]       // out of the queue
		if visit[current.Val] { // has been visited
			continue
		}
		visit[current.Val] = true // set as visited

		for _, neighbor := range current.Neighbors {
			tmp := new(Node) // get or create tmp
			if _, ok := m[neighbor.Val]; ok {
				tmp = m[neighbor.Val]
			} else {
				tmp = &Node{
					Val:       neighbor.Val,
					Neighbors: make([]*Node, 0),
				}
			}
			m[current.Val].Neighbors = append(m[current.Val].Neighbors, tmp) // append to new node neighbors
			m[neighbor.Val] = tmp                                            // write to map
			queue = append(queue, neighbor)                                  // append to queue
		}
	}
	return m[node.Val]
}
