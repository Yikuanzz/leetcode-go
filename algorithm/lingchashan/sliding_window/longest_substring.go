package slidingwindow

func lengthOfLongestSubstring(s string) int {
	windows := make([]bool, 128)
	res := 0

	left := 0
	for right, x := range s {
		for windows[x] {
			windows[s[left]] = false
			left++
		}

		windows[x] = true
		res = max(res, right-left+1)
	}
	return res
}
