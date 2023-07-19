package str

func isMatch(s string, p string) bool {
	matches := func(a, b int) bool {
		i, j := s[a], p[b]
		if i == j {
			return true
		}
		return j == '.'
	}
	n, m := len(s), len(p)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	dp[0][0] = 1
	for j := 2; j <= m; j++ {
		if p[j-1] == '*' { // 0 个字符
			dp[0][j] |= dp[0][j-2]
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if matches(i-1, j-1) {
				dp[i][j] |= dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] |= dp[i][j-2] // 0
				if matches(i-1, j-2) {
					dp[i][j] |= dp[i-1][j]
					//			  many
				}
			}

		}
	}
	if dp[n][m] != 0 {
		return true
	}
	return false
}
