package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("res1:", threeSumClosest([]int{-1, 2, 1, -4}, 1)) //2, 与 target 最接近的和是 2 (-1 + 2 + 1 = 2)
}

//最接近的三数之和
// 时间复杂度：O(N^2)O(N 2)，其中 NN 是数组 \textit{nums}nums 的长度。我们首先需要 O(N \log N)O(NlogN) 的时间对数组进行排序，随后在枚举的过程中，使用一重循环 O(N)O(N) 枚举 aa，双指针 O(N)O(N) 枚举 bb 和 cc，故一共是 O(N^2)O(N 2)。
// 空间复杂度：O(\log N)O(logN)。排序需要使用 O(\log N)O(logN) 的空间。然而我们修改了输入的数组 \textit{nums}nums，在实际情况下不一定允许，因此也可以看成使用了一个额外的数组存储了 \textit{nums}nums 的副本并进行排序，此时空间复杂度为 O(N)O(N)。

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)

	// 根据差值的绝对值来更新答案
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	// 枚举 a
	for i := 0; i < n; i++ {
		// 保证和上一次枚举的元素不相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 使用双指针枚举 b 和 c
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// 如果和为 target 直接返回答案
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				// 如果和大于 target，移动 c 对应的指针
				k0 := k - 1
				// 移动到下一个不相等的元素
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				// 如果和小于 target，移动 b 对应的指针
				j0 := j + 1
				// 移动到下一个不相等的元素
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
