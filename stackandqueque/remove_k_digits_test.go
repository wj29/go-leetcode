package stackandqueque

import "testing"

// 402. 移掉K位数字
// https://leetcode-cn.com/problems/remove-k-digits/
func Test_RemoveKdigits(t *testing.T)  {
    num := "1432219"
    k := 3
    t.Log(removeKdigits(num, k))
}

// 单调栈
// 从左至右扫描，当前的数还不确定要不要删，先入栈暂存，保留对它的记忆。
// 123531这样「高位递增」的数，会尽量删低位，不会删递增的高位。
// 432135这样「高位递减」的数，肯定想干掉高位，让高位变小，尽量变成「高位递增」
// 所以，如果当前数比栈顶更大，就还是递增态，是满意的，入栈。
// 如果当前数比栈顶更小，要当机立断删栈顶的数，不管后面有没有更大的，为什么？
// 因为栈顶的数在高位，删掉它，后面小的顶上，高位变小，优于低位变小。
// https://leetcode-cn.com/problems/remove-k-digits/solution/wei-tu-jie-dan-diao-zhan-dai-ma-jing-jian-402-yi-d/
func removeKdigits(num string, k int) string {
    if len(num) == k {
        return "0"
    }
    stack := make([]rune, 0)
    // 遍历num字符串
    for _, c := range num {
        // 只要k>0且当前的c比栈顶的小，则栈顶出栈，k--
        for k > 0 && len(stack) > 0 && stack[len(stack)-1] > c {
            stack = stack[:len(stack)-1]
            k--
        }
        // 当前的字符不是"0"，或栈非空（避免0入空栈），入栈
        if c != '0' || len(stack) != 0 {
            stack = append(stack, c)
        }
    }
    // 如果还没删够，要从stack继续删，直到k=0
    for k > 0 {
        stack = stack[:len(stack)-1]
        k--
    }
    // 如果栈空了，返回"0"，如果栈非空，转成字符串返回
    if len(stack) == 0 {
        return "0"
    }
    return string(stack)
}