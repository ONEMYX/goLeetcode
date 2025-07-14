package str

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	list := make([][]int, 0)
	sort.Ints(nums)
	// i 固定 改变j  让 i +j+ k ==0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			continue
		}
		// 增加限制
		// 1： 当i 跟j 确定时 ，则只有一个数能让他们为0
		// 2： 当 j 增大时， 则之后上次截止 位置的右侧能满足
		// 3: 当 i+j+k 小于0 J++  直接从 小于0 的位置开始计算
		// 4: 当 i+j+k 大于0 时 k-- 从当前位置开始计算
		// 5: 当 i+j+k ==0 时 J++ 从开始位置计算

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				list = append(list, []int{nums[i], nums[l], nums[r]})
				for l++; l < r+1 && nums[l] == nums[l-1]; l++ {
				}
				for r--; r > l && nums[r] == nums[r+1]; r-- {
				}
			}
			if sum > 0 {
				r--
			}
			if sum < 0 {
				l++
			}
		}
	}
	return list
}
