package main

import (
	"fmt"
	"math"
	"sort"
)

func isGood(nums []int) bool {
	sort.Ints(nums)
	cnt := make(map[int]int)
	big := len(nums) - 1
	if nums[len(nums)-1] != big {
		return false
	}
	for _, num := range nums {
		cnt[num]++
		if cnt[num] > 1 && num != big {
			return false
		}

	}
	return cnt[big] == 2
}

func sortVowels(s string) string {
	cnt := make(map[byte]bool)
	for _, i := range []byte{'a', 'e', 'i', 'o', 'u'} {
		cnt[i] = true
		cnt[i-'a'+'A'] = true
	}
	fmt.Println(cnt)

	idx := make(map[byte]int)
	all := make([]int, 0)
	for i := range s {
		b := s[i]
		if cnt[b] == true {
			idx[b]++
			all = append(all, i)
		}
	}
	var k []byte
	for key := range idx {
		k = append(k, key)
	}
	sort.Slice(k, func(i, j int) bool {
		return s[i] < s[j]
	})

	t := []byte(s)
	i := 0
	for _, key := range k {
		for j := 0; j < idx[key]; j++ {
			t[all[i]] = key
			i++
		}

	}
	return string(t)
}
func maxScore(nums []int, x int) int64 {
	n := len(nums)
	dp := make([][2]int64, n)
	getLow := func(i int) int {
		return i & 1
	}

	dp[0][getLow(nums[0])] = int64(nums[0])

	for i := 1; i < n; i++ {
		low := getLow(nums[i])
		Xlow := 1 ^ low

		dp[i][low] = max(dp[i-1][Xlow]-int64(x), dp[i-1][low]) + int64(nums[i])
		dp[i][Xlow] = dp[i-1][Xlow]
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func numberOfWays(n int, x int) int {
	f := make([]int64, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		pow := int64(math.Pow(float64(i), float64(x)))
		for j := int64(n); j >= pow; j-- { // 注意 f[i][j] 表示 前 i 个数，和为 j 的方案数，而物品的集合有 {i^x | i ^ x < n }，然后我们就利用 01 背包进行选择即可
			f[j] += f[j-pow]
		}
	}
	return int(f[n] % (1e9 + 7))
}
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
