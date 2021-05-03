package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 3, 8, 23, 22}

	fmt.Println("1:", twoSum(nums, 11))
	fmt.Println("2:", twoSum2(nums, 11))

}

// 2for  O(n2)
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); i++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}

	}
	return nil
}

//hash O(n)
func twoSum2(nums []int, target int) []int {
	hash := map[int]int{}

	for i, v := range nums {
		hash[v] = i
	}

	for i, v := range nums {
		pair := target - v
		if j, ok := hash[pair]; ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}
