package stack

import (
	"math"
)

// https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-ii/
// 2909. 元素和最小的山形三元组 II

func minimumSum(nums []int) int {
	preMin := make([]int, len(nums)+1)
	sufMin := make([]int, len(nums)+1)
	preMin[0] = math.MaxInt
	sufMin[len(nums)] = math.MaxInt

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < len(nums); i++ {
		preMin[i+1] = min(preMin[i], nums[i])
		// fmt.Printf("preMin[%d]: %d\n", i+1, preMin[i+1])
	}
	for i := len(nums) - 1; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], nums[i])
		// fmt.Printf("sufMin[%d]: %d\n", i, sufMin[i])
	}

	ans := math.MaxInt
	for index, j := range nums {
		i := preMin[index]
		k := sufMin[index+1]
		// fmt.Printf("i: %d, j: %d, k: %d\n", i, j, k)
		if i < j && k < j {
			ans = min(ans, i+j+k)
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans

}

// https://leetcode.cn/problems/unique-length-3-palindromic-subsequences/description/
//1930. 长度为 3 的不同回文子序列

func countPalindromicSubsequence(s string) int {
	set := [26][26]bool{}
	pre := [26]int{}
	suf := [26]int{}

	for i := 1; i < len(s); i++ {
		suf[s[i]-'a']++
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	pre[s[0]-'a']++
	lastIndex := len(s) - 1
	for i, v := range s {
		if i == 0 || i == lastIndex {
			continue
		}

		suf[v-'a']--
		for j := 0; j < 26; j++ {
			both := min(pre[j], suf[j])
			if both > 0 {
				set[j][v-'a'] = true
			}
		}
		pre[v-'a']++

	}
	ans := 0
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if set[i][j] {
				ans++
				// fmt.Println(string(i+'a'), string(j+'a'))
			}
		}
	}
	// fmt.Println(set)
	return ans
}

func numberOfRightTriangles(grid [][]int) int64 {
	ans := int64(0)
	xLen := len(grid)
	yLen := len(grid[0])

	xHave1S := make([][1001]int, xLen + 1)
	yHave1S := make([][1001]int, yLen + 1)
	for i := 1; i <= xLen; i++ {
		for j := 1;j <= yLen;j ++{
			xHave1S[i][j] = xHave1S[i][j-1] + grid[i-1][j-1]
			yHave1S[j][i] = yHave1S[j][i-1] + grid[i-1][j-1]
		}
	}
	// for i := 1; i <= xLen; i++ {
	// 	for j := 1; j <= yLen; j++ {
	// 		fmt.Printf("xHave1S[%d行][%d列]: %d\n", i, j, xHave1S[i][j])
	// 	}
	// }
	// for j := 1; j <= yLen; j++ {
	// 	for i := 1; i <= xLen; i++ {
	// 		fmt.Printf("yHave1S[%d列][%d行]: %d\n", j, i, yHave1S[j][i])
	// 	}
	// }



	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 0 {
				continue
			}
			var rightCount, leftCount, downCount, upCount int64
			curX := x + 1
			curY := y + 1
			rightCount = int64(xHave1S[curX][yLen] - xHave1S[curX][curY])
			leftCount = int64(xHave1S[curX][y])
			downCount = int64(yHave1S[curY][xLen] - yHave1S[curY][curX])
			upCount = int64(yHave1S[curY][x])
			fmt.Printf("x: %d, start: %d, rightCount: %d, leftCount: %d, downCount: %d, upCount: %d\n", x, y, rightCount, leftCount, downCount, upCount)
			ans += int64((downCount + upCount) * (leftCount + rightCount))
		}

	}
	return ans
}


