package main

import "fmt"

func main() {
	var n, l, r, ql, qr int
	fmt.Scan(&n, &l, &r, &ql, &qr)
	var pre, suf int
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
		suf += a[i]
	}

	var ans = suf * r + (n - 1) * qr
	for i, v := range a{
		suf -= v
		pre += v
		cost := pre *l + suf * r
		lcnt, rcnt := i + 1, n - 1 - i
		// [...] -> lcnt
		// [....] -> rcnt
		// max - min = need
		if lcnt + 1 < rcnt{
			cost += (rcnt - lcnt + 1) * qr
		}else{
			cost += (lcnt - rcnt - 1) * ql
		}
		if cost < ans {
			ans = cost
		}
	}
	fmt.Println(ans)
}
