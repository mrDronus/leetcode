package main

func isValid(s string) bool {
	bracket_stack := []rune{}

	brackets_map := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, r := range s {
		if open_bracket_rune, ok := brackets_map[r]; ok {
			if len(bracket_stack) == 0 || bracket_stack[len(bracket_stack)-1] != open_bracket_rune {
				return false
			}
			bracket_stack = bracket_stack[:len(bracket_stack)-1]
		} else {
			bracket_stack = append(bracket_stack, r)
		}
	}
	return len(bracket_stack) == 0
}
