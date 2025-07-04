package array

func findMiddleIndex(nums []int) int {
	ret := -1
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			left[i] = 0
		} else {
			left[i] += left[i-1] + nums[i-1]
		}
	}
	def := left[len(nums)-1] + nums[len(nums)-1]
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			right[i] = def - nums[i]
		} else {
			right[i] += right[i-1] - nums[i]
		}
		//right[i] = def - right[i]
	}
	left[0] = 0
	right[len(nums)-1] = 0
	for i := 0; i < len(nums); i++ {
		if left[i] == right[i] {
			return i
		}
	}

	return ret
}
