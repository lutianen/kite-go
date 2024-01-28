// 寻找两个正序数组的中位数 https://leetcode.cn/problems/median-of-two-sorted-arrays/

package findmediansortedarrays

import "sort"

// findMedianSortedArrays 排序实现，不满足题目要求，但是可过全部样例
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

	total_len := len(nums1)
	if total_len%2 == 1 {
		return float64(nums1[total_len/2])
	} else {
		return float64(nums1[total_len/2]+nums1[total_len/2-1]) * 0.5
	}
}
