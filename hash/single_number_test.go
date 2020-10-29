package hash

import (
    "testing"
)

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
// 进阶: 不使用额外空间
// 输入: [4,1,2,1,2]
// 输出: 4
// https://leetcode-cn.com/problems/single-number/  easy
// 137. 只出现一次的数字 II
// https://leetcode-cn.com/problems/single-number-ii/
func Test_SingleNumber(t *testing.T) {
    nums := []int{4, 1, 2, 1, 2}
    t.Log(singleNumber(nums))
}

// 使用hash表即map存储所有元素，出现的次数为1的那个就是单元素(出现两次的元素可以值为2也可以删除该元素)

// 不使用额外空间，即数组本身操作
// 此时我们应该想到位运算(想不到)
// 在计算机中都是二进制运算，位运算为
// & 	与 		都为1才为1
// | 	或 		都为0才为0
// ^ 	异或 	相同为0不同为1
// ～ 	取反 	1变0 0变1
// <<	左移	全部左移，高位去除，低位补0	1<<32 1左移32位 即2^32次方
// >>	右移	全部右移，对无符号数，高位补 0，对于有符号数，高位补符号位
// 异或满足交换律结合律
// 由此我们知道将数组的元素全部取异或运算
func singleNumber(nums []int) int {
    for i := 1; i < len(nums); i++ {
        nums[i] ^= nums[i-1]
    }
    return nums[len(nums)-1]
}

// 位运算的技巧
// a<<1 a>>1 左移N位相当于乘以2的N次方
// 判断奇偶 0 == a & 1 为偶数 a为3与1做与，11&&01=00，为1，3是奇数
// a的相反数为 ~a+1

// 都在说用异或等等的，提供另外一种容易理解的思路，将每个数想象成32位的二进制，对于每一位的二进制的1和0累加起来必然是3N或者3N+1，
// 为3N代表目标值在这一位没贡献，3N+1代表目标值在这一位有贡献(=1)，然后将所有有贡献的位|起来就是结果。
// 这样做的好处是如果题目改成K个一样，只需要把代码改成cnt%k，很通用
// 题解：https://leetcode-cn.com/problems/single-number-ii/comments/
func singleNumber2(nums []int) int {
    var ret uint = 0
    for i := 0; i < 32; i++ {
        cnt := 0
        for j := 0; j < len(nums); j++ {
            if uint(nums[j])&(1<<i) != 0 { // 该位都是1
                cnt++
            }
        }
        if cnt%3 != 0 {
            ret |= 1 << i
        }
    }
    return int(int32(ret)) // 给出的数是int32的
}
