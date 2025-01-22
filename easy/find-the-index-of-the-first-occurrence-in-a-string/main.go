package main

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	needleComparisonIndex := 0
	needleIndexLen := len(needle) - 1
	for i := 0; i < len(haystack); i++ {
		if needle[needleComparisonIndex] == haystack[i] {
			if needleComparisonIndex == needleIndexLen {
				return i - needleIndexLen
			}
			needleComparisonIndex++
		} else {
			if needleComparisonIndex > 0 {
				// rollback to next char from comparison started
				i -= needleComparisonIndex
				needleComparisonIndex = 0
			}
		}
	}
	return -1
}
