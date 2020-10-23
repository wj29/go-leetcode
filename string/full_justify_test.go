package string

import (
    "testing"
)

// 68. 文本左右对齐
// https://leetcode-cn.com/problems/text-justification/
func Test_FullJustify(t *testing.T) {
    words := []string{"This", "is", "an", "example", "of", "text", "justification."}
    maxWidth := 16

    words = []string{"Science","is","what","we","understand","well","enough","to","explain",
        "to","a","computer.","Art","is","everything","else","we","do"}
    maxWidth = 20
    t.Log(fullJustify(words, maxWidth))
}

func fullJustify(words []string, maxWidth int) []string {
    return nil
}

