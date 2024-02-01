// 最长回文字串 https://leetcode.cn/problems/longest-palindromic-substring/

package longestpalindromicsubstring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		wants []string
	}{
		{"case 1: ", "babad", []string{"bab", "aba"}},
		{"case 2: ", "cbbd", []string{"bb"}},
		{"case 3: ", "a", []string{"a"}},
		{"case 4: ", "aa", []string{"aa"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.s)
			ok := false
			for i := 0; i < len(tt.wants); i++ {
				if got == tt.wants[i] {
					ok = true
				}
			}
			if !ok {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.wants)
			}
		})
	}
}
