package week

import (
	"container/heap"
	"math"
	"sort"

	"google.golang.org/genproto/googleapis/container/v1"
)

func findNonMinOrMax(nums []int) int {
	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums)-1]

	for _, v := range nums {
		if v != min && v != max {
			return v
		}
	}
	return -1
}

func smallestString(s string) string {
	bts := []byte(s)
	for i := range bts {
		if bts[i] != 'a' {
			bts[i]--
		}
	}
	return string(bts)
}

func minCost(a []int, x int) int64 {
	n := len(a)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i * x
	}

	for i := 0; i < n; i++ {
		mn := i
		for j := 0; j < n; j++ {
			idx := (i + j) % n
			if a[mn] > a[idx] {
				mn = idx
			}
			s[j] += a[mn]
		}
	}
	mn := math.MaxInt
	for i := range s {
		if mn < s[i] {
			mn = s[i]
		}
	}
	return int64(mn)
}

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	type pair struct{ u, d, sum int }
	n := len(nums1)
	a := make([]pair, n)
	for i := range nums1 {
		a[i] = pair{nums1[i], nums2[i], nums1[i] + nums2[i]}
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].u != a[j].u {
			return a[i].u < a[j].u
		}
		return a[i].d < a[j].d
	})
	type pair2 struct {
		query []int
		i     int
	}
	q := make([]pair2, n)
	for i := range queries {
		q[i] = pair2{queries[i], i}
	}
	sort.Slice(q, func(i, j int) bool { return q[i].query[0] > q[j].query[0] })

	ans := make([]int, len(queries))
	stk, tt := make([]pair, n), 0
	r := n - 1
	for i := range q {
		x, y := q[i].query[0], q[i].query[1]
		for ; r >= 0 && a[r].u > x; r-- {
			for tt != 0 && (stk[tt].d < a[r].d && stk[tt].sum < a[r].sum) {
				tt--
			}
			tt++
			stk[tt] = a[r]
		}

		i := sort.Search(tt, func(i int) bool { return stk[i].d >= y })
		if i == n {
			ans[q[i].i] = -1
		} else {
			ans[q[i].i] = stk[i].sum
		}
	}
	return ans
}

