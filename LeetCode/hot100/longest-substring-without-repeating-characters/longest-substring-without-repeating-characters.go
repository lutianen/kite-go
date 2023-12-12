// 无重复字符的最长子串 https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/

package longestsubstringwithoutrepeatingcharacters

type Empty struct{}

func LengthOfLongestSubstring(str string) int {
	if len(str) <= 0 {
		return 0
	}

	uset := make(map[byte]Empty)
	left, right, length := 0, 0, 0
	for right < len(str) {
		for {
			if _, ok := uset[str[right]]; ok {
				delete(uset, str[left])
				left++
			} else {
				break
			}
		}

		uset[str[right]] = Empty{}
		length = max(length, right-left+1)
		right++
	}

	return length
}
