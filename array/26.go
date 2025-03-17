package array

// 遍历一次，双指针 如果相同继续遍历，更改位置
func RemoveDuplicates(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}
