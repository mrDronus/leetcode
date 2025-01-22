package main

func removeElement(nums []int, val int) int {
	currentIndex := 0
	for _, v := range nums {
		if v == val {
			continue
		}
		nums[currentIndex] = v
		currentIndex++
	}
	return currentIndex
}
