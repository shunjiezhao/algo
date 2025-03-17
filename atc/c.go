package main


import (
	"fmt"
	"os"
)

func main() {
	var s string
    // 使用 fmt.Fscanf 从标准输入读取
    _, err := fmt.Fscanf(os.Stdin, "%s\n", &s)
    if err != nil {
        fmt.Fprintf(os.Stderr, "读取输入时出错: %v\n", err)
        return
    }

	var w []int
	by := []byte(s)
	for i, c := range by {
		if c == 'W'{
			w = append(w, i)
			continue
		}
		if c == 'A'{
			if len(w) == 0{
				continue
			}
			by[w[0]] = 'A'
			for _, v := range w[1:]{
				by[v] = 'C'
			}
			by[i] = 'C'
			w = w[:0]
		}
		w = w[:0]
	}
	fmt.Fprintf(os.Stdout, "%s\n", string(by))
}