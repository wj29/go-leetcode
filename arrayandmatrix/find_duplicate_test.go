package arrayandmatrix

import "testing"

// 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数
// 假设只有一个重复的整数，找出这个重复的数
// https://leetcode-cn.com/problems/find-the-duplicate-number
func Test_FindDuplicate(t *testing.T) {

}

// 排序，i与i+1的值相同的时候即为重复值
// map实现

// 二分法 在1-n+1的数组中存在1-n个数，最少有一个重复的数
// 原数组是乱序的，二分1-n，即mid=(n+1)/2，遍历原数组，等于mid的个数大于1个返回mid，大于mid的个数超过一半那么重复的数据在mid-n中，
// 反之则在1-mid中
func findDuplicate1(nums []int) int {
	return 0
}

// 快慢指针
// 慢指针一次走一步，快指针一次走两步，然后记录快慢指针相遇的点
// 之后再用两个指针，一个指针从起点出发，一个指针从相遇点出发，当他们再次相遇的时候就是入口点了
// 上面的话成立的数学原理:
// 假设链表有环，设入环前长度为a，入环到相交的长度为b，环的长度为c，快指针和慢指针第一次相交时比慢指针多跑了N圈
// 此处注意，再第一次相遇时，慢指针肯定没有跑满一圈，因为快指针先入环且速度是慢指针两倍，慢指针刚入环时快指针已经走了一段距离
// 那么得到慢指针走过的距离是a+b，快指针走过的距离是a+N*c+b
// 快指针的速度是慢指针的2倍，那么 2(a+b) = a+N*c+b
// 如果快指针这是也在入口处，那么快慢指针已经相遇，圈的长度c=a，b=0,上面简化成 2a=a+c
// 公式转换成 a+b == N*c，即 N*c = a+b
// 重新使慢指针从开始，快指针从相交点开始，速度相同，上述的公式标示经过N圈之后，慢指针和快指针会在相遇出再次相遇
// 由于快慢指针速度相同，那么从入环处到相交点走过的是相同的路径，那么第一次相遇一定在入环处
func findDuplicate2(nums []int) int {
	if len(nums) <= 1 {
		panic("invalid nums length")
	}

	slow := nums[0]
	fast := nums[nums[0]]
	for nums[slow] != nums[fast] { // 找到相遇的点
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0 // 慢指针从头出发，快指针从相交点出发，再次相遇时则是入口点，此时未走到开始值就判断，头部本身也是路程的一部分
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}