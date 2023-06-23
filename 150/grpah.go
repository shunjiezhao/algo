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
	for i := 0; i < m; i++ {
		// row =  0
		if grid[0][i] == 'O' {
			dfs(0, i)
		}
		// row =  n - 1
		if grid[n-1][i] == 'O' {
			dfs(0, i)
		}
	}

	for i := 0; i < n; i++ {
		// row =  0
		if grid[i][m-1] == 'O' {
			dfs(0, i)
		}
		// row =  n - 1
		if grid[i][m-1] == 'O' {
			dfs(0, i)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 'O' {
				grid[i][j] = 'X'
			}
			if grid[i][j] == '#' {
				grid[i][j] = 'O'
			}
		}
	}
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	clone := make(map[*Node]*Node)
	var dfs func(node *Node) *Node

	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if t, ok := clone[node]; ok {
			return t
		}
		nc := &Node{Val: node.Val}
		clone[node] = nc

		for _, v := range node.Neighbors {
			nc.Neighbors = append(nc.Neighbors, dfs(v))
		}
		return nc
	}
	return dfs(node)
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	type Node struct {
		cur       string
		Val       []float64
		Neighbors []*Node
	}
	cnt := make(map[string]*Node)
	contain := func(a string) bool {
		_, ok := cnt[a]
		return ok
	}
	get := func(a string) *Node {
		if t, ok := cnt[a]; ok {
			return t
		}
		cnt[a] = &Node{cur: a}
		return cnt[a]
	}
	for i, v := range equations {
		a, b := v[0], v[1]
		an, bn := get(a), get(b)
		an.Neighbors = append(an.Neighbors, bn)
		an.Val = append(an.Val, values[i])
		bn.Neighbors = append(bn.Neighbors, an)
		bn.Val = append(bn.Val, 1/values[i])
	}
	var ans []float64
	var dfs func(cur, end string, num float64) float64
	var visit map[string]bool
	dfs = func(cur, end string, num float64) float64 {
		if !contain(cur) {
			return -1
		}

		if cur == end {
			return num
		}
		if _, ok := visit[cur]; ok {
			return -1
		}
		visit[cur] = true
		n := cnt[cur]
		for i, v := range n.Neighbors {
			if t := dfs(v.cur, end, num*n.Val[i]); t != -1 {
				return t
			}
		}
		return -1
	}

	for _, v := range queries {
		if !contain(v[0]) || !contain(v[1]) {
			ans = append(ans, -1)
			continue
		}
		visit = make(map[string]bool)
		ans = append(ans, dfs(v[0], v[1], 1))
	}
	return ans
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	type node struct{ val int }
	in := make(map[int]int)

	edges := make([][]int, numCourses)
	for _, v := range prerequisites {
		edges[v[0]] = append(edges[v[0]], v[1])
		in[v[1]]++
	}

	q := make([]int, 0)
	for k := 0; k < numCourses; k++ {
		if in[k] == 0 {
			q = append(q, k)
		}
	}

	for len(q) != 0 {
		t := q[0]
		q = q[1:]
		for _, v := range edges[t] {
			in[v]--
			if in[v] == 0 {
				q = append(q, v)
				delete(in, v)
			}
		}
	}
	return len(in) == 0
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	type node struct{ val int }
	in := make(map[int]int)

	edges := make([][]int, numCourses)
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
		in[v[0]]++
	}

	q := make([]int, numCourses)
	tt := 0
	for k := 0; k < numCourses; k++ {
		if in[k] == 0 {
			q[tt] = k
			tt++
		}
	}

	hh := 0
	for hh < tt {
		t := q[hh]
		hh++
		for _, v := range edges[t] {
			in[v]--
			if in[v] == 0 {
				q[tt] = v
				tt++
				delete(in, v)
			}
		}
	}
	if len(in) == 0 {
		return q
	}
	return nil
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}
	visit := make(map[string]bool)
	vd := make(map[string]bool)
	for _, v := range wordList {
		vd[v] = true
	}
	if !vd[endWord] {
		return 0
	}

	type pair struct {
		str  string
		step int
	}
	q := make([]pair, 1)
	q[0] = pair{beginWord, 0}
	visit[beginWord] = true
	for len(q) != 0 {
		t := q[0]
		q = q[1:]
		bts := []byte(t.str)
		for i := range bts {
			for j := byte(0); j < 26; j++ {
				pre := bts[i]
				bts[i] = (bts[i]-'a'+j)%26 + 'a'
				str := string(bts)
				bts[i] = pre
				if visit[str] || !vd[str] {
					continue
				}
				visit[str] = true
				q = append(q, pair{str, t.step + 1})
				bts[i] = pre
				if str == endWord {
					return t.step + 1
				}
			}
		}
	}
	return 0
}

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	visit := make(map[string]bool)
	vd := make(map[string]bool)
	for _, v := range bank {
		vd[v] = true
	}
	if !vd[endGene] {
		return 0
	}

	type pair struct {
		str  string
		step int
	}
	q := make([]pair, 1)
	q[0] = pair{startGene, 1}
	visit[startGene] = true
	change := []byte{'A', 'C', 'G', 'T'}
	for len(q) != 0 {
		t := q[0]
		q = q[1:]
		bts := []byte(t.str)
		for i := range bts {
			for _, j := range change {
				pre := bts[i]
				bts[i] = j
				str := string(bts)
				bts[i] = pre
				if visit[str] || !vd[str] {
					continue
				}

				visit[str] = true
				q = append(q, pair{str, t.step + 1})
				bts[i] = pre
				if str == endGene {
					return t.step
				}
			}
		}
	}
	return 0
}

