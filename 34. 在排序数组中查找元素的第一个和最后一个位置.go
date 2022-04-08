package hot100

// 暴力
func searchRange01(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	// left, right := 0, len(nums)-1
	ans := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			ans = append(ans, i)
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target {
			ans = append(ans, i)
			break
		}
	}
	if len(ans) == 2 {
		return ans
	}
	return []int{-1, -1}
}

// 