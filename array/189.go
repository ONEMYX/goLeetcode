package array

func rotate(nums []int, k int) {
	a1 := make([]int, len(nums))
	l := len(nums)
	for i, n := range nums {
		a1[(i+k)%l] = n
	}
	copy(nums, a1)
}
