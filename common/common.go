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
