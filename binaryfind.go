package main

import "fmt"

func main() {
	s := []int{11, 22, 33, 44, 55, 66}
	binaryfind(&s, 44, 0, 5)
}

func binaryfind(s *[]int, target, startIndex, endIndex int) {
	if startIndex > endIndex {
		fmt.Print("Not found")
		return
	}

	midIndex := (startIndex + endIndex) / 2

	if (*s)[midIndex] > target {
		binaryfind(s, target, startIndex, midIndex-1)
	} else if (*s)[midIndex] < target {
		binaryfind(s, target, midIndex+1, endIndex)
	} else {
		fmt.Println("Target index:%d", midIndex)
	}
}
