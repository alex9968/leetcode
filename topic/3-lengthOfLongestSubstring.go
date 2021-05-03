package main

import (
	"fmt"
)

//字母都不同的最长子串长度

func main() {

	fmt.Println("bbbbb:", lengthOfLongestSubstring("bbbbb"))       //1 b
	fmt.Println("abcabcbb:", lengthOfLongestSubstring("abcabcbb")) //3  adb
	fmt.Println("pwwkew:", lengthOfLongestSubstring("pwwkew"))     //3
}

func lengthOfLongestSubstring(s string) int {
	res, index := 0, -1//一开始左右边界都在-1处
	m := map[byte]int{}
	n := len(s)

	for i := 0; i < n; i++ {
		//移动左边界
		if i != 0 {
			delete(m, s[i-1])
		}

		//尽肯能移动右边界
		for index+1 < n && m[s[index+1]] == 0 { //保证下一个存在且不在map中
			m[s[index+1]]++ //标记在map
			index++ 
		}
		res = max(res, index-i+1) //判断当前
	}
	return res
}
