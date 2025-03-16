package stack

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int][]int)
	for i, v := range nums {
		numMap[v] = append(numMap[v], i)
	}

	for i, v := range nums {
		need := target - v
		if jlist, ok := numMap[need]; ok {
			for _, j := range jlist {
				if i != j {
					return []int{i, j}
				}
			}
		}
	}
	return []int{-1, -1}
}

func numIdenticalPairs(nums []int) int {
	ans := 0
	numMap := make(map[int][]int)
	for i, v := range nums {
		numMap[v] = append(numMap[v], i)
	}

	for i, v := range nums {
		need := v
		if jlist, ok := numMap[need]; ok {
			for _, j := range jlist {
				if i < j {
					ans++
				}
			}
		}
	}
	return ans
}

func maxSum(nums []int) int {
	ans := -1
	numMap := make(map[int][]int) //max -> idx
	for _, v := range nums {
		mx := 0
		tp := v
		for ; v != 0; v /= 10 {
			x := v % 10
			if x > mx {
				mx = x
			}
		}
		numMap[mx] = append(numMap[mx], tp)
	}

	for _, list := range numMap {
		if len(list) < 2 {
			continue
		}
		sort.Ints(list)
		n := len(list)

		if t := list[n-1] + list[n-2]; t > ans {
			ans = t
		}
	}
	return ans

}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	n := len(nums)
	mx, mi := 0, 0
	for j := indexDifference; j < n; j++ {
		i := j - indexDifference
		if nums[i] > nums[mx] {
			mx = i
		}
		if nums[i] < nums[mi] {
			mi = i
		}
		if nums[j]-nums[mi] >= valueDifference {
			return []int{mi, j}
		}
		if nums[mx]-nums[j] >= valueDifference {
			return []int{mx, j}
		}
	}
	return []int{-1, -1}

}

// https://leetcode.cn/problems/identify-the-largest-outlier-in-an-array/
func getLargestOutlier(nums []int) int {

	// 对于每个数, 计算是否是其他 n-2个数的和, 如果是, 记录最大异常值
	// 异常值 = 前缀和 + 后缀和 - 当前数
	// 如果异常值存在,且有不在当前下标的,记录

	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	ans := math.MinInt
	for _, v := range nums {
		need := sum - v
		numMap[v]--
		if numMap[need] > 0 {
			ans = v
			break
		}
	}
	return ans
}

// https://leetcode.cn/problems/best-sightseeing-pair/description/
func maxScoreSightseeingPair(values []int) int {
	vi_plus_i := values[0]
	ans := math.MinInt
	// x = vj - j + (i + vi )
	for i := 1; i < len(values); i++ {
		x := values[i] - i + vi_plus_i
		if x > ans {
			ans = x
		}
		if values[i]+i > vi_plus_i {
			vi_plus_i = values[i] + i
		}
	}
	return ans
}

// https://leetcode.cn/problems/count-nice-pairs-in-an-array/description/
func countNicePairs(nums []int) int {
	rev := func(n int) int {
		ans := 0
		for n > 0 {
			ans = ans*10 + n%10
			n /= 10
		}
		return ans
	}
	// nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
	numMap := make(map[int]int)
	ans := int64(0)
	mod := int64(1e9 + 7)
	for _, v := range nums {
		need := v - rev(v)
		// fmt.Printf("need: %d, numMap[need]: %d\n", need, numMap[need])
		if numMap[need] > 0 {
			ans += int64(numMap[need])
			ans %= mod
		}
		numMap[need]++
	}
	return int(ans)
}

// https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays/
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	v1 := maxSumTwoNoOverlapHelper(nums, firstLen, secondLen)
	v2 := maxSumTwoNoOverlapHelper(nums, secondLen, firstLen)
	if v1 > v2 {
		return v1
	}
	return v2
}

