package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	nn := len(nums1)
	n--
	m--
	for i := nn - 1; i >= 0; i-- {
		if n <= 0 {
			break
		}
		if m < 0 || nums2[n] > nums1[m] {
			nums1[i] = nums2[n]
			n--
		} else {
			nums1[i] = nums1[m]
			m--
		}
	}
}

func removeElement(nums []int, val int) int {
	l, r := 0, len(nums)
	for l < r {
		if nums[l] == val {
			nums[l], nums[r] = nums[r-1], nums[l]
			r--
		} else {
			l++
		}
	}
	return l
}

func removeDuplicates1(nums []int) int {
	return removeDuplicates(nums, 1)
}

func removeDuplicates(nums []int, allowDupCnt int) int {
	n := len(nums)
	l, r := 0, 0
	for ; r < n; r++ {
		t := r
		for r < n && nums[r] == nums[t] {
			r++
		}
		for k := 0; k < min(r-t, allowDupCnt); k++ {
			nums[l] = nums[t]
			l++
		}
		r--
	}
	return l
}

func majorityElement(nums []int) int {
	var down func(l, r int) int
	down = func(l, r int) int {
		if l >= r {
			return nums[l]
		}

		mid := (l + r) / 2
		left, right := down(l, mid), down(mid+1, r)
		if left == right {
			return left
		}
		cntInCloseRange := func(l, r, num int) (count int) {
			for i := l; i <= r; i++ {
				if nums[i] == num {
					count++
				}
			}
			return
		}
		leftCount, rightCount := cntInCloseRange(l, r, left), cntInCloseRange(l, r, right)
		if leftCount > rightCount {
			return left
		}
		return right
	}
	return down(0, len(nums)-1)
}

func rotate(nums []int, k int) {
	reverse := func(l, r int) {
		for i := 0; i < (r-l+1)/2; i++ {
			nums[i+l], nums[r-i] = nums[r-i], nums[i+l]
		}
	}
	k %= len(nums)
	if k == 0 {
		return
	}
	reverse(0, len(nums)-1)
	reverse(k, len(nums)-1)
	reverse(0, k-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
