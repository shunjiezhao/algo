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
	seen := make([]int,n + 1)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d", &seen[i])
	}
	fmt.Fscanf(reader, "%d")
	wear := make([]int,n + 1)
	wearF := make([]int,n + 1)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d", &wear[i])
		wearF[wear[i]] = i
		// fmt.Fprintf(writer, "%d %d \n", wear[i], wearF[wear[i]])
	}
	// fmt.Fprintf(writer, "%v\n", seen)
	// fmt.Fprintf(writer, "%v\n", wear)
	// fmt.Fprintf(writer, "%v\n", wearF)
	for i := 1; i <= n; i++ {
		// fmt.Fprintf(writer, "%d %d %d %d\n", i, wearF[i],seen[wearF[i]], wear[seen[wearF[i]]])
		fmt.Fprintf(writer, "%d ", wear[seen[wearF[i]]])
	}
}
