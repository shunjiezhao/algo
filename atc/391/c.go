package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, q int
	fmt.Fscanf(reader, "%d %d\n", &n, &q)
	idxMap := make([]int, n+1)
	for i := 1; i <= n; i++ {
		idxMap[i] = i
	}
	var cnt int
	cntMap := make([]int, n+1)
	for i := range cntMap {
		cntMap[i] = 1
	}
	for i := 0; i < q; i++ {
		var op int
		fmt.Fscanf(reader, "%d", &op)
		if op == 1 {
			var p, h int
			fmt.Fscanf(reader, "%d %d\n", &p, &h)
			cntMap[idxMap[p]]--
			if cntMap[idxMap[p]] == 1 {
				cnt--
			}
			idxMap[p] = h
			cntMap[h]++
			if cntMap[h] == 2 {
				cnt++
			}
		} else if op == 2 {
			fmt.Fscanf(reader, "\n")
			fmt.Fprintf(writer, "%d\n", cnt)
		}
	}

}
