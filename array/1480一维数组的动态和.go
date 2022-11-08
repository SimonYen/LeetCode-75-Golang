package array

/*
这题很简单就是累加就可以了
注意边界值
*/
func runningSum(nums []int) []int {
	//定义数组长度n
	n := len(nums)
	if n == 1 {
		return nums
	}
	//定义返回的切片
	res := make([]int, n)
	res[0] = nums[0]
	for i := 1; i < n-1; i++ {
		res[i] = res[i-1] + nums[i] //当前的值为前一个累加值加上现在数组里的值
	}
	//最后一个元素别忘记更新了
	res[n-1] = res[n-2] + nums[n-1]
	return res //有点dp里的状态方程的味道了
}
