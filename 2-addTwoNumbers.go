package main

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// nums := []int{1, 3, 8, 23, 22}
	// create l1, l2
	l1 := createList(4)
	l2 := createList(3)

	showList("l1", l1)
	showList("l2", l2)

	res := addTwoNumbers(l1, l2)
	showList("res", res)
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

func addTwoNumbers(l1, l2 *ListNode) *ListNode {

	head := new(ListNode)
	curr := head
	total := 0

	for l1 != nil || l2 != nil || total > 0 {
		curr.Next = new(ListNode)
		curr = curr.Next

		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}
		curr.Val = total % 10
		total /= 10
	}

	return head.Next
}
