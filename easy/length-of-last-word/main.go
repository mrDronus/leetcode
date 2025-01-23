package main

func lengthOfLastWord(s string) int {
	wordLen := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if wordLen > 0 {
				break
			}
			continue
		}
		wordLen++
	}
	return wordLen
}
