package hot100

import "sort"
// 调用库函数合并后快排，求中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	n := len(nums1)
	if n%2 == 0 {
		return float64((nums1[n/2-1] + nums1[n/2]) / 2)
	}
	return float64(nums1[n/2])
}

// 手动比较进行合并（类似于合并两个有序链表）
func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	var res []int
	m, n := len(nums1), len(nums2)
	l1, l2 := 0, 0
	for l1 < m && l2 < n{
		if nums1[l1] < nums2[l2]{
			res = append(res, nums1[l1])
			l1++
		}else{
			res = append(res, nums2[l2])
			l2++
		}
	}
	res = append(res, nums1[l1:]...)
	res = append(res, nums2[l2:]...)
	length := m + n
	if length % 2 == 1{
		return float64(res[length/2])
	}
	mid1 := res[length/2]
	mid2 := res[length/2-1]
	return float64(mid1 + mid2)/2.0
}

