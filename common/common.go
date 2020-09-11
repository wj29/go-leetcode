package common

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Sum(nums []int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		ret += nums[i]
	}
	return ret
}
