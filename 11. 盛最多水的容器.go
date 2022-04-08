package hot100

// !!! 双指针
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res, area := 0, 0
	for left < right {
		if height[left] <= height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}
		if res < area {
			res = area
		}
	}
	return res
}
