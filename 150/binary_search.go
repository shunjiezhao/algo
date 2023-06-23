package main

import "sort"

func searchInsert(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
}

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	lines := make([]int, n)
	for i := range matrix {
		lines[i] = matrix[i][0]
	}

	t := sort.Search(n, func(i int) bool {
		return matrix[i][0] > target
	}) - 1
	if t < 0 {
		return false
	}

	l := sort.Search(m, func(i int) bool {
		return matrix[t][i] >= target
	})
	if l == m || matrix[t][l] != target {
		return false
	}

	return true
}

func findPeakElement(nums []int) int {
	return sort.Search(len(nums), func(i int) bool {
		if i == 0 ||
			nums[i] >= nums[i-1] {
			return false
		}
		return true
	}) - 1
}

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		mid := int((l + r + 1) / 2) // last true
		if nums[0] <= nums[mid] {   // up |
			if nums[mid] > target && nums[0] <= target { // move left
				r = mid - 1
			} else {
				l = mid
			}
		} else { // nums[0] > nums[mid]
			if nums[mid] <= target && nums[mid] <= nums[n-1] {
				l = mid
			} else {
				r = mid - 1
			}
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	t := sort.Search(n, func(i int) bool {
		return nums[i] >= target
	})
	if t == n || nums[t] != target {
		return []int{-1, -1}
	}
	r := sort.Search(n, func(i int) bool {
		return nums[i] > target
	})
	return []int{t, r - 1}
}

func findMin(nums []int) int {

	n := len(nums)
	l, r := 0, n-1
	for l < r {
		mid := int((l + r) / 2)
		if nums[mid] < nums[0] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l == n && nums[l] > nums[0] { //  1 , 2, 3 ,4
		return nums[0]
	}
	return nums[l]
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	min := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	find := func(k int) int {
		up, down := 0, 0
		for k != 0 {
			if up == n {
				return nums2[down+k]
			}
			if down == m {
				return nums1[up+k]
			}
			// k == 1 k / 2 == 0
			if k == 1 {
				return min(nums1[up+1], nums2[down+1])
			}

			t := int(k / 2)
			if nums1[up+t] <= nums2[down+t] {
				up += t
			} else {
				down += t
			}
			k -= t
		}
		return -1
	}
	// jishu
	le := n + m
	ans := find(int((le-1)/2)) + find(int(le/2))

	return float64(ans / 2)

}
