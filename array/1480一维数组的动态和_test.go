package array

import "testing"

func TestRunningSum(t *testing.T) {
	in := []int{1, 2, 3, 4}
	expected := []int{1, 3, 6, 10}
	actual := runningSum(in)
	if len(expected) != len(actual) {
		t.Errorf("❌长度不一致：期待切片长度：%v\t实际切片长度%v", len(expected), len(actual))
	}
	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("❌元素值不一样：期待值：%v\t实际值：%v", v, actual[i])
			t.Errorf("❌整个切片情况：期待切片：%v\t实际切片：%v", expected, actual)
			break
		}
	}
}
