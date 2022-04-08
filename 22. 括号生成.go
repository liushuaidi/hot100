package hot100

// 回溯算法
func generateParenthesis(n int) []string {
	ans := []string{}
	backtracking("", &ans, 0, 0, n)
	return ans
}

func backtracking(path string, ans *[]string, left, right, n int) {
	if len(path) == n*2 {
		*ans = append(*ans, path)
		return
	}
	if left < n {
		path += "("
		backtracking(path, ans, left+1, right, n)
		path = path[:len(path)-1]
	}
	if right < left {
		path += "("
		backtracking(path, ans, left, right+1, n)
		path = path[:len(path)-1]
	}
}
