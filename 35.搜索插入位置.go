/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := len(nums)
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] >= target {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}
	return ans
}

// @lc code=end

