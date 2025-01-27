package main

// find result with binary search
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	max := x / 2
	min := 1
	v := x / 2
	for {
		vl := v * v
		vr := (v + 1) * (v + 1)
		if vl <= x && vr > x {
			break
		}
		if vl > x {
			max = v
		} else {
			min = v
		}
		v = min + (max-min)/2
	}
	return v
}
