package hot100

func isMatch(s string, p string) bool {
	if len(s) == 0 || len(p) == 0 {
		return false
	}
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	// 题目保证每次出现字符 * 时，前面都匹配到有效的字符
	for j := 1; j < len(p)+1; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
