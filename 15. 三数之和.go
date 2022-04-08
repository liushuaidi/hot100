package hot100

import "sort"

// 排序+双指针
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if 0-nums[i] == nums[j]+nums[k] {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if 0-nums[i] < nums[j]+nums[k] {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
