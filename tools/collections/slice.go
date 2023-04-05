package collections

// []int64 to []int
func SliceInt64ToInt(nums []int64) []int {
	s := make([]int, len(nums))
	for i, n := range nums {
		s[i] = int(n)
	}
	return s
}

// []int to []int64
func SliceIntToInt64(nums []int) []int64 {
	s := make([]int64, len(nums))
	for i, n := range nums {
		s[i] = int64(n)
	}
	return s
}
