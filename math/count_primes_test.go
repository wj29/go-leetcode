package math

import (
    "testing"
)

// 204. 计数质数
// https://leetcode-cn.com/problems/count-primes/
func Test_CountPrimes(t *testing.T) {
    n := 10
    t.Log(countPrimes(n))
}

func countPrimes(n int) int {
    isPrime := make([]bool, n)
    for i := range isPrime {
        isPrime[i] = true
    }
    cnt := 0
    for i := 2; i < n; i++ {
        if isPrime[i] {
            cnt++
            for j := 2 * i; j < n; j += i {
                isPrime[j] = false
            }
        }
    }
    return cnt
}