func nearestExit(maze [][]byte, entrance []int) int {
	type pair struct{ x, y, step int }
	n, m := len(maze), len(maze[0])
	st := make([][]bool, n)
	for i := range st {
		st[i] = make([]bool, m)
	}

	q := make([]pair, 1)
	q[0] = pair{entrance[0], entrance[1], 0}
	dirt := []pair{
		{0, 1, 0},
		{0, -1, 0},
		{1, 0, 0},
		{-1, 0, 0},
	}
	st[q[0].x][q[0].y] = true
	for len(q) != 0 {
		t := q[0]
		q = q[1:]
		for _, v := range dirt {
			x, y := v.x+t.x, v.y+t.y
			if !(x >= 0 && x < n && y >= 0 && y < m) {
				continue
			}
			if st[x][y] || maze[x][y] == '+' {
				continue
			}
			st[x][y] = true
			q = append(q, pair{x, y, t.step + 1})
			if x == 0 || x == n-1 || y == 0 || y == m-1 {
				return t.step + 1
			}
		}
	}
	return -1
}

func orangesRotting(maze [][]int) int {
	type pair struct{ x, y, step int }
	n, m := len(maze), len(maze[0])

	q := make([]pair, 0)
	for i := range maze {
		for j := range maze[0] {
			if maze[i][j] == 2 {
				q = append(q, pair{i, j, 0})
			}
		}
	}
	dirt := []pair{
		{0, 1, 0},
		{0, -1, 0},
		{1, 0, 0},
		{-1, 0, 0},
	}
	mx := 0
	for len(q) != 0 {
		t := q[0]
		q = q[1:]
		if mx < t.step {
			mx = t.step
		}
		for _, v := range dirt {
			x, y := v.x+t.x, v.y+t.y
			if !(x >= 0 && x < n && y >= 0 && y < m) {
				continue
			}
			if maze[x][y] != 1{
				continue
			}
			maze[x][y] = 2

			q = append(q, pair{x, y, t.step + 1})
		}
	}

	for i := range maze {
		for j := range maze[0] {
			if maze[i][j] == 1 {
				return -1
			}
		}
	}
	return mx
}