package dp

/*
动态规划入门题，首先要明白dp数组表示什么，
初值，状态转移方程。
这题设dp[i]为第i天能取得的最大利润
可知dp[1]=0，当天买当天卖
状态转移方程dp[i]=max(dp[i-1],prices[i]-min)，要么是前一天的最大值，要么是今天的差值。总之看谁利润最大
*/

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	//定义dp数组
	dp := make([]int, n)
	//定义最低价格
	min := prices[0]
	for i := 1; i < n; i++ {
		if min > prices[i] {
			//更新最小值
			min = prices[i]
		}
		dp[i] = max(dp[i-1], prices[i]-min)
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
