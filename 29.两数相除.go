/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	res := 0
	// for dividend <= divisor {
	// 	dividend -= divisor
	// 	res++
	// }
	for i := 31; i >= 0; i-- {
		if (dividend>>i)-divisor >= 0 {
			dividend = dividend - (divisor << i)
			res += 1 << i
		}
	}
	return res * sign
}

// @lc code=end

