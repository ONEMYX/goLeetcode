package array

import (
	"fmt"
)

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	// 1. 从num1 进行遍历，找到最小的k个乘积
	// 2. 推理得出 先找出最小的乘积，一定是 nums1 中的一个数与 nums2 中的一个数相乘
	// 3. nums1 中进行循环，
	// 4. 如果num1 为负数， 则 nums2 最右侧开始 计算值 然后确定下一个最小值
	// 4.1. 然后 计算 nums1[i+1]* nums2[j] < nums[i]*nums2[j-1] 如果成立，则 nums1[i+1] 与 nums2[j] 是最小的乘积移动i
	// 4.2 如果 nums1[i+1]* nums2[j] >= nums[i]*nums2[j-1] 则 nums1[i] 与 nums2[j-1] 是最小的乘积移动j
	// 5. 如果num1 为整数，则从nums2 最左侧开始
	// 5.1. 然后 计算 nums1[i+1]* nums2[j] < nums[i]*nums2[j+1] 如果成立，则 nums1[i+1] 与 nums2[j] 是最小的乘积移动i
	// 5.2 如果 nums1[i+1]* nums2[j] >= nums[i]*nums2[j+1] 则 nums1[i] 与 nums2[j+1] 是最小的乘积移动j
	// 6. 重复 3-5 步骤，直到找到第 k 个最小的乘积
	i := 0
	ans := int64(0)
	for k > 0 && i < len(nums1)-1 {
		if nums1[i] < 0 {
			for j := len(nums2) - 1; j >= 0; j-- {
				ans = int64(nums2[j]) * int64(nums1[i])
				fmt.Printf("i=%v, j=%v, ans=%v\n", nums1[i], nums2[j], ans)
				// 4.1. 然后 计算 nums1[i+1]* nums2[j] < nums[i]*nums2[j-1] 如果成立，则 nums1[i+1] 与 nums2[j] 是最小的乘积移动i 注意边界问题
				// 4.1.1 如果 nums1[i+1] 超出了 nums1 的范围，则 nums1[i] 与 nums2[j-1] 是最小的乘积移动j
				// 4.1.2 如果 nums2[j-1] 超出了 nums2 的范围，则 nums1[i+1] 与 nums2[j] 是最小的乘积移动i
				if j-1 > 0 && i+1 < len(nums1) {
					if nums1[i+1]*nums2[len(nums2)-1] < nums1[i]*nums2[j-1] {
						break
					}
				}
			}
		} else {
			// 5.1. 然后 计算 nums1[i+1]* nums2[j] < nums[i]*nums2[j+1] 如果成立，则 nums1[i+1] 与 nums2[j] 是最小的乘积移动i
			for j := 0; j < len(nums2); j++ {
				ans = int64(nums1[i]) * int64(nums2[j])
				fmt.Printf("i=%v, j=%v, ans=%v\n", nums1[i], nums2[j], ans)
				if j+1 < len(nums2) && i+1 < len(nums1) {
					if nums1[i+1]*nums2[0] < nums1[i]*nums2[j+1] {
						break
					}
				}
			}
		}
		k--
		i++
	}
	return ans
}
