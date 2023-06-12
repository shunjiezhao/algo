package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head.Next, head.Next.Next
	for ; slow != fast && slow != nil && fast != nil && fast.Next != nil; slow, fast = slow.Next, fast.Next.Next {
	}

	return slow == fast
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	u, d := l1, l2
	now := l1
	delta := 0
	for u != nil && d != nil {
		now.Val = u.Val + d.Val + delta
		delta = now.Val / 10
		now.Val %= 10
		u, d = u.Next, d.Next
		if u != nil {
			now = now.Next
		}
	}

	if u == nil && d == nil {
		if delta != 0 {
			now.Next = &ListNode{delta, nil}
		}
		return l1
	}

	if u == nil {
		now.Next = d
		now = now.Next
	}

	for delta != 0 {
		now.Val = now.Val + delta
		delta = now.Val / 10
		now.Val %= 10
		if now.Next == nil && delta != 0 {
			now.Next = &ListNode{}
		}
		now = now.Next
	}
	return l1
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	u, d := list1, list2
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var now, next *ListNode
	for u != nil && d != nil {
		s := d
		if u.Val <= d.Val {
			s = u
		}
		if now == nil {
			now, next = s, s
		} else {
			next.Next = s
			next = s
		}
		u, d = u.Next, d.Next
	}

	for u != nil {
		next.Next = u
		next = next.Next
		u = u.Next
	}

	for d != nil {
		next.Next = d
		next = next.Next
		d = d.Next
	}
	return now
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 0 or 1
		return head
	}

	now, next := head, head.Next
	for next != nil {
		now.Next = next.Next
		next.Next = head
		head = next
		next = now.Next
	}
	return head
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummp := &ListNode{Next: head}
	var lprev *ListNode
	l := dummp
	for i := 0; i < left; i++ {
		lprev, l = l, l.Next
	}
	if lprev == nil {
		lprev = l
	}

	for i := 0; i < right-left; i++ {
		next := l.Next
		l.Next = next.Next
		next.Next = lprev.Next
		lprev.Next = next
	}
	return head.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummp := &ListNode{}
	prev := dummp
	for now := head; now != nil; {
		next := now.Next
		for next != nil && now.Val == next.Val {
			next = next.Next
		}
		if now.Next == next {
			now.Next = nil
			prev.Next = now
			prev = prev.Next
		}
		now = next
	}
	return dummp.Next
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var lprev *ListNode
	r := head
	l := head
	for i := 0; i < n; i++ {
		r = r.Next
	}
	for r != nil {
		r = r.Next
		lprev = l
		l = l.Next
	}

	if lprev == nil {
		return l.Next
	}

	lprev.Next = l.Next
	return head
}

func partition(head *ListNode, x int) *ListNode {
	d1, d2 := &ListNode{}, &ListNode{}
	l, r := d1, d2

	for now := head; now != nil; {
		next := now.Next
		now.Next = nil
		if now.Val < x {
			l.Next = now
			l = l.Next
		} else {
			r.Next = now
			r = r.Next
		}
		now = next
	}
	l.Next = d2.Next
	return d1.Next
}

func rotateRight(head *ListNode, k int) *ListNode {
	h := head
	le := 0
	if h == nil {
		return nil
	}
	for h != nil {
		le++
		h = h.Next
	}
	k %= le
	if k == 0 {
		return head
	}

	h = head

	for i := 1; i < le-k; i++ {
		h = h.Next
	}
	dummp := &ListNode{}
	dummp.Next = head
	ans := dummp
	for h != nil && h.Next != nil {
		t := h.Next
		h.Next = t.Next
		t.Next = ans.Next
		ans.Next = t
		ans = t
	}
	return dummp.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummp := &ListNode{Next: head}
	l := dummp
	h := head
	for h != nil {
		t := h
		for i := 0; i < k-1; i++ { // get r]
			t = t.Next
			if t == nil {
				return dummp.Next
			}
		}
		if t != nil { // r)
			t = t.Next
		}

		hh := h
		for hh.Next != t {
			next := hh.Next
			hh.Next = next.Next
			next.Next = l.Next
			l.Next = next
		}
		l = hh // get (l
		h = t
	}
	return dummp.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	clone := map[*Node]*Node{}
	h := head
	for h != nil {
		_, ok := clone[h]
		if ok {
			continue
		}
		clone[h] = &Node{
			Val: h.Val,
		}
		h = h.Next
	}

	dummp := &Node{}
	ans := dummp
	h = head
	for h != nil {
		next := h.Next
		t := clone[h]
		ans.Next = t
		t.Next = clone[h.Next]
		t.Random = clone[h.Random]
		h = next
		ans = t
	}
	return dummp.Next

}


type LRUCache struct {
	li     *dlinkedList
	val2li map[int]*dlinkedNode
	cap    int
	size int
}

type dlinkedList struct {
	head, tail *dlinkedNode
}

func (node *dlinkedNode) clear() {
	node.next, node.prev = nil, nil
}

func newlinkedList() *dlinkedList {
	r := dlinkedList{}
	r.head = &dlinkedNode{}
	r.tail = &dlinkedNode{}
	r.head.next = r.tail
	r.tail.prev = r.head
	return &r
}

func (l *dlinkedList) RemoveNode(li *dlinkedNode) {
	if li == l.head || li == l.tail {
		panic("don't remove head or tail")
	}
	li.next.prev = li.prev
	li.prev.next = li.next
}

func (l *dlinkedList) addNext(prev, node *dlinkedNode) {
	node.next = prev.next
	node.prev = prev
	prev.next.prev = node
	prev.next = node
}

func (l *dlinkedList) PushFront(node *dlinkedNode) {
	l.addNext(l.head, node)
}

func (l *dlinkedList) RemoveAndPushFront(node *dlinkedNode) {
	l.RemoveNode(node)
	l.PushFront(node)
}
func (l *dlinkedList) String() {
	str := &strings.Builder{}
	for h := l.head.next; h != l.tail; h = h.next {
		str.WriteString(strconv.FormatInt(int64(h.key), 10))
	}
	fmt.Println(str.String())
}


func (l *dlinkedList) Back() *dlinkedNode {
	if l.head.next == l.tail {
		return nil
	}

	return l.tail.prev
}

type dlinkedNode struct {
	key,value      int
	next, prev *dlinkedNode
}

func Constructor(capacity int) LRUCache {
	r := LRUCache{}
	r.li = newlinkedList()
	r.val2li = make(map[int]*dlinkedNode)
	r.cap = capacity
	return r
}

func (this *LRUCache) containKey(key int) bool {
	_, ok := this.val2li[key]
	return ok
}
func (this *LRUCache) Get(key int) int {
	if this.containKey(key) == false {
		return -1
	}
	node := this.val2li[key]
	fmt.Println("get ", key)
	if node.key != key{
		panic("")
	}
	// contain
	this.li.RemoveAndPushFront(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	fmt.Println(this.size, this.cap, key, value)
	var li *dlinkedNode
	this.li.String()
	if this.containKey(key) == false {
		this.size ++
		if this.size > this.cap {
			back := this.li.Back()
			this.li.RemoveNode(back)        // remove from linkedlist
			delete(this.val2li, back.key) // delete map val2li
			this.size--
		}

		li = &dlinkedNode{}
		this.val2li[key] = li
	} else {
		li = this.val2li[key]
		this.li.RemoveNode(li)
	}
	fmt.Println(this.val2li)
	li.value = value
	this.li.PushFront(li)
	this.li.String()
	fmt.Println("-----------")
}