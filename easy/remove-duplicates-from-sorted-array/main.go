package main

func removeDuplicates(nums []int) int {
	lastLessIndex := 0
	lastLessValue := nums[0]
	for _, v := range nums {
		if v == lastLessValue {
			continue
		}
		lastLessIndex++
		lastLessValue = v
		nums[lastLessIndex] = v
	}
	return lastLessIndex + 1
}
