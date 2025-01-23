package main

func plusOne(digits []int) []int {
	m := make([]int, len(digits)+1)
	m[0] = 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			m[i+1] = 0
		} else {
			m[i+1] = digits[i] + 1
			return append(digits[:i], m[i+1:]...)
		}
	}
	return m
}

func plusOneV2(digits []int) []int {
	m := make([]int, len(digits)+1)
	increment := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 && increment > 0 {
			m[i+1] = 0
		} else {
			m[i+1] = digits[i] + increment
			increment = 0
		}
	}
	if increment > 0 {
		m[0] = 1
	} else {
		m = m[1:]
	}
	return m
}
