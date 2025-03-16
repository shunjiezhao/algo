package stack

import (
	"fmt"
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

	xHave1S := make([][1001]int, xLen+1)
	yHave1S := make([][1001]int, yLen+1)
	for i := 1; i <= xLen; i++ {
		for j := 1; j <= yLen; j++ {
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

func maximumTripletValue(nums []int) int64 {
	sufMax := make([]int, len(nums))
	sufMax[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], nums[i])
	}
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	preMax := int64(nums[0])
	var ans int64
	for i, v := range nums {
		if i == 0 || i == len(nums)-1 {
			continue
		}
		ans = max(ans, int64(preMax-int64(v))*int64(sufMax[i+1]))
		preMax = max(preMax, int64(v))
	}
	return ans
}

func subSort(array []int) []int {
	if len(array) == 0 {
		return []int{-1, -1}
	}
	type data struct {
		index int
		value int
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	stack := []data{{index: -1, value: math.MinInt}}
	tt := 0
	start := math.MaxInt
	end := -1
	for i, v := range array {
		if v >= stack[tt].value {
			stack = append(stack, data{index: i, value: v})
			tt++
			continue
		}
		// lastIndex := stack[tt].index
		j := tt
		for j > 0 && v < stack[j].value {
			// fmt.Printf("i: %d, v: %d, lastIndex: %d, array[lastIndex]: %d\n", i, v, stack[j].index, array[stack[j].index])
			j--
		}
		lastIndex := stack[j].index
		// stack = append(stack, data{index: lastIndex, value: v})
		// tt++

		start = min(start, lastIndex+1)
		end = i
	}
	if start == math.MaxInt {
		start = -1
	}
	return []int{start, end}
}

// https://leetcode.cn/problems/number-of-boomerangs/submissions/610090832/
// 447. 回旋镖的数量
func numberOfBoomerangs(points [][]int) int {
	ans := 0

	for i, v := range points {
		dis := make(map[int]int)
		for j, vi := range points {
			if j == i {
				continue
			}
			distance := (v[0]-vi[0])*(v[0]-vi[0]) + (v[1]-vi[1])*(v[1]-vi[1])
			dis[distance]++
		}
		for j, vi := range points {
			if j == i {
				continue
			}
			distance := (v[0]-vi[0])*(v[0]-vi[0]) + (v[1]-vi[1])*(v[1]-vi[1])
			dis[distance]--
			ans += dis[distance] * 2
		}
	}
	return ans
}

func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	stack := []int{
		nums[n-1],
	}
	secondMax := math.MinInt
	for i := n - 2; i >= 0; i-- {
		if nums[i] < secondMax {
			return true
		}
		// 找到,逆序对, 然后更新k能取的最大值
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			secondMax = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	type data struct {
		idx   int
		value int
	}
	grid := make([][]data, n)
	for _, edge := range edges {
		grid[edge[0]] = append(grid[edge[0]], data{idx: edge[1], value: edge[2]})
		grid[edge[1]] = append(grid[edge[1]], data{idx: edge[0], value: edge[2]})
	}
	var dfs func(cur int, father int, dis int, ans int) int
	dfs = func(cur int, father int, dis int, ans int) int {
		if dis%signalSpeed == 0 {
			ans++
		}
		// fmt.Printf("cur: %d, father: %d, dis: %d, ans: %d\n", cur, father, dis, ans)

		for _, v := range grid[cur] {
			if v.idx == father {
				continue
			}
			ans = dfs(v.idx, cur, dis+v.value, ans)
		}
		// fmt.Printf("return cur: %d, father: %d, dis: %d, ans: %d\n", cur, father, dis, ans)
		return ans
	}
	ansList := make([]int, n)

	for i := 0; i < n; i++ {
		// 当前是 c
		if len(grid[i]) < 2 {
			ansList[i] = 0
			continue
		}
		cnt := []int{}
		for _, v := range grid[i] {
			val := dfs(v.idx, i, v.value, 0)
			if val != 0 {
				cnt = append(cnt, val)
				fmt.Printf("i: %d, v.idx: %d, cnt: %d\n", i, v.idx, val)
			}
		}
		ans := 0
		for i, v := range cnt {
			for j := i + 1; j < len(cnt); j++ {
				ans += v * cnt[j]
			}
		}

		ansList[i] = ans

	}
	return ansList
}
