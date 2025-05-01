package slidingwindow

func numSubArrayProductionLessThank(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	windows := 1
	res := 0

	left := 0
	for right, x := range nums {
		windows *= x
		for windows >= k {
			windows /= nums[left]
			left += 1
		}

		res += right - left + 1
	}

	return res
}
