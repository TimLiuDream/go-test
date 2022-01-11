package main

// 单链表的反转

type ListNode struct {
	val  int
	Next *ListNode
}

// 1->2->3->4->nil
// 2->3->4->nil->1
func main() {
	head := new(ListNode)
	head.val = 1

	first := new(ListNode)
	first.val = 2

	head.Next = first
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
