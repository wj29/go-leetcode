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
    //t.Log(ladderLength(beginWord, endWord, wordList))
    //t.Log(ladderLength2(beginWord, endWord, wordList))

    t.Log(findLadders(beginWord, endWord, wordList))
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

// 有问题 应先用bfs求出最短路径长 dfs求出所有路径
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    if common.StringInSlice(wordList, endWord) < 0 || beginWord == endWord {
        return nil
    }
    wordList = append(wordList, beginWord)
    changeMapping := make(map[string][]string)
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
    ans := make([][]string, 0)
    wordList = wordList[:len(wordList)-1]
    queue := make([]*ladd, 0)
    queue = append(queue, &ladd{
        val:   beginWord,
        count: 1,
        arr:   []string{beginWord},
    })
    visited := make(map[string]bool)
    visited[beginWord] = true
    flag := false

    for len(queue) != 0 {
        current := queue[0]
        queue = queue[1:]
        if current.val == endWord {
            if !flag {
                ans = append(ans, current.arr)
                flag = true
                continue
            }
            if len(current.arr) > len(ans[0]) {
                break
            }
            ans = append(ans, current.arr)
            continue
        }
        for _, v := range changeMapping[current.val] {
            if visited[v] {
                continue
            }
            visited[v] = true
            queue = append(queue, &ladd{
                val:   v,
                arr:   append(current.arr, v),
                count: current.count + 1,
            })
        }
    }
    return ans
}

type ladd struct {
    val   string
    count int
    arr   []string
}