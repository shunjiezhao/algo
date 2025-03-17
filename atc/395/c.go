package main


import (
	"fmt"
	"os"
	"math"
)

func main() {
	var n int
    // 使用 fmt.Fscanf 从标准输入读取
    _, err := fmt.Fscanf(os.Stdin, "%d\n", &n)
    if err != nil {
        fmt.Fprintf(os.Stderr, "读取输入时出错: %v\n", err)
        return
    }
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(os.Stdin, "%d", &arr[i])
		// fmt.Println(i, arr[i])
	}
	// fmt.Println(arr)
	cntMap := map[int]int{}
	cntMap[arr[0]] = 0
	ans := math.MaxInt
	for i :=1;i < n;i ++{
		idx, exist := cntMap[arr[i]]
		if exist{
			ans = min(ans, i - idx + 1)
		}else{
			cntMap[arr[i]] = i
		}
	}
	if ans == math.MaxInt{
		ans = -1
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
