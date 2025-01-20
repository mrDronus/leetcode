package main

// O(N^2) + less memory, - slower
func twoSum(nums []int, target int) []int {
	num_len := len(nums)
	// test for input len
	if num_len < 2 {
		return []int{}
	}

	for i := 0; i <= num_len-2; i++ {
		for j := i + 1; j <= num_len-1; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// O(N) - more memory, + faster
func twoSum2(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	looked_up := make(map[int]int, len(nums))

	for i, right := range nums {
		// determining number for solution
		left := target - right
		if j, ok := looked_up[left]; ok {
			return []int{j, i}
		}
		looked_up[right] = i
	}
	return []int{}
}
