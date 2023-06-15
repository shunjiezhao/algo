package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	n := len(nums)
	var down func(l, r int) *TreeNode
	down = func(l, r int) *TreeNode {
		if r <= l {
			return nil
		}

		mid := (l + r) / 2
		left, right := down(l, mid), down(mid+1, r)
		root := &TreeNode{Val: nums[mid]}
		root.Left = left
		root.Right = right
		return root
	}
	return down(0, n)
}

func sortedListToBST(head *ListNode) *TreeNode {
	le := 0
	h := head
	for h != nil {
		h = h.Next
		le++
	}

	cursor := head
	var down func(l, r int) *TreeNode
	down = func(l, r int) *TreeNode {
		if r <= l {
			return nil
		}
		mid := (l + r) / 2
		root := &TreeNode{}
		root.Left = down(l, mid)
		root.Val = cursor.Val
		cursor = cursor.Next
		root.Right = down(mid+1, r)
		return root
	}
	return down(0, le)
}

func sortList(head *ListNode) *ListNode {
	le := 0
	h := head
	for h != nil {
		h = h.Next
		le++
	}

	cursor := head
	var down func(l, r int) *ListNode
	down = func(l, r int) *ListNode {
		if r < l {
			return nil
		}
		if r == l {
			p := cursor
			cursor = cursor.Next
			p.Next = nil
			return p
		}

		m := (l + r) / 2
		left, right := down(l, m), down(m+1, r)
		dummp := &ListNode{}
		head := dummp

		for left != nil && right != nil {
			if left.Val <= right.Val {
				head.Next = left
				left = left.Next
			} else {
				head.Next = right
				right = right.Next
			}
			head = head.Next
		}

		for ; left != nil; left, head = left.Next, head.Next {
			head.Next = left
		}
		for ; right != nil; right, head = right.Next, head.Next {
			head.Next = right
		}
		return dummp.Next
	}
	return down(0, le-1)

}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {

}

func mergeKLists(lists []*ListNode) *ListNode {

	var down func(l, r int) *ListNode
	down = func(l, r int) *ListNode {
		if r < l {
			return nil
		}
		if r == l {
			return lists[l]
		}

		m := (l + r) / 2
		left, right := down(l, m), down(m+1, r)
		dummp := &ListNode{}
		head := dummp
		for left != nil && right != nil {
			if left.Val <= right.Val {
				head.Next = left
				left = left.Next
			} else {
				head.Next = right
				right = right.Next
			}
			head = head.Next
		}

		for ; left != nil; left, head = left.Next, head.Next {
			head.Next = left
		}
		for ; right != nil; right, head = right.Next, head.Next {
			head.Next = right
		}
		return dummp.Next
	}
	return down(0, len(lists)-1)
}

func construct(grid [][]int) *Node {
	calc := func(ll, lr, rl, rr int) bool {
		for i := ll; i <= rl; i++ {
			for j := lr; j <= rr; j++ {
				if grid[i][j] != grid[ll][lr] {
					return false
				}
			}
		}
		return true
	}

	var down func(ll, lr, rl, rr int) *Node
	down = func(ll, lr, rl, rr int) *Node {
		if calc(ll, lr, rl, rr) {
			ans := &Node{}
			if grid[ll][lr] == 1 {
				ans.Val = true
			}
			ans.IsLeaf = true
			return ans
		}

		ans := &Node{}
		mr, mc := (ll+rl)/2, (lr+rr)/2
		ans.TopLeft = down(ll, lr, mr, mc)
		ans.TopRight = down(ll, mc+1, mr, rr)
		ans.BottomLeft = down(mr + 1, lr, rl, mc)
		ans.BottomRight = down(mr + 1, mc + 1, rl, rr)
		return ans
	}
	n := len(grid)
	m := len(grid[0])
	return down(0, 0,n-1,m-1)

}
