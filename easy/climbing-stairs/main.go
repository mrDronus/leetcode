package main

// fibonacci!
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	prev := 1
	v := 1
	for i := 2; i <= n; i++ {
		tmp := v
		v += prev
		prev = tmp
	}
	return v
}
