package main

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	window := make([]int, 0, len(nums))
	type tmp struct{ len, l, r int }
	var ans tmp
	n := len(nums)
	for i := 0; i < n; i++ {
		sum += nums[i]
		window = append(window, nums[i])
		for len(window) > 0 && sum >= target {
			if ans.len == 0 || ans.len > len(window) {
				ans.len = len(window)
				ans.l, ans.r = i-len(window)-1, i
			}

			sum -= window[0]
			window = window[1:]
		} // move window

	}
	return ans.len
}

func lengthOfLongestSubstring(s string) int {
	cnt := map[byte]struct{}{}
	contain := func(c byte) bool {
		_, ok := cnt[c]
		return ok
	}

	n := len(s)
	ans := 0

	for l, r := 0, 0; r < n; r++ {
		for ; l <= r && contain(s[r]); l++ {
			delete(cnt, s[l])
		}
		cnt[s[r]] = struct{}{}
		if ans < (r - l + 1) {
			ans = r - l + 1
		}

	}
	return ans

}

func minWindow(s string, t string) string {
	ogn, now := map[byte]int{}, map[byte]int{}
	for i := range t {
		ogn[t[i]]++
	}
	type tmp struct{ len, l, r int }
	var ans tmp
	n := len(s)
	ans.len = n + 1
	isAnswer := func() bool {
		for k := range ogn {
			if ogn[k] < now[k] {
				return false
			}

		}
		return true
	}
	containInOgn := func(b byte) bool {
		_, ok := ogn[b]
		return ok
	}
	check := func(b byte) bool {
		if !containInOgn(b) {
			return false // t not contain this char
		}
		for k := range ogn {
			if ogn[k] > now[k] {
				return true
			}
		}
		return false
	}

	for l, r := 0, 0; r < n; r++ {
		for ; l <= r && check(s[r]); l++ {
			now[s[l]]--
		}
		now[s[r]]++

		for ; l <= r && !containInOgn(s[l]); l++ {
			now[s[l]]--
		}

		if isAnswer() {
			if ans.len < r-l+1 {
				ans.len = r - l + 1
				ans.l, ans.r = l, r+1
			}
		}

	}
	return s[ans.l:ans.r]
}

func findSubstring(s string, words []string) []int {
	ans := make([]int, 0)
	wMap := make(map[string]int)
	for _, v := range words {

		wMap[v]++

	}

	wcnt, wLen := len(words), len(words[0])
	for l := 0; l < wLen && l < len(s)-wcnt*wLen+1; l++ {
		cnt := make(map[string]int)

		descKey := func(s string) {
			cnt[s]--
			if cnt[s] == 0 {
				delete(cnt, s)
			}
		}
		incKey := func(s string) {
			cnt[s]++
			if cnt[s] == 0 {
				delete(cnt, s)
			}
		}

		for i := l; i < l+wcnt*wLen; i += wLen {
			incKey(s[i : i+wLen])
		}

		for _, v := range words {
			descKey(v)
		}
		// fmt.Println(cnt)
		// last word + 1
		for i := l; i <= len(s)-wcnt*wLen; i += wLen {
			// old word
			if i != l {
				descKey(s[i-wLen : i])
				end := i + wcnt*wLen
				incKey(s[end-wLen : end])
				// fmt.Println("delete[", i-wLen,i,"] add",end-wLen, end, "]")
			}
			// fmt.Println(i, cnt)
			if len(cnt) == 0 {
				ans = append(ans, i)
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	return -max(-a, -b)
}
