package string

import (
	"testing"
)

// 剑指 Offer 20. 表示数值的字符串
// https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
// 65. 有效数字
// https://leetcode-cn.com/problems/valid-number/
func Test_IsNumber(t *testing.T) {
	s := "5e2"
	t.Log(isNumber(s))
}

// 一个合法的数值字符串
// 符号位(+\-)(后面必须是数字或小数点)
// 整数部分(由0-9组成)
// 小数点(前后最少一侧存在数字)
// 小数部分(同整数部分构成)
// 指数部分(e\E,可选符号位,整数部分)

// 一个自动机，总能够回答某种形式的「对于给定的输入字符串 S，判断其是否满足条件 P」的问题
// 确定有限状态自动机（以下简称「自动机」）是一类计算模型。它包含一系列状态，这些状态中：
// 有一个特殊的状态，被称作「初始状态」
// 还有一系列状态被称为「接受状态」，它们组成了一个特殊的集合。其中，一个状态可能既是「初始状态」，也是「接受状态」
// TODO: 先定义状态，再画出状态转移图，最后编写代码
// https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/solution/biao-shi-shu-zhi-de-zi-fu-chuan-by-leetcode-soluti/
type state int
type charType int

const (
	STATE_INITIAL state = iota
	STATE_BIT_SIGN
	STATE_INTEGER
	STATE_POINT
	STATE_POINT_WITHOUT_INTEGER
	STATE_FRACTION
	STATE_EXPONENT
	STATE_EXPONENT_SIGN
	STATE_EXPONENT_INTEGER
	STATE_END
	STATE_ILLEGAL
)

const (
	CHAR_NUMBER charType = iota
	CHAR_SIGN
	CHAR_POINT
	CHAR_EXPONENT
	CHAR_SPACE
	CHAR_ILLEGAL
)

var (
	transferMapping = map[state]map[charType]state{
		STATE_INITIAL: {
			CHAR_SPACE:  STATE_INITIAL,
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_SIGN:   STATE_BIT_SIGN,
			CHAR_POINT:  STATE_POINT_WITHOUT_INTEGER,
		},
		STATE_BIT_SIGN: {
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INTEGER,
		},
		STATE_INTEGER: {
			CHAR_NUMBER:   STATE_INTEGER,
			CHAR_POINT:    STATE_POINT,
			CHAR_EXPONENT: STATE_EXPONENT,
			CHAR_SPACE:    STATE_END,
		},
		STATE_POINT: {
			CHAR_NUMBER:   STATE_FRACTION,
			CHAR_EXPONENT: STATE_EXPONENT,
			CHAR_SPACE:    STATE_END,
		},
		STATE_POINT_WITHOUT_INTEGER: {
			CHAR_NUMBER: STATE_FRACTION,
		},
		STATE_FRACTION: {
			CHAR_NUMBER:   STATE_FRACTION,
			CHAR_EXPONENT: STATE_EXPONENT,
			CHAR_SPACE:    STATE_END,
		},
		STATE_EXPONENT: {
			CHAR_SIGN:   STATE_EXPONENT_SIGN,
			CHAR_NUMBER: STATE_EXPONENT_INTEGER,
		},
		STATE_EXPONENT_SIGN: {
			CHAR_NUMBER: STATE_EXPONENT_INTEGER,
		},
		STATE_EXPONENT_INTEGER: {
			CHAR_NUMBER: STATE_EXPONENT_INTEGER,
			CHAR_SPACE:  STATE_END,
		},
		STATE_END: {
			CHAR_SPACE: STATE_END,
		},
	}

	acceptState = []state{STATE_INTEGER, STATE_POINT, STATE_FRACTION, STATE_EXPONENT_INTEGER, STATE_END}
)

func isNumber(s string) bool {
	state := STATE_INITIAL
	for i := 0; i < len(s); i++ {
		state = transfer(state, toCharType(s[i]))
		if state == STATE_ILLEGAL {
			return false
		}
	}
	for _, v := range acceptState {
		if v == state {
			return true
		}
	}
	return false
}

func toCharType(s byte) charType {
	switch s {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case ' ':
		return CHAR_SPACE
	case '+', '-':
		return CHAR_SIGN
	case 'e', 'E':
		return CHAR_EXPONENT
	case '.':
		return CHAR_POINT
	default:
		return CHAR_ILLEGAL
	}
}

func transfer(state state, s charType) state {
	if v, ok := transferMapping[state][s]; ok {
		return v
	}
	return STATE_ILLEGAL
}
