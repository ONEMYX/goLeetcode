package array

import "sort"

// 输入：citations = [3,0,6,1,5]
// 输入：citations = [1,3,1]
// [1，1，3]
// [0,1,3,5,6]
// 可以理解为 至少有i 篇文章 引用大于等于i
func hIndex(citations []int) int {
	h := 0
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return h
}
