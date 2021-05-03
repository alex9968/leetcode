package main

import (
	"fmt"
)

//给你一个字符串 s，找到 s 中最长的回文子串。

func main() {

	fmt.Println("babad:", longestPalindrome("babad")) //1
	fmt.Println("ac:", longestPalindrome("ac"))       //3
	fmt.Println("a:", longestPalindrome("a"))         //3
}

// 1. 先规划之后才可以选择

func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}
