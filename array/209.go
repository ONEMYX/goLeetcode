package array

import "math"

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MinInt
	left, right := 0, 0
	sum := 0
	for right < n {
		sum += nums[right]
		for sum >= target {
			if sum == target {
				ans = min(ans, right-left+1)
			}
			sum -= nums[left]
			left++
		}
		if sum < target {
			right++
		}
	}
	return ans
}
