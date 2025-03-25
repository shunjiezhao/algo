package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"math"
	"os"
	"sort"
)
// 串葫芦
// 算横着的一排的最大价值 (价值:y坐标值)
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, w int
	fmt.Fscanf(reader, "%d %d\n", &n, &w)
	x,y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &x[i], &y[i])
		x[i]--
		y[i]--
	}
	// fmt.Println("======",x, y)
	// t := make([]int, n)
	line := make([][]int, w)
	for i := 0;i < n;i ++{
		line[x[i]] = append(line[x[i]], i)
	}
	for x := range line{
		sort.Slice(line[x], func(i, j int) bool {
			return y[line[x][i]] < y[line[x][j]]
		})
	}
	// fmt.Println("======",line)
	no := make([]int, n)
	for i := 0;i < n;i ++{
		no[i] = math.MaxInt
	}
	for i := 0; i < int(10e9); i++ {
		done := false
		max := 0
		for j := 0;j < w;j ++{
			if len(line[j]) <= i{
				done = true
				break
			}
			if y[line[j][i]] > max{
				max = y[line[j][i]]
			}
		}
		if done{
			break
		}
		for j := 0;j < w;j ++{
			no[line[j][i]] = max + 1
		}
		// fmt.Printf("j:%d max:%d\n", i, max)
	}
	// for i := 0;i < n;i ++{
	// 	fmt.Fprintf(writer, "no %d %d\n", i, no[i])
	// }

	var q int
	fmt.Fscanf(reader, "%d\n", &q)
	for i := 0;i < q;i ++{
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		b--
		if no[b] > a{
			fmt.Fprintf(writer, "Yes\n")
		}else{
			fmt.Fprintf(writer, "No\n")
		}
	}
}
