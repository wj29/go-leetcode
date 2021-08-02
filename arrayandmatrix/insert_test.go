package arrayandmatrix

import (
    "github.com/wujiang1994/go-leetcode/common"
    "testing"
)

// 57. 插入区间
// https://leetcode-cn.com/problems/insert-interval/

func Test_Insert(t *testing.T) {
    intervals := [][]int{
        {1, 5},
        {6, 8},
    }
    newInterval := []int{3, 7}
    t.Log(insert(intervals, newInterval))
    t.Log(insert2(intervals, newInterval))

}

// 判断两个区间是否有交集
// 对于区间S1=[l1, r1] S2=[l2,r2] 如果没有交集则需要满足r1<l2 || l1>r2 即在其左侧或右侧
// 不满足上述条件则必相交
// 其交集为[max(l1,l2), min(r1,r2)] 并集为[min(l1,l2), max(r1,r2)]
func insert(intervals [][]int, newInterval []int) [][]int {
    ret := make([][]int, 0)
    merge := false
    left, right := newInterval[0], newInterval[1]
    for _, interval := range intervals {
        if interval[0] > right {
            if !merge {
                ret = append(ret, []int{left, right})
                merge = true
            }
            ret = append(ret, interval)
        } else if interval[1] < left {
            ret = append(ret, interval)
        } else {
            left = common.Min(left, interval[0])
            right = common.Max(right, interval[1])
        }
    }
    if !merge {
        ret = append(ret, []int{left, right})
    }
    return ret
}

// 先寻找左边界，再寻找右边界
func insert2(intervals [][]int, newInterval []int) [][]int {
    ret := make([][]int, 0)
    for i := 0; i < len(intervals); i++ {
        tmp := make([]int, 2)
        if newInterval[0] <= intervals[i][0] {
            tmp[0] = newInterval[0]
            if newInterval[1] < intervals[i][0] {
                tmp[1] = newInterval[1]
                ret = append(ret, tmp)
                ret = append(ret, intervals[i:]...)
                return ret
            } else if newInterval[1] == intervals[i][0] {
                tmp[1] = intervals[i][1]
                ret = append(ret, tmp)
                ret = append(ret, intervals[i+1:]...)
                return ret
            } else {
                if newInterval[1] <= intervals[i][1] {
                    tmp[1] = intervals[i][1]
                    ret = append(ret, tmp)
                    ret = append(ret, intervals[i+1:]...)
                    return ret
                } else {
                    for j := i + 1; j < len(intervals); j++ {
                        if newInterval[1] < intervals[j][0] {
                            tmp[1] = newInterval[1]
                            ret = append(ret, tmp)
                            ret = append(ret, intervals[j:]...)
                            return ret
                        } else if newInterval[1] == intervals[j][0] {
                            tmp[1] = intervals[j][1]
                            ret = append(ret, tmp)
                            ret = append(ret, intervals[j+1:]...)
                            return ret
                        } else {
                            if newInterval[1] <= intervals[j][1] {
                                tmp[1] = intervals[j][1]
                                ret = append(ret, tmp)
                                ret = append(ret, intervals[j+1:]...)
                                return ret
                            }
                        }
                    }
                    tmp[1] = newInterval[1]
                    ret = append(ret, tmp)
                    return ret
                }
            }
        } else {
            if newInterval[0] <= intervals[i][1] {
                tmp[0] = intervals[i][0]
                if newInterval[1] <= intervals[i][1] {
                    tmp[1] = intervals[i][1]
                    ret = append(ret, tmp)
                    ret = append(ret, intervals[i+1:]...)
                    return ret
                } else {
                    for j := i + 1; j < len(intervals); j++ {
                        if newInterval[1] < intervals[j][0] {
                            tmp[1] = newInterval[1]
                            ret = append(ret, tmp)
                            ret = append(ret, intervals[j:]...)
                            return ret
                        } else if newInterval[1] == intervals[j][0] {
                            tmp[1] = intervals[j][1]
                            ret = append(ret, tmp)
                            ret = append(ret, intervals[j+1:]...)
                            return ret
                        } else {
                            if newInterval[1] <= intervals[j][1] {
                                tmp[1] = intervals[j][1]
                                ret = append(ret, tmp)
                                ret = append(ret, intervals[j+1:]...)
                                return ret
                            }
                        }
                    }
                    tmp[1] = newInterval[1]
                    ret = append(ret, tmp)
                    return ret
                }
            } else {
                ret = append(ret, intervals[i])
                continue
            }
        }
    }
    return append(ret, newInterval)
}
