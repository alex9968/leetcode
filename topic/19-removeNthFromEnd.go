package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func main() {
	l1 := createList(8)

	showList("l1", l1)

	res := removeNthFromEnd(l1, 3)
	showList("res", res)
}

//给你一个链表，删除链表的 倒数 第 n 个结点，并且返回链表的头结点。

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getListLen(head)

	dummy := &ListNode{0, head}
	curr := dummy

	for i := 0; i < length-n; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return dummy.Next
}
