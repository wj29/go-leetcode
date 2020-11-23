package greedy

import (
    "github.com/wujiang1994/go-leetcode/common"
    "sort"
    "testing"
)

// 621. 任务调度器
// https://leetcode-cn.com/problems/task-scheduler/
func Test_LeastInterval(t *testing.T) {
    tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
    tasks = []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
    n := 2
    t.Log(leastInterval(tasks, n))
    t.Log(leastInterval2(tasks, n))
}

// 以n+1为一轮，每次选择任务次数最多的n+1个任务安排
// 贪心的原则表示如果将出现次数多的任务后安排，会导致大量空闲时间的出现
func leastInterval(tasks []byte, n int) int {
    cnt := make([]int, 26)
    for _, v := range tasks {
        cnt[v-'A']++
    }
    sort.Ints(cnt)
    time := 0
    for cnt[25] > 0 {
        index := 0
        for index <= n { // n+1为一轮
            if cnt[25] == 0 {
                break
            }
            if index < 26 && cnt[25-index] > 0 {
                cnt[25-index]--
            }
            time++
            index++
        }
        sort.Ints(cnt)
    }
    return time
}

// 假设数组 ["A","A","A","B","B","C"]，n = 2，A的频率最高，记为count = 3，所以两个A之间必须间隔2个任务，
// 才能满足题意并且是最短时间（两个A的间隔大于2的总时间必然不是最短），因此执行顺序为： A->X->X->A->X->X->A，
// 这里的X表示除了A以外其他字母，或者是待命，不用关心具体是什么，反正用来填充两个A的间隔的。
// 上面执行顺序的规律是： 有count - 1个A，其中每个A需要搭配n个X，再加上最后一个A，所以总时间为 (count - 1) * (n + 1) + 1

// 要注意可能会出现多个频率相同且都是最高的任务，比如 ["A","A","A","B","B","B","C","C"]，所以最后会剩下一个A和一个B，因此最后要加上频率最高的不同任务的个数 maxCount

// 公式算出的值可能会比数组的长度小，如["A","A","B","B"]，n = 0，此时要取数组的长度
// 当maxCount > n 的情况，此时必然存在不需要待命的最短时间的执行顺序，所以此时的最短时间就是tasks.length
// https://leetcode-cn.com/problems/task-scheduler/comments/
func leastInterval2(tasks []byte, n int) int {
    cnt := make([]int, 26)
    for _, v := range tasks {
        cnt[v-'A']++
    }
    sort.Ints(cnt)
    max := cnt[25]
    maxCount := 1
    for i := 24; i >= 0; i-- {
        if cnt[i] == max {
            maxCount++
        }
    }
    ret := (cnt[25]-1)*(n+1) + maxCount
    return common.Max(ret, len(tasks))
}
