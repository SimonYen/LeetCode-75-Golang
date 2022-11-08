package array

/*
第一次看这个题目的时候完全没有啥思路
后来看官方答案才理解了，要采取一些数学方法
定义整个数组和为total，当前下标指向的值为val，下标前面的数组和（前缀和）为sum
那么可知后面的数组和为total-val-sum
题目求的是sum=total-val-sum
可得val=total-2*sum
*/
func pivotIndex(nums []int) int {
	//首先先求出总共的值total
	total := 0
	for _, v := range nums {
		total += v
	}
	//定义前缀和sum
	sum := 0
	for i, v := range nums {
		//如果满足当前值为total-2*sum，说明找到了
		if v == total-2*sum {
			return i
		}
		sum += v //更新前缀和
	}
	return -1 //没找到
}
