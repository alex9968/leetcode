package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4} //：[[-1,-1,2],[-1,0,1]]

	res := threeSum(nums)

	fmt.Printf("res:%v", res)

}

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		//枚举b
		for second := first + 1; second < n; second++ {
			//需要和上次枚举的数不一样
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			//遍历c
			for second < third && nums[second]+nums[third] > target {
				third--
			}

			if second == third {
				break //枚举结束
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}

		}

	}
	return ans
}
