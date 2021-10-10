/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	c, x0 := float64(x), float64(x)
	for {
		x1 := 0.5 * (x0 + c/x0)
		if math.Abs(x0-x1) < 1e-7 {
			break
		}
		x0 = x1
	}
	return int(x0)
}

// @lc code=end

