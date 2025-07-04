package array

// 循环数组
func RemoveElement(nums []int, val int) int {
	j := 0
	for _, v := range nums {
		if v != val {
			nums[j] = v
			j++
		}
	}
	//fmt.Println(nums, nums[:j])
	return j
}
