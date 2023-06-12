package main

import (
	"sort"
	"strings"
)

func canConstruct(ransomNote string, magazine string) bool {
	cnt := make(map[byte]int)
	for i := range magazine {
		cnt[magazine[i]]++
	}

	for i := range ransomNote {
		if _, ok := cnt[ransomNote[i]]; !ok {
			return false
		}
		cnt[ransomNote[i]]--
		if cnt[ransomNote[i]] == 0 {
			delete(cnt, ransomNote[i])
		}
	}
	return true
}

// 双射
func isIsomorphic(s string, t string) bool {
	mp1, mp2 := make(map[byte]byte), make(map[byte]byte)
	for i := range s {
		c1, ok := mp1[t[i]]
		if !ok {
			mp1[t[i]] = s[i]
			c1 = s[i]
		}
		c2, ok := mp2[s[i]]
		if !ok {
			mp2[s[i]] = t[i]
			c2 = t[i]
		}
		if c1 != s[i] || c2 != t[i] {
			return false
		}
	}
	return true
}

// 双射
func wordPattern(pattern string, s string) bool {
	mp1 := make(map[byte]string)
	mp2 := make(map[string]byte)
	sA := strings.Split(s, " ")
	if len(sA) != len(pattern) {
		return false
	}
	for i := range pattern {
		c1, ok := mp1[pattern[i]]
		if !ok {
			mp1[pattern[i]] = sA[i]
			c1 = sA[i]
		}
		c2, ok := mp2[sA[i]]
		if !ok {
			mp2[sA[i]] = pattern[i]
			c2 = pattern[i]
		}
		if c1 != sA[i] || c2 != pattern[i] {
			return false
		}
	}
	return true

}

func isAnagram(s string, t string) bool {
	cnt1, cnt2 := make(map[byte]int), make(map[byte]int)
	for i := range s {
		cnt1[s[i]]++
	}
	for i := range t {
		cnt2[t[i]]++
	}
	for k := range cnt1 {
		if cnt1[k] != cnt2[k] {
			return false
		}
	}
	for k := range cnt2 {
		if cnt1[k] != cnt2[k] {
			return false
		}
	}
	return true
}

func groupAnagrams(strs []string) [][]string {
	cnt := make(map[string][]string)

	for _, v := range strs {
		bts := []byte(v)
		sort.Slice(bts, func(i, j int) bool {
			return bts[i] < bts[j]
		})
		sTr := string(bts)
		cnt[sTr] = append(cnt[sTr], v)
	}
	ans := make([][]string, 0, len(cnt))
	for _, v := range cnt {
		ans = append(ans, v)
	}

	return ans
}

func twoSum(nums []int, target int) []int {
	mp := map[int]int{}
	for i, v := range nums {
		if t, ok := mp[target-v]; ok {
			return []int{t, i}
		}
		mp[v] = i
	}
	return []int{-1, -1}
}

func isHappy(n int) bool {
	cnt := make(map[int]int)
	getNext := func(a int) (next int) {
		for ; a > 0; a /= 10 {
			next += (a % 10) * (a % 10)
		}
		return
	}

	for n != 1 && cnt[n] == 0 {
		cnt[n]++
		n = getNext(n)
	}
	return n == 1
}

func containsNearbyDuplicate(nums []int, k int) bool {
	idx := make(map[int][]int)
	for i, v := range nums {
		idx[v] = append(idx[v], i)
	}
	for _, v := range idx {
		for i := 1; i < len(v); i++ {
			if v[i]-v[i-1] <= k {
				return true
			}
		}
	}
	return false
}

func longestConsecutive(nums []int) int {
	idx := make(map[int]int)

	for i, v := range nums {
		t, ok := idx[v]
		if ok {
			continue
		}
		idx[v] = t
		// left
		l, ok := idx[v-1]
		if ok {
			idx[l] = i
		}
		// right
		r, ok := idx[v+1]
		if ok {
			idx[v] = r
		}
	}

	var find func(a int) int
	// like union set
	find = func(a int) int {
		if nums[idx[a]] != a {
			idx[a] = find(nums[idx[a]])
		}
		return idx[a]
	}


	ans := 0
	for k, v := range idx{
		t := find(k)
		if ans < t - v + 1{
			ans = t - v + 1
		}
	}
	return ans
}
