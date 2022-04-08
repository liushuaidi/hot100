package hot100

/*
1、从右向左开始找，找到第一个小于右边的值
比如 12385764，此时5就是找到的那个值；
2、从右向左开始找，找到第一个大于上一步发现的值，发现6满足；
3、两者对调位置 变成：12386754
4、把第一步中5的下标位置之后的所有元素反转，12386457；
*/

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// fmt.Println(nums[i])
	// 如果不是最后一个排列
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		// fmt.Println(nums[j])
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
