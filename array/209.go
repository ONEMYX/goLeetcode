package array

import "sort"

func minSubArrayLen(target int, nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	num := 0
	for i := 0; i < n; i++ {
		sum := nums[i]
		if sum == target {
			num = 1
			break
		}
		for j := i + 1; j < n; j++ {
			sum += nums[j]
			if sum < target {
				continue
			}
			if sum > target {
				break
			}
			if sum == target {
				if num == 0 {
					num = j - i + 1
				} else {
					num = min(num, j-i+1)
				}
			}
		}
	}

	return num
}
