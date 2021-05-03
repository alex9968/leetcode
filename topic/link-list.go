package main

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func createList(length int) *ListNode {
	head := new(ListNode)
	curr := head

	for i := 1; ; i++ {
		//new node
		curr.Next = new(ListNode)
		curr = curr.Next

		//set new node
		// rand.Seed(time.Now().Unix())
		val := rand.Intn(10)
		if val > 0 {
			curr.Val = val
		} else {
			curr.Val = 1
		}

		if i == length {
			break
		}
	}
	return head.Next
}
func showList(listName string, l *ListNode) {
	curr := l
	fmt.Printf("show %s:", listName)
	for curr != nil {
		fmt.Printf("%d \t", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func getListLen(head *ListNode) int {
	length := 0
	for ; head != nil; head = head.Next {
		length++
	}
	return length
}