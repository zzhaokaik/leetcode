package leetcode
//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
// pre node1 node2
// pre node2 node1
// pre = node1 下一次循环
func swapPairs(head *ListNode) *ListNode {
	pre:=&ListNode{}
	pre.Next = head
	tmp:=pre
	for tmp.Next!=nil &&tmp.Next.Next!=nil{
		a:=tmp.Next
		b:=tmp.Next.Next
		tmp.Next = b
		a.Next = b.Next
		b.Next = a

		tmp = a
	}
	return pre.Next
}