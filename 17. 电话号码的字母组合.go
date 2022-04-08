package hot100

// 回溯
func letterCombinations(digits string) []string {
	if len(digits) == 0 || len(digits) > 4 {
		return nil
	}
	letters := [10]string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	ans := []string{}
	backtracking("", digits, letters, 0, &ans)
	return ans
}
func backtracking(path, dights string, letters [10]string, idx int, ans *[]string) {
	if len(path) == len(dights) {
		*ans = append(*ans, path)
		return
	}
	words := letters[dights[idx]-'0']
	for i := 0; i < len(words); i++ {
		path += string(words[i])
		backtracking(path, dights, letters, idx+1, ans)
		path = path[:len(path)-1]
	}
}
