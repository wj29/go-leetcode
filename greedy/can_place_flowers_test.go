package greedy

import "testing"

// 假设你有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花卉不能种植在相邻的地块上，它们会争夺水源，两者都会死去
//  给定一个花坛（表示为一个数组包含0和1，其中0表示没种植花，1表示种植了花），和一个数n
// 能否在不打破种植规则的情况下种入 n 朵花？能则返回True，不能则返回False
// https://leetcode-cn.com/problems/can-place-flowers easy
func Test_CanPlaceFlowers(t *testing.T) {
	nums := []int{1, 0, 0, 0, 1}
	n := 1
	t.Log(canPlaceFlowers(nums, n))
}

// 每两个1之间可以加入多少个1，总和>=n即为true
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	f := make([]int, 0)
	f = append(f, -2)
	for i, v := range flowerbed {
		if v == 1 {
			f = append(f, i)
		}
	}
	f = append(f, len(flowerbed)+1)
	for i := 0; i < len(f)-1; i++ {
		count += (f[i+1]-f[i])/2 - 1
	}
	return count >= n
}

// 每个0的左右如果都是0，那么这个位置可以种花