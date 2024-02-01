// 最长回文字串 https://leetcode.cn/problems/longest-palindromic-substring/

package longestpalindromicsubstring

// 动态规划
// dp[i][j] - 表示 s[i]...[j] 是否为回文子串，dp[0][0] = true
// dp[i][j] = dp[i+1][j-1]^s[i]==s[j]
// 只有当s[i]==s[j]且dp[i+1][j-1]为真时，dp[i][j]才为真
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	stIdx, length := 0, 1

	for r := 1; r < len(s); r++ {
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l+1 > length {
					length = r - l + 1
					stIdx = l
				}
			}
		}
	}

	return s[stIdx : stIdx+length]
}
