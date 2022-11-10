package list

/*
其实这题做法很简单，就是双指针，互相比较谁最小，小的插入到新链表尾部
但是链表题目是会有一些很细节的操作
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//定义返回的链表
	dummy := new(ListNode)
	last := dummy
	cur1, cur2 := list1, list2
	//开始遍历
	for cur1 != nil && cur2 != nil {
		//比较两指针指向的值
		if cur1.Val < cur2.Val {
			last.Next = cur1
			cur1 = cur1.Next
		} else {
			last.Next = cur2
			cur2 = cur2.Next
		}
		last = last.Next
	}
	//结束后，看看还有没有剩余
	if cur1 != nil {
		last.Next = cur1
	}
	if cur2 != nil {
		last.Next = cur2
	}
	return dummy.Next
}
