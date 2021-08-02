package dynamicprogramming

import (
    "testing"
)

func Test_NumSubarraysWithSum(t *testing.T) {
    nums := []int{1, 0, 1, 0, 1}
    goal := 2
    //nums = []int{0, 0, 0, 0, 0}
    //goal = 0
    //t.Log(numSubarraysWithSum(nums, goal))
    t.Log(numSubarraysWithSum2(nums, goal))
}

// 暴力解法
func numSubarraysWithSum(nums []int, goal int) int {
    if len(nums) == 0 {
        return 0
    }
    preSum := make([]int, len(nums))
    dp := make([][]int, len(nums))
    for i := 0; i < len(nums); i++ {
        dp[i] = make([]int, len(nums))
        if i == 0 {
            preSum[i] = nums[i]
            if nums[i] == goal {
                dp[0][0] = 1
            }
        } else {
            preSum[i] = preSum[i-1] + nums[i]
        }
    }

    for i := 1; i < len(nums); i++ {
        dp[0][i] += dp[0][i-1]
        if preSum[i] == goal {
            dp[0][i] ++
        }
        for k := i - 1; k >= 0; k-- {
            if preSum[i]-preSum[k] == goal {
                dp[0][i] ++
            }
        }
    }
    return dp[0][len(nums)-1]
}

// 计算nums的前缀和数组sum，然后从前往后扫描nums，对于右端点r，通过前缀和数组可以在 O(1)复杂度内求得 [0, r]连续一段的和，
// 根据容斥原理，我们还需要求得某个左端点l，使得[0,r] 减去[0,l−1] 和为t，即满足sum[r]−sum[l−1]=t，
// 这时候利用哈希表记录扫描过的sum[i] 的出现次数，可以实现O(1) 复杂度内求得满足要求的左端点个数
func numSubarraysWithSum2(nums []int, goal int) int {
    m := make(map[int]int)
    ans := 0
    sum := 0
    for _, num := range nums {
        m[sum]++
        sum += num
        ans += m[sum-goal]
    }
    return ans
}
