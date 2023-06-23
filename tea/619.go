package main

import (
	"fmt"
)

func main619() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	var ans int
	cnt := make(map[int]int)
	l := 0
	for i, v := range a {
		cnt[v]++
		if len(cnt) > 2 {
			for l < i && len(cnt) > 2 {
				cnt[a[l]]--
				if cnt[a[l]] == 0 {
					delete(cnt, a[l])
				}
				l++
			}
		}
		if ans < i-l+1 {
			ans = i - l + 1
		}
	}
	fmt.Println(ans)
}

func longestSubarray(a []int, limit int) int {
		qMin := make([]int, 0)
	qMax := make([]int, 0)
	ans := 0

	l := 0
	for r, v := range a {
		for len(qMin) > 0 && qMin[len(qMin)-1] < v {
			qMin = qMin[:len(qMin)-1]
		}
		for len(qMax) > 0 && qMax[len(qMax)-1] > v {
			qMax = qMax[:len(qMax)-1]
		}
		qMin = append(qMin, v)
		qMax = append(qMax, v)


		for len(qMin) > 0 && len(qMax) > 0 && qMax[0] - qMin[0] > limit{
			if a[l] == qMin[0]{
				qMin = qMin[1:]
			}
			if a[l] == qMax[0]{
				qMax = qMax[1:]
			}
			l++
		}
		if ans < r -l+1{
			ans = r-l+1
		}
	}
	return ans
}
