package array

func countPairs(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (i*j)%k == 0 && nums[i] == nums[j] {
				res++
			}
		}
	}
	return res
}
