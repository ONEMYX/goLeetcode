package array

func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	num := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, nums[i]+i)
		if end == i {
			end = maxPosition
			num++
		}
	}
	return num
}
