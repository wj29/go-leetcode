package backtracking

import (
	"sort"
	"testing"
)

// 332. 重新安排行程
// https://leetcode-cn.com/problems/reconstruct-itinerary/

func Test_FindItinerary(t *testing.T) {
	//tickets := [][]string{
	//	{"AXA", "EZE"}, {"EZE", "AUA"}, {"ADL", "JFK"}, {"ADL", "TIA"}, {"AUA", "AXA"}, {"EZE", "TIA"}, {"EZE", "TIA"},
	//	{"AXA", "EZE"}, {"EZE", "ADL"}, {"ANU", "EZE"}, {"TIA", "EZE"}, {"JFK", "ADL"}, {"AUA", "JFK"}, {"JFK", "EZE"},
	//	{"EZE", "ANU"}, {"ADL", "AUA"}, {"ANU", "AXA"}, {"AXA", "ADL"}, {"AUA", "JFK"}, {"EZE", "ADL"}, {"ANU", "TIA"},
	//	{"AUA", "JFK"}, {"TIA", "JFK"}, {"EZE", "AUA"}, {"AXA", "EZE"}, {"AUA", "ANU"}, {"ADL", "AXA"}, {"EZE", "ADL"},
	//	{"AUA", "ANU"}, {"AXA", "EZE"}, {"TIA", "AUA"}, {"AXA", "EZE"}, {"AUA", "SYD"}, {"ADL", "JFK"}, {"EZE", "AUA"},
	//	{"ADL", "ANU"}, {"AUA", "TIA"}, {"ADL", "EZE"}, {"TIA", "JFK"}, {"AXA", "ANU"}, {"JFK", "AXA"}, {"JFK", "ADL"},
	//	{"ADL", "EZE"}, {"AXA", "TIA"}, {"JFK", "AUA"}, {"ADL", "EZE"}, {"JFK", "ADL"}, {"ADL", "AXA"}, {"TIA", "AUA"},
	//	{"AXA", "JFK"}, {"ADL", "AUA"}, {"TIA", "JFK"}, {"JFK", "ADL"}, {"JFK", "ADL"}, {"ANU", "AXA"}, {"TIA", "AXA"},
	//	{"EZE", "JFK"}, {"EZE", "AXA"}, {"ADL", "TIA"}, {"JFK", "AUA"}, {"TIA", "EZE"}, {"EZE", "ADL"}, {"JFK", "ANU"},
	//	{"TIA", "AUA"}, {"EZE", "ADL"}, {"ADL", "JFK"}, {"ANU", "AXA"}, {"AUA", "AXA"}, {"ANU", "EZE"}, {"ADL", "AXA"},
	//	{"ANU", "AXA"}, {"TIA", "ADL"}, {"JFK", "ADL"}, {"JFK", "TIA"}, {"AUA", "ADL"}, {"AUA", "TIA"}, {"TIA", "JFK"},
	//	{"EZE", "JFK"}, {"AUA", "ADL"}, {"ADL", "AUA"}, {"EZE", "ANU"}, {"ADL", "ANU"}, {"AUA", "AXA"}, {"AXA", "TIA"},
	//	{"AXA", "TIA"}, {"ADL", "AXA"}, {"EZE", "AXA"}, {"AXA", "JFK"}, {"JFK", "AUA"}, {"ANU", "ADL"}, {"AXA", "TIA"},
	//	{"ANU", "AUA"}, {"JFK", "EZE"}, {"AXA", "ADL"}, {"TIA", "EZE"}, {"JFK", "AXA"}, {"AXA", "ADL"}, {"EZE", "AUA"},
	//	{"AXA", "ANU"}, {"ADL", "EZE"}, {"AUA", "EZE"},
	//}

	//tickets := [][]string{
	//	{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"},
	//}

	tickets := [][]string{
		{"DRW", "HBA"}, {"EZE", "DRW"}, {"ANU", "EZE"}, {"AXA", "EZE"}, {"DRW", "HBA"}, {"ANU", "SYD"}, {"EZE", "ANU"}, {"CNS", "DRW"}, {"HBA", "BNE"}, {"JFK", "CNS"}, {"BNE", "EZE"}, {"HBA", "EZE"}, {"EZE", "AXA"}, {"ANU", "TIA"}, {"CNS", "ANU"}, {"ADL", "CNS"}, {"TIA", "ANU"}, {"EZE", "ADL"},
	}

	t.Log(findItinerary(tickets))
}

func findItinerary(tickets [][]string) []string {
	edges := make(map[string][]string)
	visit := make(map[string]int)
	ret := ""
	flag := false
	for i := 0; i < len(tickets); i++ { // 构造图的边和访问数组
		edges[tickets[i][0]] = append(edges[tickets[i][0]], tickets[i][1])
		visit[tickets[i][0]+tickets[i][1]]++
	}
	for k := range edges {
		sort.Strings(edges[k])
	}
	backTrack(edges, (len(tickets)+1)*3, "JFK", visit, &ret, &flag)

	r := make([]string, 0)
	for i := 3; i <= len(ret); i += 3 {
		r = append(r, ret[i-3:i])
	}
	return r
}

func backTrack(edges map[string][]string, length int, pre string, visit map[string]int, ret *string, flag *bool) {
	if *flag { // 发现最短路径直接返回
		return
	}
	if len(pre) == length {
		*ret = pre
		*flag = true
		return
	}
	src := pre[len(pre)-3:]
	for i := 0; i < len(edges[src]); i++ {
		if i+1 < len(edges[src]) && edges[src][i] == edges[src][i+1] { // 同一层循环中相同的目的地只需要跑一次
			continue
		}
		if visit[src+edges[src][i]] == 0 { // 该目的地已经全部跑过
			continue
		}
		visit[src+edges[src][i]]--
		backTrack(edges, length, pre+edges[src][i], visit, ret, flag)
		visit[src+edges[src][i]]++
	}
}
