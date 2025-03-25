package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	k:=make([]int,n)
	cntMapList := make([]map[int]int,n)
	
	for i := 0; i < n; i++ {
		cntMap := make(map[int]int)
		fmt.Fscanf(reader, "%d ", &k[i])
		for j := 0; j < k[i]; j++ {
			var a int
			fmt.Fscanf(reader, "%d ", &a)
			cntMap[a]++
		}
		cntMapList[i] = cntMap
	}

	ans := float64(0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := k[i] * k[j]
			eq := 0
			for k := range cntMapList[i] {
				if cntMapList[j][k] > 0 {
					eq += cntMapList[i][k] * cntMapList[j][k]
				}
			}
			if tmp := float64(eq) / float64(sum); tmp > ans {
				ans = tmp
			}
		}
	}
	
	fmt.Fprintf(writer, "%.16f\n", ans)
}
