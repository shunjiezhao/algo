package main

import (
	"fmt"
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

			if r-i == 2 &&  path[i+1] == '.' { // ..
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
