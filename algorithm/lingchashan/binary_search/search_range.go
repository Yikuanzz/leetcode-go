package binarysearch

func searchRange(nums []int, target int) []int {
	first := lower_bound_1(nums, target)
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}
	end := lower_bound_1(nums, target+1) - 1
	return []int{first, end}
}
