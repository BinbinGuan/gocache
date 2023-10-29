package leetcode

import "math"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}

func mySqrt1(x int) int {
	l, r := x, 0
	ans := -1

	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			l = mid + 1
			ans = mid
		} else {
			r = mid - 1
		}
	}

	return ans
}
