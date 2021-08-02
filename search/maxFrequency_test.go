package search

import (
    "github.com/wujiang1994/go-leetcode/common"
    "sort"
    "testing"
)

func Test_MaxFrequency(t *testing.T) {
    nums := []int{1, 4, 8, 13}
    k := 5
    //nums = []int{1, 2, 4}
    k = 5
    //nums = []int{3, 9, 6}
    //k = 2
    t.Log(maxFrequency(nums, k))
}

// 设ans为数组nums[i:j]所有数据变成nums[j-1]的操作次数，那么ans<=k，频次即为i-j的跨度，即j-i+1
// 针对已排序的nums中的数据，从i-j的的操作次数<=从i-(j+1)的操作次数
// i-j的操作次数满足ans<=k时，j++，增大窗口直到i-j的操作次数>k，此时不能增大j了，增大j操作次数会更多，不会满足
// 此时增大i缩小窗口，直到i-j的操作次数再次满足ans<=k，j++重复上述过程，找到最大窗口
func maxFrequency(nums []int, k int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    sort.Ints(nums)
    ans := 1
    window := 0
    for l, r := 0, 1; r < len(nums); r++ {
        window += (nums[r] - nums[r-1]) * (r - l)
        for window > k {
            window -= nums[r] - nums[l]
            l++
        }
        ans = common.Max(ans, r-l+1)
    }
    return ans
}
