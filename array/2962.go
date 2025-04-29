package array

import "fmt"

func countSubarrays(nums []int, k int) int64 {
	mx := nums[0]
	for _, v := range nums {
		if v > mx {
			mx = v
		}
	}
	cur := 0
	left := 0
	ans := int64(0)
	for _, v := range nums {
		if v == mx {
			cur++
		}
		for cur == k {
			if nums[left] == mx {
				cur--
			}
			left++
		}
		ans += int64(left)
	}
	fmt.Println(ans)
	return ans

}
