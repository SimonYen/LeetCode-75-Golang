package tree

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		currentLen := len(queue)
		tmp := []int{}
		for i := 0; i < currentLen; i++ {
			currentNode := queue[0]
			queue = queue[1:]
			tmp = append(tmp, currentNode.Val)
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}
