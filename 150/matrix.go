package main

import (
	"fmt"
	"sort"
	"strconv"
)

// isValidSudoku: https://leetcode.cn/problems/valid-sudoku/submissions/438938668/
func isValidSudoku(board [][]byte) bool {
	n, m := len(board), len(board[0])
	row, col, cet := make([]int, n), make([]int, m), make([]int, 9)

	get := func(a, b int) int {
		return (a/3)*3 + (b / 3)
	}

	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		b := board[x][y]

		tbit := 1 << (b - '0')
		if row[x]&tbit != 0 ||
			col[y]&tbit != 0 ||
			cet[get(x, y)]&tbit != 0 {
			return false
		}
		row[x] |= tbit
		col[y] |= tbit
		cet[get(x, y)] |= tbit
		return true
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] != '.' && !dfs(i, j) {
				return false
			}
		}
	}
	return true
}

// solveSudoku https://leetcode.cn/problems/sudoku-solver/submissions/438938609/
func solveSudoku(board [][]byte) {
	n, m := len(board), len(board[0])
	row, col, cet := make([]int, n), make([]int, m), make([]int, 9)

	point := make([][2]int, 0)
	get := func(a, b int) int {
		return (a/3)*3 + (b / 3)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] != '.' {
				row[i] |= 1 << (board[i][j] - '0')
				col[j] |= 1 << (board[i][j] - '0')
				cet[get(i, j)] |= 1 << (board[i][j] - '0')
			} else {
				point = append(point, [2]int{i, j})
			}
		}
	}

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i == len(point) {
			return true
		}
		x, y := point[i][0], point[i][1]
		for j := 1; j <= 9; j++ {
			tbit := 1 << j
			if row[x]&tbit != 0 ||
				col[y]&tbit != 0 ||
				cet[get(x, y)]&tbit != 0 {
				continue
			}
			row[x] |= tbit
			col[y] |= tbit
			cet[get(x, y)] |= tbit
			board[x][y] = byte(j + '0')

			if dfs(i + 1) {
				return true
			}

			row[x] ^= tbit
			col[y] ^= tbit
			cet[get(x, y)] ^= tbit
			board[x][y] = byte('.')
		}
		return false
	}
	dfs(0)
}

func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	ans := make([]int, 0)
	var inc, desc *int
	i, j := 0, 0
	update := func() {
		if inc != nil {
			*inc++
		}
		if desc != nil {
			*desc--
		}
	}

	// [)
	// r t -> i == rowBegin,j == colEnd, inc = &i, j --, rowBegin ++
	// r d -> i == rowEnd, j == colEnd, desc = &j, i --, colEnd--
	// l d -> i == rowEnd, j < colBegin,desc = &i,j++, rowEnd--
	// l t ->  i < rowBegin, j == colBegin, i--, inc = &j, colBegin++
	rowBegin, rowEnd := 0, n
	colBegin, colEnd := 0, m
	inc = &j
	ans = append(ans, matrix[0][0])
	for len(ans) != n*m {
		switch {
		case i == rowBegin && j == colEnd:
			inc, desc = &i, nil
			j--
			rowBegin++
		case i == rowEnd && j == colEnd-1:
			inc, desc = nil, &j
			i--
			colEnd--
		case i == rowEnd-1 && j < colBegin:
			inc, desc = nil, &i
			j++
			rowEnd--
		case i < rowBegin && j == colBegin:
			inc, desc = &j, nil
			i++
			colBegin++
		}
		update()
		if i >= rowBegin && j >= colBegin && i < rowEnd && j < colEnd {
			ans = append(ans, matrix[i][j])
		}
	}
	return ans
}

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			t := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = t
		}
	}
}

func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	row, col := make([]bool, n), make([]bool, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		if row[i] {
			for j := 0; j < m; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < m; j++ {
		if col[j] {
			for i := 0; i < n; i++ {
				matrix[i][j] = 0
			}
		}
	}

}

func gameOfLife(board [][]int) {
	// -1 0 -> 1
	// 2  1 -> 0
	n, m := len(board), len(board[0])
	cnt := func(x, y int) (live int) {
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				if i >= 0 && i < n && j >= 0 && j < m && i != x && j != y {
					fmt.Println(x, y, i, j)
					if board[i][j] > 0 {
						live++
					}
				}
			}
		}
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			live := cnt(i, j)
			fmt.Println(i, j, live)
			if live < 2 || live > 3 {
				// dead
				if board[i][j] == 1 {
					board[i][j] = 2
				}
			}
			if live == 3 {
				if board[i][j] == 0 {
					board[i][j] = -1
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == -1 {
				board[i][j] = 1
			}
			if board[i][j] == 2 {
				board[i][j] = 0
			}
		}
	}
}
