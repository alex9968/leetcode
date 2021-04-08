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
	l1 := createList(8)

	showList("l1", l1)

	res := removeNthFromEnd(l1, 3)
	showList("res", res)
}

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 进阶：你能尝试使用一趟扫描实现吗？

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)

	dummy := &ListNode{0, head}
	curr := dummy

	for i := 0; i < length-n; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return dummy.Next
}

func getLength(head *ListNode) int {
	length := 0
	for ; head != nil; head = head.Next {
		length++
	}
	return length
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
