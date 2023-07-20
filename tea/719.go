package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"io"
	"os"
	"sort"
)

func CF1765D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int64
	fmt.Fscan(in, &n, &m)
	a := make([]int64, n)
	ans := n
	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &a[i])
		ans += a[i]
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	l, r := int64(0), n-1
	for l < r {
		if a[l]+a[r] <= m {
			l++ // 如果还有下一个
			if l < r {
				ans -= 2 // 需要将本次的观看时间减去
				// 因为 a[l - 1] + a[r] <= m  ->
				// 							a[l - 1] + a[r - 1] <= m
			} else {
				ans-- // 如果没有下一个，那么只需要减去一个 a[r] 的下载时间
			}
		}
		r--
	}
	fmt.Fprint(out, ans)

}

func main() { CF1765D(os.Stdin, os.Stdout) }
