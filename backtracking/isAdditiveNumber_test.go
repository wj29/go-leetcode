package backtracking

import (
    "strconv"
    "strings"
    "testing"
)

func Test_IsAdditiveNumber(t *testing.T) {
    num := "112358"
    num = "199100199"
    num = "111"
    num = "1023"
    t.Log(isAdditiveNumber(num))
}

func isAdditiveNumber(num string) bool {
    if len(num) < 3 {
        return false
    }
    flag := false
loop:
    for i := 0; i < len(num)-2; i++ {
        if flag {
            break
        }
        if len(num[:i+1]) != 1 && strings.HasPrefix(num[:i+1], "0") {
            continue
        }
        for j := i + 1; j < len(num)-1; j++ {
            if len(num[i+1:j+1]) != 1 && strings.HasPrefix(num[i+1:j+1], "0") {
                continue loop
            }
            backTraceForIsAdditiveNumber(&flag, j+1, toInt(num[:i+1]), toInt(num[i+1:j+1]), num)
        }
    }
    return flag
}

func backTraceForIsAdditiveNumber(flag *bool, start, first, second int, num string) {
    if start == len(num) {
        *flag = true
    }
    if *flag {
        return
    }
    for i := start; i < len(num); i++ {
        if len(num[start:i+1]) != 1 && strings.HasPrefix(num[start:i+1], "0") {
            return
        }
        if first+second != toInt(num[start:i+1]) {
            continue
        }
        backTraceForIsAdditiveNumber(flag, i+1, second, first+second, num)
    }
    return
}

func toInt(s string) int {
    r, _ := strconv.Atoi(s)
    return r
}
