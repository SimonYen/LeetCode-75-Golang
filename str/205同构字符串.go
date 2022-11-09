package str

/*
其实如果想到了正确的数据结构来解题，这题目其实简单
可以把题目中的字符串一个看成明文，一个看成密文，也就是说字母的转换需要一一对应
所以用哈希表可以很好解决问题
*/
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	//定义一个哈希表
	hash := make(map[rune]byte)
	//再定义一个哈希表，用于记录已保存的字符（防止字符重复映射）
	exist := make(map[byte]bool)
	//遍历s
	for i, v := range s {
		//首先先查看哈希表中是否有当前的字符
		out, ok := hash[v]
		if ok {
			//有的话，比较out是否与t[i]一致
			if t[i] != out {
				return false
			}
		} else {
			//没有的话就保存下来
			//先看看之前是否保存过了
			if exist[t[i]] {
				return false
			}
			hash[v] = t[i]
			exist[t[i]] = true
		}
	}
	//所以的都符合的话，就是正确的
	return true
}
