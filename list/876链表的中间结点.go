package list

/*
非常经典的快慢指针题目，快指针一次走两步，慢指针一次走一步，
当快指针走到尽头时，慢指针就走到中间了，
具体数学原理很简单，因为快指针的路程是整个链表，那么速度只有它一半的慢指针就是中间点
*/

func middleNode(head *ListNode) *ListNode {
	//定义快慢指针
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
