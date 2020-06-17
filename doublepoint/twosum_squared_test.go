package doublepoint

import (
	"fmt"
	"math"
	"testing"
)

// 给定一个非负数c，判断是否有整数a，b使得 a^2+b^2=c easy
// https://leetcode-cn.com/problems/sum-of-square-numbers/description/

func Test_TwoSumSquared(t *testing.T) {
	c := 9
	t.Log(twoSumSquared(c))
}

// 同有序数组两数和，b取sqrt(c)即可 时间负载度n
func twoSumSquared(c int) bool {
	if c < 0 {
		return false
	}
	a := 0
	b := int(math.Sqrt(float64(c)) + 1)
	for {
		target := a*a + b*b
		if target == c {
			fmt.Println(a, b)
			return true
		} else if target > c {
			b--
		} else {
			a++
		}
		if a > b {
			return false
		}
	}
}
