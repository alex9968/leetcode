package main

import (
	"fmt"
)

func main() {

	fmt.Printf("res:%v \n", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) //49
	fmt.Printf("res:%v \n", maxArea([]int{1, 2, 1}))                   //2
}

//盛最多水的容器
//使用左右两个指针，一个在头，一个在尾，每次只移动小的指针，来比较面积的大小，直到相遇
//O(N)
func maxArea(height []int) int {
	res, left, right := 0, 0, len(height)-1
	for left < right {
		v := (right - left) * min(height[left], height[right])
		res = max(res, v)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
