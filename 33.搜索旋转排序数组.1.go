/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	// ans := -1
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

// @lc code=end

