package list

/*
这道题也可用快慢指针来搞，但是推理太复杂了，
最简单的是哈希表
*/

func detectCycle(head *ListNode) *ListNode {
	//用一个哈希表来记录便利的结点是否之前出现过
	hash := make(map[*ListNode]int)
	for head != nil {
		//检测是否出现过
		_, ok := hash[head]
		if ok {
			return head
		}
		//没有出现，插入
		hash[head] = 1
		head = head.Next
	}
	return nil
}
