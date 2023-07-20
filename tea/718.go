package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"io"
	"math"
	"os"
)

func CF191A(_r io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	in := bufio.NewReader(_r)
	var n, ans int
	var a []byte
	f := [26][26]int{}

	dp := [26][26]int{}
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			dp[i][j] = math.MinInt32
		}
	}
	for fmt.Fscan(in, &n); n > 0; n-- {
		fmt.Fscan(in, &a)
		// start
		l := a[0] - 'a'
		r := a[len(a)-1] - 'a'
		// s[l:r] -> 右端点为r
		// 左端点就是需要遍历
		for j := 0; j < 26; j++ {
			dp[j][r] = max(dp[j][r], dp[j][l]+len(a))
		}
		dp[l][r] = max(dp[l][r], len(a))

	}
	for i := 0; i < 26; i++ {
		ans = max(ans, dp[i][i])
		//fmt.Println(string(i+'a'), ans)
	}

	for i := 0; i < 26; i++ {
		ans = max(ans, f[i][i])
	}
	fmt.Fprint(out, ans)
}

func main2() { CF191A(os.Stdin, os.Stdout) }
