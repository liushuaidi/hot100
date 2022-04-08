package hot100

// 中心拓展法
func longestPalindrome(s string) string {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		x, y := extand(s, i, i)
		if y-x > right-left {
			left, right = x, y
		}
		x, y = extand(s, i, i+1)
		if y-x > right-left {
			left, right = x, y
		}
	}
	return s[left : right+1]
}

func extand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
