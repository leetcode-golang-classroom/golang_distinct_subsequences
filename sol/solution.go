package sol

func numDistinct(s string, t string) int {
	sLen, tLen := len(s), len(t)
	if sLen < tLen {
		return 0
	}
	// dp[i][j]: number of s[:i] subsequence of t[:j]
	dp := make([][]int, sLen+1)
	for row := range dp {
		dp[row] = make([]int, tLen+1)
		dp[row][0] = 1
	}
	for sEnd := 1; sEnd <= sLen; sEnd++ {
		for tEnd := 1; tEnd <= tLen; tEnd++ {
			if s[sEnd-1] == t[tEnd-1] {
				dp[sEnd][tEnd] = dp[sEnd-1][tEnd-1] + dp[sEnd-1][tEnd]
			} else {
				dp[sEnd][tEnd] = dp[sEnd-1][tEnd]
			}
		}
	}
	return dp[sLen][tLen]
}
