package twopointers

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	n := len(nums)

	res := [][]int{}

	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if nums[i]+nums[n-1]+nums[n-2] < 0 {
			continue
		}

		l, r := i+1, n-1
		for l < r {
			total := nums[l] + nums[r] + x
			if total < 0 {
				l++
			} else if total > 0 {
				r--
			} else {
				res = append(res, []int{
					x, nums[l], nums[r],
				})

				l += 1
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r -= 1
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}

	}
	return res
}
