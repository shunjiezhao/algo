package main
func numIslands(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	st := make([][]bool, n)
	for i := range grid {
		st[i] = make([]bool, m)
	}
	type tmp struct{ dx, dy int }
	dir := []tmp{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= n ||
			y < 0 || y >= m {
			return
		}
		if st[x][y] {
			return
		}

		st[x][y] = true
		if grid[x][y] != '1' {
			return
		}
		for _, r := range dir {
			dfs(x+r.dx, y+r.dy)
		}
	}
	var ans int
	for i := range grid {
		for j := range grid[i] {
			if st[i][j] || grid[i][j] != '1' {
				continue
			}
			ans++
			dfs(i, j)
		}
	}
	return ans
}

func solve(grid [][]byte) {
	n, m := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(i, j int) {
		if i < 0 || i >= n ||
			j < 0 || j >= m || grid[i][j] == 'X' || grid[i][j] == '#' {
			return
		}

		grid[i][j] = '#'
		dfs(i-1, j) // 上
		dfs(i+1, j) // 下
		dfs(i, j-1) // 左
		dfs(i, j+1) // 右

	}
	// 4ge
	for i := 0;i < m;i ++{
		// row =  0
		if grid[0][i] == 'O'{
			dfs(0,i)
		}
		// row =  n - 1
		if grid[n - 1][i] == 'O'{
			dfs(0,i)
		}
	}

	for i := 0;i < n;i ++{
		// row =  0
		if grid[i][m - 1] == 'O'{
			dfs(0,i)
		}
		// row =  n - 1
		if grid[i][m - 1] == 'O'{
			dfs(0,i)
		}
	}

	for i := 0 ;i < n;i ++{
		for j:= 0;j < m;j ++{
			if grid[i][j] == 'O'{
				grid[i][j] = 'X'
			}
			if grid[i][j] == '#'{
				grid[i][j] = 'O'
			}
		}
	}
}