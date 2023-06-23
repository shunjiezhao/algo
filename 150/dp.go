package main

import "sort"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := range s {
		i++
		for _, v := range wordDict {
			le := len(v)
			if i < le {
				continue
			}
			if s[i-le:i] == v && dp[i-le] {
				dp[i] = true
			}
		}
	}
	return dp[n]
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if v > i || dp[i-v] == -1 {
				continue
			}
			if dp[i] == -1 || dp[i] > dp[i-v]+1 {
				dp[i] = dp[i-v] + 1
			}
		}
	}
	return dp[amount]
}

func lengthOfLIS(nums []int) int {
	q := make([]int, 0)
	for _, v := range nums {
		if len(q) == 0 || q[len(q)-1] < v {
			q = append(q, v)
		} else {
			t := sort.Search(len(q), func(i int) bool {
				return q[i] >= v
			})
			if q[t] == v {
				t--
			}
			q[t+1] = v
		}
	}
	return len(q)
}

func rob(nums []int) int {
	f := make([]int, len(nums)+1)
	f[0] = 0
	f[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		f[i] = f[i-2] + nums[i-1]
		if f[i] < f[i-1] {
			f[i] = f[i-1]
		}
	}
	return f[len(nums)]
}
