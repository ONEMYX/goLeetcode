package array

import (
	"testing"
)

func TestFindSmallestInteger(t *testing.T) {
	// 测试用例1: 基本用例
	nums1 := []int{1, -10, 7, 13, 6, 8}
	value1 := 5
	expected1 := 4
	result1 := findSmallestInteger(nums1, value1)
	if result1 != expected1 {
		t.Errorf("测试用例1失败: 期望 %d, 得到 %d", expected1, result1)
	}
}
