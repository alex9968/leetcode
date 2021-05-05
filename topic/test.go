package main

import "sort"

func main() {

}

func Test(nums []int, target int) (res [][]int) {
	sort.Ints(nums)
	n := len(nums)

	for i:=0;  i< n-3 && nums[i] + nums[i+1] + nums[i+2] + nums[i+3] <= target; i++ {
		if i> 0 && nums[i] == nums[i-1] || nums[i] + nums[n-3]+ nums[n-2]+ nums[n-1] < target {
			continue
		}

		for j := i+1; j< n-2 && nums[i] + nums[j]

	} 

}
