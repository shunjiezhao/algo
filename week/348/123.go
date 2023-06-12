package main

import (
	"sort"
	"strconv"
)

func isFascinating(n int) bool {
	q := strconv.Itoa(n*3) + strconv.Itoa(n*2) + strconv.Itoa(n)
	cnt := make(map[int]int)
	for i := 0; i < len(q); i++ {
		cnt[int(q[i]-'0')]++
	}
	for i := 1; i <= 9; i++ {
		if cnt[i] != 1 {
			return false
		}
	}
	return true
}

func longestSemiRepetitiveSubstring(s string) int {
	n := len(s)
	idx := make([]int, 0)
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			idx = append(idx, i-1)
		}
	}

	start, ans := 0, 0
	for i, v := range idx {
		if i == 0 {
			continue
		}
		if v-start+1 > ans {
			ans = v - start + 1
		}
		start = idx[i-1] + 1
	}
	if n-start > ans {
		ans = n - start
	}
	return ans
}

func sumDistance(nums []int, s string, d int) int {
	for i, v := range s {
		if v == 'L' {
			nums[i] -= d
		} else {
			nums[i] += d
		}
	}

	sort.Ints(nums)
	sum := 0
	mod := int(10e9 + 7)
	ans := 0
	for i, v := range nums {
		ans = (ans + sum + (v*i)%mod) % mod
		sum = (sum + v) % mod
	}
	return ans
}

func getLastMoment(n int, left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	ans := 0
	if len(right) != 0 {
		ans = n - right[0]
	}

	if len(left) != 0 {
		if ans < left[len(left)-1] {
			ans = left[len(left)-1]
		}
	}
	return ans

}