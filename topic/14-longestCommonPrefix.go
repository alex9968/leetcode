package main

import (
	"fmt"
)

func main() {
	fmt.Println("res1:", longestCommonPrefix([]string{"flower", "flow", "flight"})) //fl
	fmt.Println("res2:", longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""                                         // ""
}

//最长公共前缀
// 时间复杂度：O(mn)，其中 m 是字符串数组中的字符串的平均长度，n是字符串的数量。最坏情况下，字符串数组中的每个字符串的每个字符都会被比较一次。
// 空间复杂度：O(1)。使用的额外空间复杂度为常数。

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0] //len逐渐变小
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(s1, s2 string) string { //返回两个单词相同的前缀
	n := min(len(s1), len(s2))
	index := 0
	for index < n && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}
