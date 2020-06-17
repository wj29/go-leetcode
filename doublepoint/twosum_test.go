package doublepoint

import "testing"

// 有序数组的两数和
// 从给定有序数组中找到两个数使其和为目标值 easy
// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/description/

func Test_TwoSum(t *testing.T) {
	arr := []int{3, 5, 2, 1, 6, 9, 10}
	sum := 11
	a, b := twoSum(arr, sum)
	t.Log(a, b)
	arr = []int{1, 2, 3, 4, 5, 6, 7}
	sum = 6
	c := twoSumPoint(arr, sum)
	t.Log(c)
}

// 将数组对应的值与目标的差值作为key，下标作为value，当数组的另一个值为map的key时，返回该该的下标和key对应的value
func twoSum(arr []int, sum int) (int, int) {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		v, ok := m[arr[i]]
		if ok {
			return v, i
		}
		m[sum-arr[i]] = i
	}
	return 0, 0
}

// 使用双向指针，双端逼近，和较小移动较小元素，和较大移动较大元素 时间负载度n
func twoSumPoint(arr []int, sum int) []int {
	if len(arr) <= 2 {
		return []int{0, 0}
	}
	l := 0
	r := len(arr) - 1
	for {
		if arr[l]+arr[r] == sum {
			return []int{l, r}
		} else if arr[l]+arr[r] > sum {
			r--
		} else {
			l++
		}
		if l >= r {
			return []int{0, 0}
		}
	}
}
