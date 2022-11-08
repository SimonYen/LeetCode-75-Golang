package array

import "testing"

func TestPivotIndex(t *testing.T) {
	in := []int{1, 7, 3, 6, 5, 6}
	expected := 3
	actual := pivotIndex(in)
	if expected != actual {
		t.Errorf("❌：输入数组：%v\n期待值：%d\t实际值：%d", in, expected, actual)
	}
}
