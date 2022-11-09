package str

/*
这题如果要解决后续挑战需要使用dp方法，好家伙真是万物都是dp
我使用最简单的双指针方法
*/

func isSubsequence(s string, t string) bool {
	//定义指针，长度
	i, j, n, m := 0, 0, len(s), len(t)
	if n == 0 {
		return true
	}
	if m == 0 {
		return false
	}
	for i < n && j < m {
		if s[i] == t[j] {
			//如果两个指针指向的内容是一致的，那么就同时前进
			i++
			j++
		} else {
			//否则就在t的下一个字符进行匹配
			j++ //这里其实可以把j++提出来
		}
	}
	if i < n {
		//i没有走到末尾，证明还有一些字符没有匹配
		return false
	}
	//如果走到了末尾，还需要比较一下
	if s[i-1] != t[j-1] {
		return false //对应 abc abd的这种情况
	}
	return true
}
