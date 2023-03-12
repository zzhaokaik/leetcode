package leetcode

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}


// k 个一组反转链表


// 1 2 3 4 5


// 1 分组
// 2 组内进行反转
// 3 组间连接


func reverseKGroup(head *ListNode, k int) *ListNode {
	project :=&ListNode{Val: 0,Next: head}
	pre:=project


	for head !=nil{

		end:=getNext(head,k)
		if end ==nil{
			break
		}


		nextHead:=end.Next
		reverseInGroup(head,nextHead)

		pre.Next = end
		head.Next = nextHead

		pre = head
		head = nextHead




	}

	return pre.Next

}

func reverseInGroup(head,nextHead *ListNode){

	project:=&ListNode{
		Val: 0,
		Next: head,
	}
	 pre :=project

	for head !=nextHead{
		next:=head.Next
		head.Next = pre
		pre = head
		head = next


	}

}

func getNext(head *ListNode,k int )  *ListNode{
	for head !=nil{
		if k>0{
			k--
			head = head.Next
		}
		return head
	}
	return nil


}