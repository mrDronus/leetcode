package main

func romanToInt(s string) int {
	defs := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	result := 0
	for i, v := range s {
		result += defs[v]

		if i > 0 && defs[rune(s[i-1])] < defs[v] {
			result -= defs[rune(s[i-1])] * 2
		}
	}
	return result
}
