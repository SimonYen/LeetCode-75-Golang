package str

/*
之前一直没搞懂的，现在终于搞懂了，
就是首先先统计各字符出现的个数，
偶数字符直接放两边，奇数字符拿一个放中间，然后再放两边，
但是奇数字符只能用一次。
*/

func longestPalindrome(s string) int {
	//先定义一个哈希表，统计字符个数
	hash := make(map[rune]int)
	for _, c := range s {
		hash[c]++
	}
	ans := 0
	for _, v := range hash {
		ans += v / 2 * 2
		if v%2 == 1 && ans%2 == 0 {
			ans++
		}
	}
	return ans
}
