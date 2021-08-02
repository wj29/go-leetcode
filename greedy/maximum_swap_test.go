package greedy

import (
    "math"
    "testing"
)

// 670. 最大交换
// https://leetcode-cn.com/problems/maximum-swap/
func Test_MaximumSwap(t *testing.T) {
    num := 2736
    num = 91
    t.Log(maximumSwap(num))
}

// 优先从左边交换，找到地位比高位高切最高的位进行交换得到的数最大
func maximumSwap(num int) int {
    arr := make([]int, 0)
    for num != 0 {
        arr = append(arr, num%10)
        num /= 10
    }
    for i := 0; i < len(arr)/2; i++ {
        arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
    }
    for i := 0; i < len(arr); i++ {
        maxIndex := -1
        for j := len(arr) - 1; j > i; j-- {
            if maxIndex < 0 && arr[j] > arr[i] {
                maxIndex = j
            }
            if maxIndex >= 0 && arr[j] > arr[maxIndex] {
                maxIndex = j
            }
        }
        if maxIndex >= 0 {
            arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
            break
        }
    }

    ret := 0
    for i := 0; i < len(arr); i++ {
        ret += arr[i] * int(math.Pow10(len(arr)-i-1))
    }
    return ret
}
