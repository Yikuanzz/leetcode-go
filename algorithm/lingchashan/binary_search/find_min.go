package binarysearch

func findMin(nums []int) int {
	left, right := -1, len(nums)-1
	last := len(nums) - 1

	for left+1 < right {
		mid := left + ((right - left) >> 1)
		if nums[mid] <= nums[last] {
			right = mid
		} else {
			left = mid
		}
	}
	return nums[right]
}
