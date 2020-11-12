package greedy

import (
    "sort"
    "testing"
)

// 406. 根据身高重建队列
// https://leetcode-cn.com/problems/queue-reconstruction-by-height/
func Test_ReconstructQueue(t *testing.T) {
    people := [][]int{
        {7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
    }
    t.Log(reconstructQueue(people))
}

// 身高从高到低排序的好处是，对于前面已经排好的队，
// 1.如果下一个人(h,k)比前面所有人都矮，那么，他插入队列的k处，使其达到k的要求，对其他人没影响
// 2.如果下一个人跟之前排好队的人中最矮的身高一样，这时候，就体现为什么之前排序时候，先考虑身高，再按照k的升序了，
// 这时候，新来的人虽然与之前最矮之人身高一样，但是由于他的k比之前最矮的人的k都大，所以，他插入的地方一定在已经排好队的，和他身高一样的，最矮之人的后面，
// 对这些最矮人们没有影响，当然，对其他比他高的人就更没有影响了。
func reconstructQueue(people [][]int) [][]int {
    // 身高倒序，k正序排列
    sort.SliceStable(people, func(i, j int) bool {
        if people[i][0] != people[j][0] {
            return people[i][0] > people[j][0]
        }
        return people[i][1] < people[j][1]
    })

    for i := 1; i < len(people); i++ {
        // 插入k的位置即可，其他人向后平移
        tmp := people[i]
        for j := i; j > tmp[1]; j-- {
            people[j] = people[j-1]
        }
        people[tmp[1]] = tmp
    }
    return people
}
