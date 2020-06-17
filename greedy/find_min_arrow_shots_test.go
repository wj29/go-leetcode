package greedy

import (
	"sort"
	"testing"
)

// 在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标
// 由于它是水平的，所以y坐标并不重要，因此只要知道开始和结束的x坐标就足够了。开始坐标总是小于结束坐标。平面内最多存在104个气球。
// 一支弓箭可以沿着x轴从不同点完全垂直地射出。在坐标x处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend，
// 且满足  xstart ≤ x ≤ xend，则该气球会被引爆。可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。
// 我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量
// https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/ medium

func Test_FindMinArrowShots(t *testing.T)  {
	
}

// 此题同求二维数组最大不相交子集，及射出的最少弓箭数量
// func eraseOverlapIntervals
func findMinArrowShots(points [][]int) int {
	if len(points) <= 0 {
		return 0
	}
	// 排序
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][len(points[i])-1] < points[j][len(points[j])-1]
	})
	r := make([][]int, 0)
	r = append(r, points[0])
	for i := range points {
		end := r[len(r)-1][len(r[len(r)-1])-1] // 最后一个元素的end
		if points[i][0] > end {
			r = append(r, points[i])
		}
	}
	return len(r)
}