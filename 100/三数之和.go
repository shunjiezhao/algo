package _00

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for i, v := range nums {
		if i > 0 && nums[i-1] == v {
			continue
		}
		target := -v
		l, r := i+1, n-1
		for l < r {
			sum := nums[l] + nums[r]
			if sum == target {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
			}
			if sum >= target { // move big
				for rval := nums[r]; l < r && nums[r] == rval; r-- {
				}
			} else {
				for lval := nums[l]; l < r && nums[l] == lval; l++ {
				}
			}
		}
	}
	return ans
}
