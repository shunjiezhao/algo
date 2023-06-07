package main

import (
	"sort"
	"strings"
)

func isPalindrome(s string) bool {
	var tmp []byte
	s = strings.ToLower(s)
	for _, c := range s {
		cs := string(c)
		if (cs >= "a" && cs <= "z") ||
			(cs >= "0" && cs <= "9") {
			tmp = append(tmp, byte(c))
		}
	}
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		if tmp[i] != tmp[j] {
			return false
		}
	}
	return true
}

func isSubsequence(s string, t string) bool {
	index := 0
	for i := 0; i < len(t); i++ {
		if index < len(s) && s[index] == t[i] {
			index++
		}
		if index == len(s) {
			return true
		}
	}
	return index == len(s)
}

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{-1, -1}
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int

	for k := 0; k < len(nums); k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		v := -nums[k]
		for i, j := k+1, len(nums)-1; i < j; {
			sum := nums[i] + nums[j]
			if sum == v {
				ans = append(ans, []int{v, nums[i], nums[j]})
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				j--
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				i++
				continue
			}
			if sum > v {
				j--
			} else {
				i++
			}
		}
	}
	return ans
}

func maxArea(height []int) int {
	var ans int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		return -max(-a, -b)
	}

	for i, j := 0, len(height)-1; i < j; {
		ans = max(ans, min(height[i], height[j])*(j-i))
		if height[i] < height[j]{
			i++
		}else{
			j--
		}
	}
	return ans
}
