/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start

// func binarySearch(nums []int, target int, right bool) int {
// 	left, right := 0, len(nums)-1
// 	for left <= right {
// 		mid := (right-left)/2 + left
// 		if nums[mid] == target {
// 			return mid
// 		}
// 		if target < nums[mid] {
// 			right = mid - 1
// 		} else {
// 			left = mid + 1
// 		}
// 	}
// 	return -1
// }

func searchRange(nums []int, target int) []int {
	leftMost := sort.SearchInts(nums, target)
	if leftMost == len(nums) || nums[leftMost] != target {
		return []int{-1, -1}
	}
	rightMost := sort.SearchInts(nums, target+1) - 1
	return []int{leftMost, rightMost}
}

// @lc code=end

