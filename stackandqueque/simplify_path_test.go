package stackandqueque

import (
	"strings"
	"testing"
)

// 71. 简化路径
// https://leetcode-cn.com/problems/simplify-path/
func Test_SimplifyPath(t *testing.T) {
	path := "/a/./b/../../c/"
	t.Log(simplifyPath(path))
}

// 通过/划分成不同的路径
// 空格(两个//)和.代表当前路径，不存入结果
// ..返回上一目录(上一个目录存在，即非根目录)
// 其余情况则是路径
func simplifyPath(path string) string {
	ret := make([]string, 0)
	paths := strings.Split(path, "/")
	for _, v := range paths {
		if v == "" || v == "." {
			continue
		} else if len(ret) > 0 && v == ".." {
			ret = ret[:len(ret)-1]
		} else if v != ".." {
			ret = append(ret, v)
		}
	}
	return "/" + strings.Join(ret, "/")
}
