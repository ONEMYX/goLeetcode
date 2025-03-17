package array

func twoSum(nums []int, target int) []int {
	a := map[int]int{}

	for i, x := range nums {
		if p, ok := a[target-x]; ok {
			return []int{p, i}
		}
		a[x] = i
	}
	return nil
}
