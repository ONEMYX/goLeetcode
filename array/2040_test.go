package array

import "testing"

func TestKthSmallestProduct(t *testing.T) {
	//输入：nums1 = [2,5], nums2 = [3,4], k = 2
	//输出：8
	nums1 := []int{2, 5}
	nums2 := []int{3, 4}
	k := int64(2)
	ans := kthSmallestProduct(nums1, nums2, k)
	if ans != 8 {
		t.Errorf("kthSmallestProduct(%v, %v, %v) = %v; want 8", nums1, nums2, k, ans)
	}
	//输入：nums1 = [-4,-2,0,3], nums2 = [2,4], k = 6
	//输出：0
	nums1 = []int{-4, -2, 0, 3}
	nums2 = []int{2, 4}
	k = int64(6)
	ans = kthSmallestProduct(nums1, nums2, k)
	if ans != 0 {
		t.Errorf("kthSmallestProduct(%v, %v, %v) = %v; want 0", nums1, nums2, k, ans)
	}
	//输入：nums1 = [-2,-1,0,1,2], nums2 = [-3,-1,2,4,5], k = 3
	//输出：-6
	nums1 = []int{-2, -1, 0, 1, 2}
	nums2 = []int{-3, -1, 2, 4, 5}
	k = int64(3)
	ans = kthSmallestProduct(nums1, nums2, k)
	if ans != -6 {
		t.Errorf("kthSmallestProduct(%v, %v, %v) = %v; want -6", nums1, nums2, k, ans)
	}
}
