package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isValid(s string) bool {
	stk := make([]rune, 0)
	mp := make(map[rune]rune)
	mp['{'] = '}'
	mp['('] = ')'
	mp['['] = ']'
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			stk = append(stk, v)
		} else {
			if len(stk) > 0 {
				l := len(stk) - 1
				t := stk[l]
				if mp[t] == v {
					stk = stk[:l]
					continue
				}
			}
			stk = append(stk, v)
		}
	}
	return len(stk) == 0

}

func simplifyPath(path string) string {
	stk := make([]string, 0)
	if path[0] == '/' {
		path = path[1:]
	}

	find := func(l int) int {
		var i int
		for i = l; i < len(path); i++ {
			if path[i] == '/' {
				break
			}
		}
		return i
	}
	n := len(path)
	for i := 0; i < n; i++ {
		v := path[i]
		switch v {
		case '.':
			r := find(i)
			if r == i+1 { // .
				continue
			}

			if r-i == 2 && path[i+1] == '.' { // ..
				if len(stk) >= 1 {
					stk = stk[:len(stk)-1] // pop

				}
				continue
			}

			stk = append(stk, path[i:r]) // valid path name like ...
			i = r - 1
		case '/':
			for i < n && path[i] == '/' {
				i++
			}
			i--
		default:
			r := find(i)
			stk = append(stk, path[i:r])
			i = r // skip '/'
		}
		fmt.Println(stk)
	}
	t := strings.Builder{}

	for _, v := range stk {
		t.WriteString("/" + v)
	}
	if t.String() == "" {
		return "/"
	}

	return t.String()
}

type MinStack struct {
	stk, MinStk []int
}

func Constructor() MinStack {
	this := MinStack{}
	this.MinStk = append(this.MinStk, math.MaxInt)
	return this
}

func (this *MinStack) Push(val int) {
	this.stk = append(this.stk, val)
	t := this.MinStk[len(this.MinStk)-1]
	if t > val {
		t = val
	}
	this.MinStk = append(this.MinStk, t)
}

func (this *MinStack) Pop() {
	this.stk = this.stk[:len(this.stk)-1]
	this.MinStk = this.MinStk[:len(this.MinStk)-1]
}

func (this *MinStack) Top() int {
	return this.stk[len(this.stk)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinStk[len(this.MinStk)-1]
}

func evalRPN(tokens []string) int {
	stk := make([]string, 0)
	for _, v := range stk {
		var l, r int
		switch v {
		case "/", "*", "+", "-":
			r, _ = strconv.Atoi(stk[len(stk)-1])
			l, _ = strconv.Atoi(stk[len(stk)-2])
			stk = stk[:len(stk)-2]
		default:
			stk = append(stk, v)
		}
		if v == "/" {
			stk = append(stk, strconv.FormatInt(int64(l/r), 10))
		} else if v == "*" {
			stk = append(stk, strconv.FormatInt(int64(l*r), 10))
		} else if v == "+" {
			stk = append(stk, strconv.FormatInt(int64(l+r), 10))
		} else if v == "-" {
			stk = append(stk, strconv.FormatInt(int64(l-r), 10))
		}
		fmt.Println(stk)
	}
	ans, _ := strconv.Atoi(stk[len(stk)-1])
	return ans
}

// https://oi-wiki.org/misc/expression/ 
func calculate(s string) int {
	stk, op := make([]int, 0), make([]rune, 0)

	isOp := func(op rune) bool {
		switch op {
		case '*', '+', '-', '/':
			return true
		}
		return false
	}
	priorityOp := func(op rune) int {
		if op < 0 {
			return 3
		}
		switch op {
		case '+', '-':
			return 1
		case '*', '/':
			return 2
		}
		return -1
	}
	processOp := func(op rune) {
		if op < 0{
			l := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			op = -op
			if op == '+'{
				stk = append(stk, l)
			}else{
				stk = append(stk, -l)
			}
			return
		}
		r, l := stk[len(stk)-1], stk[len(stk)-2]
		stk = stk[:len(stk)-2]
		switch op {
		case '*':
			stk = append(stk, l*r)
		case '/':
			stk = append(stk, l/r)
		case '+':
			stk = append(stk, l+r)
		case '-':
			stk = append(stk, l-r)
		}
	}

	isUnary := func(op rune) bool {
		if op == '+' || op == '-' {
			return true
		}
		return false
	}
	mayUnary := true
	for i := 0;i < len(s);i ++{
		v := rune(s[i])
		if v == ' ' {
			continue
		}

		if v == '(' {
			op = append(op, v)
			mayUnary = true
		} else if v == ')' {
			for len(op) > 0 && op[len(op)-1] != '(' {
				processOp(op[len(op)-1])
				op = op[:len(op)-1]
			}
			op = op[:len(op)-1]
		} else if isOp(v) {
			curOp := v
			if mayUnary && isUnary(v) {
				curOp = -curOp
			}

			for len(op) > 0 &&
				((curOp >= 0 && priorityOp(op[len(op)-1]) >= priorityOp(curOp)) ||
					(curOp < 0 && priorityOp(op[len(op)-1]) > priorityOp(curOp))) {
				processOp(op[len(op)-1])
				op = op[:len(op)-1]
			}
			op = append(op, curOp)
			mayUnary = true
		}else{
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9'{
				num = num * 10 + int(s[i] -'0')
				i++
			}
			i--
			stk = append(stk, num)
			mayUnary = false
		}
	}
	for len(op) != 0{
		processOp(op[len(op)-1])
		op = op[:len(op)-1]
	}
	return stk[len(stk)-1]
}