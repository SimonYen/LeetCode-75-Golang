package str

import "testing"

func TestIsIsomorphic(t *testing.T) {
	t.Run("测试1正确情况", func(t *testing.T) {
		a1, a2 := "egg", "add"
		out := isIsomorphic(a1, a2)
		if !out {
			t.Errorf("❌：函数输入参数为：%s\t%s", a1, a2)
		}
	})
	t.Run("测试2错误情况", func(t *testing.T) {
		a1, a2 := "foo", "bar"
		out := isIsomorphic(a1, a2)
		if out {
			t.Errorf("❌：函数输入参数为：%s\t%s", a1, a2)
		}
	})
	t.Run("测试3错误情况", func(t *testing.T) {
		a1, a2 := "badc", "baba"
		out := isIsomorphic(a1, a2)
		if out {
			t.Errorf("❌：函数输入参数为：%s\t%s", a1, a2)
		}
	})
}
