package leetcode

// 反转链表

type ListNode struct {
	Val  int
	Next *ListNode
}


// 背


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
