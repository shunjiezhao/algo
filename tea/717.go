package main

import (
	"bufio"
	. "fmt"
	"io"
)

func CF1195C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	var n, ans int
	Fscan(in, &n)
	a, b := make([]int, n+1), make([]int, n+1)
	f := make([][2]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &b[i])
	}
	for i := 1; i <= n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][1]+a[i])
		f[i][1] = max(f[i-1][1], f[i-1][0]+b[i])
	}

	for i := 1; i <= n; i++ {
		ans = max(ans, f[i][0])
		ans = max(ans, f[i][1])
	}
	Fprint(out, ans)
}
