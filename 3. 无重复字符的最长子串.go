package hot100

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	res := 0
	m := make(map[byte]int)
	left := 0
	for i := 0; i < len(s); i++ {
		if right, ok := m[s[i]]; ok && right >= left {
			// 遇到元素重复，就从窗口左边界的下一位重新开始记录
			left = right + 1
		}
		m[s[i]] = i
		// 更新窗口大小
		if res < i-left+1 {
			res = i - left + 1
		}
	}
	return res
}
