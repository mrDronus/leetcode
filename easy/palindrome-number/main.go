package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	m := make([]int, 0, 10)
	l := x
	for l >= 10 {
		m = append(m, l%10)
		l = l / 10
	}
	m = append(m, l)
	ml := len(m)
	for i := 0; i < ml/2; i++ {
		if m[i] != m[ml-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	mdl := x
	rev := 0
	for mdl > 0 {
		rev *= 10
		rev += mdl % 10
		mdl /= 10
	}
	return rev == x
}
