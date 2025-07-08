package array

// [3,2,1,0,4]
// [3,0,8,2,0,0,1]
func canJump(nums []int) bool {
	mx := len(nums) - 1
	maNums := nums[0]
	for i := 0; i < mx; i++ {
		maNums = maNums - 1
		if i+nums[i] >= mx {
			return true
		}
		maNums = max(maNums, nums[i])
		if maNums == 0 {
			return false
		}
	}
	return false
}
