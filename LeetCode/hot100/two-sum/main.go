// 两数之和 https://leetcode.cn/problems/two-sum/description/

package main

import "fmt"

// 暴力：O(n^2)
func TowSum(nums []int, target int) []int {
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
func TowSum2(nums []int, target int) []int {
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

func main() {
	nums := []int{2, 4, 7, 11, 15}
	target := 9

	ret := TowSum(nums, target)

	for _, dat := range ret {
		fmt.Printf("%d ", dat)
	}
	fmt.Println()

	ret = TowSum2(nums, target)
	for _, dat := range ret {
		fmt.Printf("%d ", dat)
	}
	fmt.Println()
}
