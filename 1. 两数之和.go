package hot100

// 法1：暴力

// 法2：哈希表
func twoSum(nums []int, target int) []int {
	n := len(nums)
	if n < 2 {
		return nil
	}
	res := []int{}
	hashtable := make(map[int]int)
	for i := 0; i < n; i++ {
		tmp := target - nums[i]
		if _, ok := hashtable[tmp]; ok {
			res = append(res, hashtable[tmp], i)
			return res
		}
		hashtable[nums[i]] = i
	}
	return res
}
