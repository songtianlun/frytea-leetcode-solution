/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	slow := 0
	for fast := 0; fast < length; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// @lc code=end

