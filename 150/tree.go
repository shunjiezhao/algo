package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maintain depth when dfs
func maxDepth(root *TreeNode) int {
	var ans int
	var dfs func(root *TreeNode, depth int)

	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			if ans < depth {
				ans = depth
			}
			return
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}

	dfs(root, 0)
	return ans
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// dfs maintain tht root
	// if root != root return false
	// then check left and right

	var dfs func(l, r *TreeNode) bool
	dfs = func(l, r *TreeNode) bool {
		if l == nil || r == nil {
			return l == r
		}
		if l.Val != r.Val {
			return false
		}
		return dfs(l.Left, r.Left) && dfs(l.Right, r.Right)
	}

	return dfs(p, q)
}

func invertTree(root *TreeNode) *TreeNode {
	// we don't change root
	// when we think the root
	// the left and right if changed
	// and we just swap them

	var dfs func(root *TreeNode) *TreeNode

	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		root.Right, root.Left = dfs(root.Left), dfs(root.Right)
		return root
	}

	dfs(root)
	return root
}

func isSymmetric(root *TreeNode) bool {
	// maintain the two root node
	// if left.val == r.val then true
	// dfs this child
	// compare (left.r, right.l) (left.l, right.r) this subtree
	if root == nil {
		return true
	}

	var dfs func(l, r *TreeNode) bool
	dfs = func(l, r *TreeNode) bool {
		if l == nil || r == nil {
			return l == r
		}

		if l.Val != r.Val {
			return false
		}

		return dfs(l.Left, r.Right) && dfs(l.Right, r.Left)
	}
	return dfs(root.Left, root.Right)
}

/*
	 now we should think about what is preoder, inorder, suforder

		when  tree like
				a
			b		c
	root is say to root
	 preorder: a b c
	 inorder: b a c
	 suforder: b c a
*/
// buildTree
func buildTree(preorder []int, inorder []int) *TreeNode {

	// accord to preorder , we can know root
	// 			 inorder, we can know left and right scope

	findIndexByVal := func(val int) int {
		for i, v := range inorder {
			if v == val {
				return i
			}
		}
		return -1
	}

	var dfs func(l, r int) *TreeNode
	index := 0
	dfs = func(l, r int) *TreeNode {
		if r < l {
			return nil
		}
		val := preorder[index]
		root := &TreeNode{Val: val}
		index++
		i := findIndexByVal(val)
		root.Left, root.Right = dfs(l, i-1), dfs(i+1, r)
		return root
	}
	return dfs(0, len(inorder)-1)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	// accord to suforder , we can know root
	// 			 inorder, we can know left and right scope

	findIndexByVal := func(val int) int {
		for i, v := range inorder {
			if v == val {
				return i
			}
		}
		return -1
	}

	var dfs func(l, r int) *TreeNode
	index := len(inorder) - 1
	dfs = func(l, r int) *TreeNode {
		if r < l {
			return nil
		}
		val := postorder[index]
		root := &TreeNode{Val: val}
		index--
		i := findIndexByVal(val)
		root.Right = dfs(i+1, r)
		root.Left = dfs(l, i-1)
		return root
	}
	return dfs(0, len(inorder)-1)
}

func connect(root *Node) *Node {

	var dfs func(root *Node) *Node
	var find func(root *Node) *Node
	find = func(root *Node) *Node {
		for root.Next != nil {
			if root.Next.Left != nil {
				return root.Next.Left
			}

			if root.Next.Right != nil {
				return root.Next.Right
			}

			root = root.Next
		}
		return nil
	}

	dfs = func(root *Node) *Node {
		if root == nil || (root.Left == nil && root.Right == nil) {
			return root
		}

		if root.Left != nil && root.Right != nil {
			root.Left.Next = root.Right
			root.Right.Next = find(root)
		}

		if root.Left == nil {
			root.Right.Next = find(root)
		}

		if root.Right == nil {
			root.Left.Next = find(root)
		}

		root.Right = dfs(root.Right)
		root.Left = dfs(root.Left)
	}

}

func flatten(root *TreeNode) {
	var dfs func(root *TreeNode) *TreeNode

	dfs = func(root *TreeNode) *TreeNode {
		if root == nil || (root.Left == nil && root.Right == nil) {
			return root
		}
		l, r := dfs(root.Left), dfs(root.Right)
		root.Left, root.Right = nil, nil

		if l != nil {
			root.Right = l
			for l.Right != nil {
				l = l.Right
			}
			l.Right = r
		}
		if l == nil {
			root.Right = r
		}

		return root
	}
	dfs(root)

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var dfs func(r *TreeNode, sum int) bool

	dfs = func(r *TreeNode, sum int) bool {
		if sum == targetSum && (r.Left == nil && r.Right == nil) {
			return true
		}
		if r == nil {
			return false
		}
		if r.Left != nil && dfs(r.Left, sum+r.Left.Val) {
			return true
		}

		if r.Right != nil && dfs(r.Right, sum+r.Right.Val) {
			return true
		}
		return false
	}
	return dfs(root, root.Val)
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0

	var dfs func(r *TreeNode, sum int) bool

	dfs = func(r *TreeNode, sum int) bool {
		if r == nil {
			return false
		}
		sum = sum*10 + r.Val
		if r.Left == nil && r.Right == nil {
			ans += sum
			return true
		}
		if r.Left != nil && dfs(r.Left, sum) {
			return true
		}

		if r.Right != nil && dfs(r.Right, sum) {
			return true
		}
		return false
	}
	dfs(root, 0)
	return ans
}

func maxPathSum(root *TreeNode) int {
	var dfs func(*TreeNode) int
	var ans int = math.MinInt64

	dfs = func(cur *TreeNode) int {
		if cur == nil {
			return 0
		}

		l := dfs(cur.Left)
		r := dfs(cur.Right)
		dist := max(cur.Val, max(cur.Val+l, cur.Val+r))
		ans = max(ans, max(dist, r+l+cur.Val))
		return dist
	}
	dfs(root)
	return ans
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var dfs func(r *TreeNode) *TreeNode

	dfs = func(r *TreeNode) *TreeNode {
		if r == nil {
			return r
		}
		if p == r || q == r {
			return r
		}

		l := dfs(r.Left)
		rt := dfs(r.Right)
		if l == nil {
			return rt
		}
		if rt == nil {
			return l
		}
		return r
	}
	return dfs(root)

}

func countNodes(root *TreeNode) int {
	h := 0
	t := root
	for t != nil {
		h++
		t = t.Left
	}

	var dfs func(*TreeNode, int) bool
	ans := (1 << h) - 1

	dfs = func(t *TreeNode, depth int) bool {
		if depth == h {
			if t == nil {
				ans--
				return false
			}
			return true
		}

		if dfs(t.Right) || dfs(t.Left) {
			return true
		}
		return false

	}
	dfs(root, 1)
	return ans

}

type BSTIterator struct {
	stk []*TreeNode
	cur *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	t := BSTIterator{}
	t.cur = root
	return t
}

func (this *BSTIterator) Next() int {
	for this.cur != nil {
		this.stk = append(this.stk, this.cur)
		this.cur = this.cur.Left
	}

	this.cur = this.stk[len(this.stk)-1]
	this.stk = this.stk[:len(this.stk)-1]
	ret := this.cur.Val
	this.cur = this.cur.Right
	return ret
}

func (this *BSTIterator) HasNext() bool {
	return !(this.cur == nil && len(this.stk) == 0)

}
