package greedy

import (
	"sort"
	"testing"
)

// 给定一个区间的集合，找到需要移除区间的最小数量，使得剩余区间不重叠 medium
// https://leetcode-cn.com/problems/non-overlapping-intervals/description/
// [1,2] [2,3]接触不重叠
// eg: 输入 [ [1,2], [2,3], [3,4], [1,3] ] 输出 1(移除[1,3])
// eg: 输入 [ [1,2], [1,2], [1,2] ] 输出2(移除两个[1,2])

func Test_EraseOverlapIntervals(t *testing.T) {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	intervals = [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	t.Log(eraseOverlapIntervals(intervals))
}

// 贪心算法 局部最优
// 假定已存在不重叠的区域，判断后序加入的区域会不会让已不重叠的区域重叠，未加入个数即是移除的最小个数(不会实现)

// 如何实现上面的算法，以下借鉴leetcode题解
// 1.从区间集合中选end最小的区间x，放入新的集合，end为集合中最后一个区间的end
// 2.把所有与x相交的区间从集合中删除，start<end，更新集合中下一个区间到新的集合中
// 3.重复1和2直到集合为空，新的集合为最大不相交子集
// https://leetcode-cn.com/problems/non-overlapping-intervals/solution/tan-xin-suan-fa-zhi-qu-jian-diao-du-wen-ti-by-labu/
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 0 {
		return 0
	}
	// 排序
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][len(intervals[i])-1] < intervals[j][len(intervals[j])-1]
	})
	r := make([][]int, 0)
	r = append(r, intervals[0])
	for i := range intervals {
		end := r[len(r)-1][len(r[len(r)-1])-1] // 最后一个元素的end
		if intervals[i][0] >= end {
			r = append(r, intervals[i])
		}
	}
	return len(intervals) - len(r)
}

// 给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。
// 现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。
// 给定一个对数集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。
// https://leetcode-cn.com/problems/maximum-length-of-pair-chain  medium
// 最大不相交子集
func findLongestChain(pairs [][]int) int {
	if len(pairs) <= 0 {
		return 0
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	ret := make([][]int, 0)
	ret = append(ret, pairs[0])
	for i := range pairs {
		end := ret[len(ret)-1][len(ret[len(ret)-1])-1]
		if pairs[i][0] > end {
			ret = append(ret, pairs[i])
		}
	}
	return len(ret)
}
