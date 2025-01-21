package main

func longestCommonPrefix(strs []string) string {
	rbucket := make([]byte, 0, 200)
	current_index := 0
	for {
		all_matches := true
		for _, s := range strs {
			if len(s) <= current_index || s[current_index] != strs[0][current_index] {
				all_matches = false
				break
			}
		}
		if !all_matches {
			break
		}
		rbucket = append(rbucket, strs[0][current_index])
		current_index++
	}
	return string(rbucket)
}
