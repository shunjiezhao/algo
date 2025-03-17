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

	var n, q int
	fmt.Fscanf(reader, "%d %d\n", &n, &q)

	geziIndex := make([]int, n+1)
	room2po := make([]int, n+1)
	po2room := make([]int, n+1)

	for i := 1; i <= n; i++ {
		room2po[i] = i
		po2room[i] = i
		geziIndex[i] = i
	}

	for i := 0; i < q; i++ {
		var a, b, op int
		fmt.Fscanf(reader, "%d", &op)
		if op == 1 || op == 2 {
			fmt.Fscanf(reader, "%d %d\n", &a, &b)
		} else {
			fmt.Fscanf(reader, "%d\n", &a)
		}
		if op == 1 {
			geziIndex[a] = room2po[b]
		}
		if op == 2 {
			ra, rb := room2po[a], room2po[b]
			room2po[a], room2po[b] = room2po[b], room2po[a]
			po2room[ra], po2room[rb] = po2room[rb], po2room[ra]
		}
		if op == 3 {
			fmt.Fprintf(writer, "%d\n", po2room[geziIndex[a]])
		}
	}
}
