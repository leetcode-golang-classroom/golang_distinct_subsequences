package sol

func numDistinctDFS(s string, t string) int {
	sLen := len(s)
	tLen := len(t)
	if sLen < tLen {
		return 0
	}
	cache := make([][]int, sLen)
	for row := range cache {
		cache[row] = make([]int, tLen)
		for col := range cache[row] {
			cache[row][col] = -1
		}
	}
	var dfs func(sIndex, tIndex int) int
	dfs = func(sIndex, tIndex int) int {
		if tIndex == tLen { // match to the end
			return 1
		}
		if sIndex == sLen { // not match
			return 0
		}
		if cache[sIndex][tIndex] != -1 {
			return cache[sIndex][tIndex]
		}
		var result int
		if s[sIndex] == t[tIndex] {
			// choose sIndex + not choose
			result = dfs(sIndex+1, tIndex+1) + dfs(sIndex+1, tIndex)
		} else {
			result = dfs(sIndex+1, tIndex)
		}
		cache[sIndex][tIndex] = result
		return result
	}
	return dfs(0, 0)
}
