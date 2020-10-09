package search

import (
	"github.com/wujiang1994/go-leetcode/common"
	"testing"
)

// 127. 单词接龙
// https://leetcode-cn.com/problems/word-ladder/
func Test_LadderLength(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	t.Log(ladderLength(beginWord, endWord, wordList))
	t.Log(ladderLength2(beginWord, endWord, wordList))
}

// 回溯法(超时)(深度优先搜索)
// 使用过的词不需要使用，重复过程无意义
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if common.StringInSlice(wordList, endWord) < 0 {
		return 0
	}
	if beginWord == endWord {
		return 0
	}
	visited := make(map[string]bool)
	visited[beginWord] = true
	ret := len(wordList) + 1
	backTrackForLadderLength(wordList, visited, &ret, 0, beginWord, endWord)
	if ret > len(wordList) {
		return 0
	}
	return ret + 1
}

func backTrackForLadderLength(wordList []string, visited map[string]bool, ret *int, count int, beginWord, endWord string) {
	if beginWord == endWord {
		*ret = common.Min(*ret, count)
		return
	}
	for i := 0; i < len(wordList); i++ {
		if visited[wordList[i]] || !diffOne(beginWord, wordList[i]) {
			continue
		}
		visited[wordList[i]] = true
		backTrackForLadderLength(wordList, visited, ret, count+1, wordList[i], endWord)
		visited[wordList[i]] = false
	}
}

func diffOne(s, t string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}

// 广度优先搜索
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	if common.StringInSlice(wordList, endWord) < 0 {
		return 0
	}
	if beginWord == endWord {
		return 0
	}
	visited := make(map[string]bool)
	changeMapping := make(map[string][]string)
	wordList = append(wordList, beginWord)
	for i := 0; i < len(wordList); i++ {
		changeMapping[wordList[i]] = make([]string, 0)
		for j := 0; j < len(wordList); j++ {
			if wordList[i] == wordList[j] {
				continue
			}
			if diffOne(wordList[i], wordList[j]) {
				changeMapping[wordList[i]] = append(changeMapping[wordList[i]], wordList[j])
			}
		}
	}
	wordList = wordList[:len(wordList)-1]
	queue := make([]*lad, 0)
	queue = append(queue, &lad{
		val:   beginWord,
		count: 1,
	})
	visited[beginWord] = true
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		if current.val == endWord {
			return current.count
		}
		for _, v := range changeMapping[current.val] {
			if visited[v] {
				continue
			}
			visited[v] = true
			queue = append(queue, &lad{
				val:   v,
				count: current.count + 1,
			})
		}
	}
	return 0
}

type lad struct {
	val   string
	count int
}

// TODO
// 双端bfs，从头和尾都bfs，选择范围较小的一边进行bfs