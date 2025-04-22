package twopointers

func trap_1(height []int) int {
	n := len(height)

	left_max := make([]int, n)
	right_max := make([]int, n)

	left_max[0] = height[0]
	for i := 1; i < n; i++ {
		left_max[i] = max(height[i], left_max[i-1])
	}

	right_max[n-1] = height[n-1]
	for j := n - 2; j > -1; j-- {
		right_max[j] = max(height[j], right_max[j+1])
	}

	res := 0
	for k, h := range height {
		res += min(left_max[k], right_max[k]) - h
	}
	return res
}

func trap_2(height []int) int {
	n := len(height)
	res := 0

	left_max, right_max := 0, 0
	i, j := 0, n-1

	for i <= j {
		left_max = max(left_max, height[i])
		right_max = max(right_max, height[j])

		if left_max < right_max {
			res += left_max - height[i]
			i += 1
		} else {
			res += right_max - height[j]
			j -= 1
		}
	}
	return res
}
