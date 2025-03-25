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

	var n int64
	fmt.Fscanf(reader, "%d\n", &n)
	end := math.Cbrt(float64(n))
	for d := int64(1); d <= int64(end); d++ {
		if n%d != 0 {
			continue
		}

		k := calc(3, 3*d, d*d-n/d)
		if k != -1 {
			fmt.Fprintf(writer, "%d %d\n", k+d, int64(k))
			return
		}
	}
	fmt.Fprintf(writer, "-1\n")
	return
}

func calc(a, b, c int64) int64 {
	var l, r int64 = 1, 600000001
	for l < r {
		mid := (l + r) / 2
		if a*mid*mid+b*mid+c >= 0 {
			r = mid
		} else {
			l = mid + 1
		}

	}
	if l*l*a+l*b+c == 0 {
		return l
	}
	return -1
}
