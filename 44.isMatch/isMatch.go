package frytea_leetcode

import (
	"fmt"
)

// 失败，难以穷举
func isMatch_classic(s string, p string) bool {
	i, j := 0, 0
	if len(s) == 0 {
		return true
	}
	for i < len(s) && j < len(p) {
		fmt.Println(i, s[i], j, p[j])
		if s[i] == p[j] {
			i++
			j++
		} else if p[j] == '*' {
			// 匹配0次
			if j+2 < len(p) && s[i] == p[j+1] {
				i++
				j += 2
			} else if j+1 < len(p) && i+1 < len(s) && s[i+1] == p[j+1] {
				i++
				j++
			} else {
				i++
			}
		} else if p[j] == '?' {
			i++
			j++
		} else {
			return false
		}
	}
	fmt.Println(i, j)
	if i == len(s) && j == len(p) || (j == len(p)-1 && p[j] == '*') {
		return true
	}
	return false
}

// Time  O(mn)
// Space O(mn)
func isMatch_dp(s string, p string) bool {
	m, n := len(s), len(p)
    // 初始化 dp
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
    // 前 * 个 dp[o0][j] 为 true，只有 * 能匹配空串
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = true
		} else {
			break
		}
	}
    // 空匹配空为 true
	dp[0][0] = true
    // 状态转移
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
                dp[i][j] = false
            }
		}
	}
	return dp[m][n]
}

func isMatch(s string, p string) bool {
	return isMatch_dp(s, p)
}
