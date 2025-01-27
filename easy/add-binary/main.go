package main

func addBinary(a string, b string) string {
	lenA := len(a)
	lenB := len(b)
	maxLen := lenA
	if lenA < lenB {
		maxLen = lenB
	}

	r := make([]byte, maxLen+1)
	c := 0
	for i := 0; i < maxLen; i++ {
		ai := 0
		bi := 0
		if i < lenA && a[lenA-i-1] == '1' {
			ai = 1
		}
		if i < lenB && b[lenB-i-1] == '1' {
			bi = 1
		}
		s := ai + bi + c
		c = s / 2
		r[maxLen-i] = '0' + byte(s%2)
	}
	if c > 0 {
		r[0] = '1'
		return string(r)
	} else {
		return string(r[1:])
	}
}
