package greedy

import "testing"

// 134. 加油站
// https://leetcode-cn.com/problems/gas-station/
func Test_CanCompleteCircuit(t *testing.T) {
    gas := []int{1, 2, 3, 4, 5}
    cost := []int{3, 4, 5, 1, 2}
    t.Log(canCompleteCircuit(gas, cost))
}

// O(N^2)f]时间复杂度
func canCompleteCircuit(gas []int, cost []int) int {
    diff := make([]int, len(gas))
    for i := 0; i < len(gas); i++ {
        diff[i] = gas[i] - cost[i]
    }
    for i := 0; i < len(diff); i++ {
        if diff[i] < 0 {
            continue
        }
        tmp := diff[i]
        j := i
        for {
            if j == len(diff)-1 {
                j = -1
            }
            tmp += diff[j+1]
            if tmp < 0 {
                break
            }
            j++
            if j == i {
                return j
            }
        }
    }
    return -1
}

// 一次遍历
// https://leetcode-cn.com/problems/gas-station/solution/jia-you-zhan-by-leetcode/
func canCompleteCircuit2(gas []int, cost []int) int {
    rest, run, start := 0, 0, 0
    for i := 0; i < len(gas); i++ {
        run += gas[i] - cost[i]
        rest += gas[i] - cost[i]
        if run < 0 {
            start = i + 1
            run = 0
        }
    }
    if rest < 0 {
        return -1
    }
    return start
}
