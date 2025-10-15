package array

func maxIncreasingSubarrays(nums []int) int {
	// 滑动窗口
	cnt, precnt, ans := 1, 1, 1
	//for i, num := range nums {
	//	cnt++
	//	if i == len(nums)-1 || num >= nums[i+1] {
	//		ans = max(ans, min(precnt, cnt))
	//		ans = max(ans, cnt/2)
	//		precnt = cnt
	//		cnt = 0
	//	}
	//}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cnt++
		} else {
			precnt = cnt
			cnt = 1
		}
		ans = max(ans, min(precnt, cnt))
		ans = max(ans, cnt/2)
	}

	return ans
}
