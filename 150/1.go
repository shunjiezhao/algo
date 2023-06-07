package main

import (
	"math"
	"math/rand"
	"sort"
	"strings"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nn := len(nums1)
	n--
	m--
	for i := nn - 1; i >= 0; i-- {
		if n <= 0 {
			break
		}
		if m < 0 || nums2[n] > nums1[m] {
			nums1[i] = nums2[n]
			n--
		} else {
			nums1[i] = nums1[m]
			m--
		}
	}
}

func removeElement(nums []int, val int) int {
	l, r := 0, len(nums)
	for l < r {
		if nums[l] == val {
			nums[l], nums[r] = nums[r-1], nums[l]
			r--
		} else {
			l++
		}
	}
	return l
}

func removeDuplicates1(nums []int) int {
	return removeDuplicates(nums, 1)
}

func removeDuplicates(nums []int, allowDupCnt int) int {
	n := len(nums)
	l, r := 0, 0
	for ; r < n; r++ {
		t := r
		for r < n && nums[r] == nums[t] {
			r++
		}
		for k := 0; k < min(r-t, allowDupCnt); k++ {
			nums[l] = nums[t]
			l++
		}
		r--
	}
	return l
}

func majorityElement(nums []int) int {
	var down func(l, r int) int
	down = func(l, r int) int {
		if l >= r {
			return nums[l]
		}

		mid := (l + r) / 2
		left, right := down(l, mid), down(mid+1, r)
		if left == right {
			return left
		}
		cntInCloseRange := func(l, r, num int) (count int) {
			for i := l; i <= r; i++ {
				if nums[i] == num {
					count++
				}
			}
			return
		}
		leftCount, rightCount := cntInCloseRange(l, r, left), cntInCloseRange(l, r, right)
		if leftCount > rightCount {
			return left
		}
		return right
	}
	return down(0, len(nums)-1)
}

func rotate(nums []int, k int) {
	reverse := func(l, r int) {
		for i := 0; i < (r-l+1)/2; i++ {
			nums[i+l], nums[r-i] = nums[r-i], nums[i+l]
		}
	}
	k %= len(nums)
	if k == 0 {
		return
	}
	reverse(0, len(nums)-1)
	reverse(k, len(nums)-1)
	reverse(0, k-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	return sort.Search(len(citations), func(i int) bool {
		return citations[i] < i+1
	})
}

type RandomizedSet struct {
	idx  map[int]int
	nums []int
}

func Constructor() RandomizedSet {

	return RandomizedSet{
		idx: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idx[val]; ok {
		return false
	}

	this.idx[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.idx[val]; !ok {
		return false

	}

	i, lst := this.idx[val], this.nums[len(this.nums)-1]
	this.nums[i] = lst
	this.idx[lst] = i
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.idx, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = 1
		if i != 0 {
			ans[i] = ans[i-1] * v
		}
	}

	suf := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		ans[i] *= suf
		suf *= nums[i]
	}

	return ans
}

func canCompleteCircuit(gas []int, cost []int) int {
	sum, mIndex, minVal := 0, 0, math.MaxInt
	for i := 0; i < len(cost); i++ {
		sum += gas[i] - cost[i]
		if sum < minVal {
			mIndex, minVal = i, sum
		}
	}
	if minVal < 0 {
		return -1
	}
	return (mIndex + 1) % len(gas)
}

func trap(height []int) int {
	n := len(height)
	pre, suf := make([]int, 2+n), make([]int, 2+n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		return -max(-a, -b)
	}

	for i := n; i > 0; i-- {
		suf[i] = max(suf[i+1], height[i-1])
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		pre[i] = max(pre[i-1], height[i-1])
		ans += max(0, min(pre[i], suf[i])-height[i-1])
	}
	return ans

}

func candy(ratings []int) int {
	n := len(ratings)
	left := make([]int, n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	left[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	r := 1
	ans := max(1, left[n-1])
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			r++
		} else {
			r = 1
		}
		ans += max(r, left[i])
	}
	return ans
}

func intToRoman(num int) string {
	var valueSymbols = []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var ans string
	for _, v := range valueSymbols {
		for num > v.value {
			ans += v.symbol
		}
		if num == 0 {
			break
		}
	}
	return ans
}
func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	return len(s) - strings.LastIndex(s, " ")
}

func longestCommonPrefix(strs []string) string {
	var index int
	for {
		for _, v := range strs {
			if len(v) <= index || v[index] != strs[0][index] {
				goto done // not have
			}
		}
		index++
	}
done:
	return strs[0][:index]
}

func reverseWords(s string) string {
	sArray := make([]string, 0)
	reverse := func(n int, swap func(i, j int)) {
		for i := 0; i < n/2; i++ {
			swap(i, n-i-1)
		}
	}
	n := len(s)
	for i := 0; i < n; i++ {
		j := i + 1
		for j < n && s[j] == ' ' {
			j++
		}
		start := j
		for j < n && s[j] != ' ' {
			j++
		}
		if start == j {
			break
		}
		sArray = append(sArray, string(s[start:j]))

		i = j - 1
	}

	reverse(len(sArray), func(i, j int) {
		sArray[i], sArray[j] = sArray[j], sArray[i]
	})

	return strings.Join(sArray, " ")
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	index := make([]string, numRows)
	down, up := 0, 1

	d := down // down
	r := 0
	for _, c := range s {
		if r == numRows {
			d = up
			r -= 2
		}
		if r < 0 {
			d = down
			r += 2
		}
		index[r] += string(c)
		if d == down {
			r++
		} else {
			r--
		}
	}
	b := strings.Builder{}
	for _, v := range index {
		b.WriteString(v)
	}
	return b.String()
}

func kmpGetNext(s []byte) []int {
	n := len(s)
	next := make([]int, n)
	for i, j := 2, 0; i < len(s); i++ {
		for j > 0 && s[j+1] != s[i] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
	return next
}
func str2bytes(s string) []byte {
	n := len(s)
	sb := make([]byte, 1, n+1)
	sb = append(sb, []byte(s)...)
	return sb
}

func kmp(s, p string) int{
	sb, pb := str2bytes(s), str2bytes(p)
	next := kmpGetNext(pb)

	for i, j := 1, 0; i < len(sb); i++ {
		for j > 0 && sb[i] != pb[j+1] {
			j = next[j]
		}
		if sb[i] == pb[j + 1]{
			j++
		}
		if j == len(p) {
			return i - len(p) 
		}
	}
	return -1
}

func strStr(haystack string, needle string) int {
	return kmp(haystack, needle)
}