func maxSumTwoNoOverlapHelper(nums []int, firstLen int, secondLen int) int {
	// fmt.Println("firstLen: ", firstLen, "secondLen: ", secondLen)
	// 分框
	// firstLen 和 secondLen 是固定的
	// 滑动 两个窗口 (firstLen, secondLen), 并维护 firstLen 窗口的最大值
	// 返回最大的和
	pre := 0
	for i := 0; i < firstLen; i++ {
		pre += nums[i]
	}
	// fmt.Println("pre: ", pre)
	preMax := pre
	sum := 0
	for i := 0; i < secondLen-1; i++ {
		sum += nums[i+firstLen]
	}

	ans := 0
	for i := firstLen + secondLen - 1; i < len(nums); i++ {
		sum += nums[i]
		if ans < sum+preMax {
			ans = sum + preMax
		}
		// fmt.Printf("i: %d, ans: %d, sum: %d, preMax: %d\n", i	, ans, sum, preMax)

		hh := i - firstLen - secondLen + 1
		pre += nums[hh+firstLen] - nums[hh]
		hh++
		if pre > preMax {
			preMax = pre
		}
		// fmt.Printf("hh: %d, pre: %d\n", hh, pre)
		sum -= nums[i-secondLen+1]
	}
	return ans
}

// 2555. 两个线段获得的最多奖品 https://leetcode.cn/problems/maximize-win-from-two-segments/
func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	leftMax := make([]int, n+1)
	left := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, v := range prizePositions {
		for v-prizePositions[left] > k {
			left++
		}
		leftMax[i+1] = max(leftMax[i], i-left+1)
	}
	right := 0
	ans := 0
	for i, v := range prizePositions {
		for v-prizePositions[right] > k {
			right++
		}
		ans = max(ans, i-right+1+leftMax[right])
	}
	return ans
}
func countQuadruplets(nums []int) int {
	n := len(nums)
	cnt := make(map[int]int)
	ans := 0
	for b := n - 3; b >= 0; b-- {
		for j := b + 2; j < n; j++ {
			cnt[nums[j]-nums[b+1]]++
		}
		for j := 0; j < b; j++ {
			ans += cnt[nums[b]+nums[j]]
		}
	}
	return ans

}

func numberOfSubsequences(nums []int) int64 {
	divMap := make(map[string]int64)

	ans := int64(0)
	n := len(nums)
	for q := n - 4; q >= 0; q-- {
		for r := q + 2; r < n; r++ {
			for s := r + 2; s < n; s++ {
				div := gcd(nums[r], nums[s])
				key := fmt.Sprintf("%d/%d", nums[s]/div, nums[r]/div)
				fmt.Printf("q: %d, r: %d, s: %d, key: %s\n", q, r, s, key)
				divMap[key]++
			}
		}
		for p := 0; p < q-1; p++ {
			div := gcd(nums[p], nums[q])
			key := fmt.Sprintf("%d/%d", nums[p]/div, nums[q]/div)
			ans += divMap[key]
			if divMap[key] > 0 {
				fmt.Printf("p: %d, q: %d, key: %s, divMap[key]: %d\n", p, q, key, divMap[key])
			}
		}
	}
	return ans
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func pairSums(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		switch {
		case sum == target:
			ans = append(ans, []int{nums[i], nums[j]})
			i++
		case sum < target:
			i++
		case sum > target:
			j--
		}
	}
	return ans
}

func countPairs(nums []int) int {
	sort.Ints(nums)
	ans := 0
	numMap := make(map[int]int)
	atoi := func(s string) int {
		num, _ := strconv.Atoi(s)
		return num
	}
	for i := 0; i < len(nums); i++ {
		set := map[int]bool{}
		set[nums[i]] = true
		bytes := []byte(strconv.Itoa(nums[i]))
		m := len(bytes)
		for j := 0; j < m; j++ {
			for k := j + 1; k < m; k++ {
				bytes[j], bytes[k] = bytes[k], bytes[j]
				set[atoi(string(bytes))] = true
				for z := 0; z < m; z++ {
					for x := z + 1; x < m; x++ {
						bytes[z], bytes[x] = bytes[x], bytes[z]
						set[atoi(string(bytes))] = true
						bytes[z], bytes[x] = bytes[x], bytes[z]
					}
				}
				bytes[j], bytes[k] = bytes[k], bytes[j]
			}
		}
		for s := range set {
			ans += numMap[s]
		}
		numMap[nums[i]]++
	}
	return ans
}



func findMaxK(nums []int) int {
	numMap := make(map[int]bool)
	ans := 0
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	for _, v := range nums {
		if numMap[-v] {
			ans = max(ans, abs(v))
		}
		numMap[v] = true
	}
	return ans

    
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
