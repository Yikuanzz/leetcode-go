package slidingwindow

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	res := n + 1
	l := 0
	windows := 0

	for r, v := range nums {
		windows += v
		for windows-nums[l] >= target {
			windows -= nums[l]
			l += 1
		}

		if windows >= target {
			res = min(res, r-l+1)
		}
	}

	if res < n {
		return res
	}
	return 0
}
