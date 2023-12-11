// 两数之和 https://leetcode.cn/problems/two-sum/description/

package twosum

import (
	"testing"
)

func TestTowSum(t *testing.T) {
	var tests = []struct {
		input  []int
		target int
	}{
		{[]int{2, 4, 7, 11, 15}, 9},
	}

	for _, test := range tests {
		got := TwoSum(test.input, test.target)
		if test.input[got[0]]+test.input[got[1]] != test.target {
			t.Errorf("TwoSum(%v, %d) = %v", test.input, test.target, got)
		}
	}
}

func TestTwoSum2(t *testing.T) {
	var tests = []struct {
		input  []int
		target int
	}{
		{[]int{2, 4, 7, 11, 15, 19}, 9},
		{[]int{2, 4, 7, 11, 15, 19}, 13},
		{[]int{2, 4, 7, 11, 15, 19}, 15},
	}

	for _, test := range tests {
		got := TwoSum2(test.input, test.target)
		if test.input[got[0]]+test.input[got[1]] != test.target {
			t.Errorf("TwoSum2(%v, %d) = %v", test.input, test.target, got)
		}
	}
}
