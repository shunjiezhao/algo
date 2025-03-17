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

	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var ans int
	mp :=map[string]bool{}
	

	for i := 0; i < m; i++ {
		var from, to  int
		fmt.Fscanf(reader, "%d %d\n", &from, &to)
		if from == to{
			ans ++
			continue
		}
		key1 := fmt.Sprintf("%d->%d", from, to)
		if mp[key1]{
			ans ++
			continue
		}
		key2 := fmt.Sprintf("%d->%d", to, from)
		if mp[key2]{
			ans ++
			continue
		}
		mp[key1] = true
	}
	fmt.Fprintf(writer, "%d\n", ans)
}
