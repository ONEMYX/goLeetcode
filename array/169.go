package array

func majorityElement(nums []int) int {
	ans := nums[0]
	count := 0
	for _, v := range nums {
		if count == 0 {
			ans = v
			count++
		} else {
			if v == ans {
				count++
			} else {
				count--
			}
		}
	}
	return ans
}
