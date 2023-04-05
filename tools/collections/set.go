package collections

import "sort"

// 取两个整型数组交集
func Intersect(nums1, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			nums1[k] = nums1[i]
			k++
			i++
			j++
		}
	}
	return nums1[:k]
}

// nums1對nums2求差集
func Difference(nums1, nums2 []int) []int {
	var difference []int
	nums2Map := make(map[int]bool)
	for _, v := range nums2 {
		nums2Map[v] = true
	}

	for _, v := range nums1 {
		if !nums2Map[v] {
			difference = append(difference, v)
		}
	}
	return difference
}
