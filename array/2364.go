package array

// nums[i]-i != num[j]-j
func countBadPairs(nums []int) int64 {
	num := int64(0)
	mp := make(map[int]int)
	for k, v := range nums {
		key := v - k
		num += int64(k - mp[key])
		mp[key]++
	}
	return num
}
