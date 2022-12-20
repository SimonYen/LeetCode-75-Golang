package list

func reverseList(head *ListNode) *ListNode {
	//实在是理解不了，算了死记硬背
	prev := new(ListNode)
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
