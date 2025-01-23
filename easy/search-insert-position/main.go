package main

// O(log n ) + recursion
func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	if target == nums[len(nums)-1] {
		return len(nums) - 1
	}

	left := nums[:len(nums)/2]
	right := nums[len(nums)/2:]
	if target > left[len(left)-1] {
		return len(left) + searchInsert(right, target)
	} else {
		return searchInsert(left, target)
	}
}
