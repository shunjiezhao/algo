package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Fscanf(os.Stdin, "%s\n", &s)
	by := []byte(s)
	stack := []byte{}
	pair := map[byte]byte{
		')': '(',
		']': '[',
		'>': '<',
	}
	for _, v := range by{
		if v == '(' || v == '[' || v == '<' || len(stack) == 0{
			stack = append(stack, v)
			continue
		}
		// 找
		if stack[len(stack) -1] == pair[v]{
			stack = stack[:len(stack) -1]
		}else{
			stack = append(stack, v)
		}
	}
	// for _, v := range stack{
	// 	fmt.Println(string(v))
	// }
	if len(stack) == 0{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}

}
