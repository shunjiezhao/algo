package main

import (
	"sort"
)

func letterCombinations(digits string) []string {
	dig2Char := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var dfs func(u int)
	s := make([]byte, len(digits))
	var ans []string
	dfs = func(u int) {
		if u == len(digits) {
			if u != 0 {
				ans = append(ans, string(s))
			}
			return
		}

		for _, v := range dig2Char[digits[u]] {
			s[u] = byte(v)
			dfs(u + 1)
		}
	}
	dfs(0)
	return ans
}

// combine 组合： https://leetcode.cn/problems/combinations/?envType=study-plan-v2&envId=top-interview-150
func combine(n int, k int) [][]int {
	var dfs func(u, idx int)
	var ans [][]int
	path := make([]int, k)
	dfs = func(u, idx int) {
		if idx == k {
			var tmp []int
			tmp = append(tmp, path...)
			ans = append(ans, tmp)
			return
		}
		if u > n {
			return
		}

		// chose
		path[idx] = u
		dfs(u+1, idx+1)
		// don't chose
		dfs(u+1, idx)
	}
	dfs(1, 0)
	return ans
}

func permute(nums []int) [][]int {
	var dfs func(u int)
	var ans [][]int
	n := len(nums)
	path := make([]int, n)
	choose := make([]bool, n)
	dfs = func(u int) {
		if u == n {
			var tmp []int
			tmp = append(tmp, path...)
			ans = append(ans, tmp)
			return
		}

		for i, v := range choose {
			if v {
				continue
			}
			choose[i] = true
			path[u] = nums[i]
			dfs(u + 1)
			choose[i] = false
		}
	}
	dfs(0)
	return ans
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var dfs func(u, sum int)
	var ans [][]int
	n := len(candidates)
	path := make([]int, 0, 40)
	dfs = func(u, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			var tmp []int
			tmp = append(tmp, path...)
			ans = append(ans, tmp)
			return
		}
		if u >= n {
			return
		}
		l := len(path)
		num := candidates[u]
		for i := 0; i <= target; i += num {
			if sum+i > target {
				break
			}
			if i != 0 {
				path = append(path, num)
			}
			dfs(u+1, sum+i)
		}
		path = path[:l]
	}
	dfs(0, 0)
	return ans
}

func totalNQueens(n int) int {
	col, r, l := make([]bool, n), make([]bool, 2*n), make([]bool, 2*n)

	var dfs func(row int)

	var ans int
	dfs = func(row int) {
		if row == n {
			ans++
			return
		}

		for i := 0; i < n; i++ {
			if col[i] ||
				r[row+i] ||
				l[row-i+n] {
				continue
			}
			r[row+i], l[row-i+n], col[i] = true, true, true
			dfs(row + 1)
			r[row+i], l[row-i+n], col[i] = false, false, false
		}
	}
	dfs(0)
	return ans
}

func generateParenthesis(n int) []string {
	end := n * 2

	var dfs func(u, sum int)
	path := make([]byte, end)
	var ans []string

	// ( +1 ) -1
	dfs = func(u, sum int) {
		if sum < 0 { // count( ')' ) > count( '(' )
			return
		}
		if u == end && sum == 0 {
			ans = append(ans, string(path))
			return
		}
		if u == end {
			return
		}

		// ')'
		path[u] = ')'
		dfs(u+1, sum-1)
		path[u] = '('
		dfs(u+1, sum+1)
	}
	dfs(0, 0)
	return ans
}

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	var dfs func(row, col, idx int) bool
	type pair struct{ dx, dy int }
	dirt := []pair{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	st := make([][]bool, n)
	for i := range board {
		st[i] = make([]bool, m)
	}

	dfs = func(row, col, idx int) bool {
		if idx == len(word) {
			return true
		}
		if row < 0 || row >= n || col < 0 || col >= m || idx > len(word) {
			return false
		}
		if st[row][col] {
			return false
		}
		st[row][col] = true
		defer func () {
			st[row][col] = false
		}()
		if word[idx] != board[row][col] {
			return false
		}

		for _, v := range dirt {
			if dfs(row+v.dx, col+v.dy, idx+1) {
				return true
			}
		}
		return false
	}

	for i := range board {
		for j := range board[0] {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
