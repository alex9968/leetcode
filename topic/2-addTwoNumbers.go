package main

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

//两个数做加法
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
