package main

import (
	"fmt"
)

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

func main() {

	fmt.Println("bbbbb:", lengthOfLongestSubstring("bbbbb"))       //1
	fmt.Println("abcabcbb:", lengthOfLongestSubstring("abcabcbb")) //3
	fmt.Println("pwwkew:", lengthOfLongestSubstring("pwwkew"))     //3
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	res, index := 0, -1
	m := map[byte]int{}
	n := len(s)

	for i := 0; i < n; i++ {
		//移动左边
		if i != 0 {
			delete(m, s[i-1])
		}

		//移动右边
		for index+1 < n && m[s[index+1]] == 0 { //保证下一个存在且不在map中
			m[s[index+1]]++
			index++
		}
		res = max(res, index-i+1)
	}
	return res
}
func lengthOfLongestSubstring2(s string) int {
	m := map[byte]int{}
	n := len(s)
	res, index := 0, -1

	for i := 0; i < n; i++ {
		//左边界移动
		if i != 0 {
			delete(m, s[i-1])
		}

		//右边界移动
		for index+1 < n && m[s[index+1]] == 0 {
			m[s[index+1]]++
			index++
		}
		res = max(res, index-i+1)
	}
	return res
}
