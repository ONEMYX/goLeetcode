package array

// 1. 双指针
// 2. 增加一个只获取两个相同的函数
func removeDuplicates(nums []int) int {
	j, num := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			num = 0
			nums[j] = nums[i]
		} else {
			num++
			if num == 1 {
				j++
				nums[j] = nums[i]
			}
		}
	}
	return j + 1
}
