package str

import "testing"

func TestIsSubsequence(t *testing.T) {
	t.Run("正确情况", func(t *testing.T) {
		a1, a2 := "abc", "ahbgdc"
		if !isSubsequence(a1, a2) {
			t.Error("❌")
		}
	})
	t.Run("错误情况1", func(t *testing.T) {
		a1, a2 := "axc", "ahbgdc"
		if isSubsequence(a1, a2) {
			t.Error("❌")
		}
	})
	t.Run("错误情况2", func(t *testing.T) {
		a1, a2 := "abc", "abd"
		if isSubsequence(a1, a2) {
			t.Error("❌")
		}
	})
}
