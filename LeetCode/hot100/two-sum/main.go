// 两数之和 https://leetcode.cn/problems/two-sum/description/

package twosum

// 暴力：O(n^2)
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Hashtable: O(n)
func TwoSum2(nums []int, target int) []int {
	mapNumIdx := make(map[int]int)
	for idx, num := range nums {
		mapNumIdx[num] = idx
	}

	for i := 0; i < len(nums); i++ {
		left := target - nums[i]
		if idx, ok := mapNumIdx[left]; ok && i != idx {
			return []int{i, idx}
		}
	}

	return nil
}
