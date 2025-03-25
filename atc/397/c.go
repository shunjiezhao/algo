package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	cnt := make(map[int]int)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &arr[i])
		cnt[arr[i]]++
	}
	ans := math.MinInt
	preCnt := make(map[int]int)
	for _, x := range arr {
		cnt[x]--
		if cnt[x] == 0 {
			delete(cnt, x)
		}
		preCnt[x]++
		disCount := len(preCnt) + len(cnt)
		ans = max(ans, disCount)
	}
	fmt.Fprintf(writer, "%d\n", ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
