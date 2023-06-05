package main

func removeTrailingZeros(num string) string {
	n := len(num)
	for i := n - 1; i > 0; i-- {
		if num[i] != '0' {
			break
		}
		num = num[:i]
	}
	return num
}

func differenceOfDistinctValues(grid [][]int) [][]int {
	n, m := len(grid), len(grid[0]) // row, col
	ans := make([][]int, n)
	contain := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < m
	}
	calc := func(i, j int) {
		if !contain(i, j) {
			return
		}
		help := func(delta int) int {
			cnt := make(map[int]bool)
			a := 0
			for x, y := i+delta, j+delta; contain(x, y); x, y = x+delta, y+delta {
				if _, ok := cnt[grid[x][y]]; ok {
					continue
				}
				a++
				cnt[grid[x][y]] = true
			}
			return a
		}
		ans[i][j] = abs(help(1) - help(-1))
	}
	for i := 0; i < n; i++ {
		ans[i] = make([]int, m)
		for j := 0; j < m; j++ {
			calc(i, j)
		}
	}
	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a

}

func minimumCost(s string) int64 {
	n := len(s)
	ans := int64(0)
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			ans += int64(min(i, n-i))
		}
	}
	return ans
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	println(minimumCost("010101"))
}
