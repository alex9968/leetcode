package main

import (
	"fmt"
)

func main() {
	fmt.Println("res1:", Test([]string{"flower", "flow", "flight"})) //fl
	fmt.Println("res2:", Test([]string{"dog", "racecar", "car"}))    // ""                                         // ""
}

func Test(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	n := len(strs)

	for i := 0; i < n; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}

	return prefix
}

func lcp(s1, s2 string) string {
	n := min(len(s1), len(s2))
	index := 0
	for index < n && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}
