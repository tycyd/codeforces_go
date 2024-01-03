package main

import (
	"fmt"
)

// static bottom -> top
type SegmentNode struct {
	val int
	L   int
	R   int
	F   int
}

func NewSegmentNode(val int, l int, r int) *SegmentNode {
	node := SegmentNode{}
	node.val = val
	node.L = val
	node.R = val
	node.F = val
	return &node
}

type SegmentTree struct {
	nodes []*SegmentNode
	size  int
}

func (t *SegmentTree) build(n int, a []int) {
	t.size = 1
	for t.size < n {
		t.size <<= 1
	}

	t.nodes = make([]*SegmentNode, t.size*2)
	for i := 0; i < n; i++ {
		t.nodes[t.size+i] = NewSegmentNode(a[i], i, i)
	}
	for i := t.size - 1; i > 0; i-- {
		t.nodes[i] = NewSegmentNode(-1, i<<1, i<<1+1)
		t.merge(t.nodes[i], t.nodes[i<<1], t.nodes[i<<1+1])
	}
}

func (t *SegmentTree) modify(i int, v int) {
	t.nodes[t.size+i].val = v
	t.nodes[t.size+i].L = v
	t.nodes[t.size+i].F = v
	t.nodes[t.size+i].R = v

	for i := (t.size + i) >> 1; i > 0; i >>= 1 {
		t.merge(t.nodes[i], t.nodes[i<<1], t.nodes[i<<1+1])
	}
}

func (t *SegmentTree) merge(n *SegmentNode, ln *SegmentNode, rn *SegmentNode) {
	if ln == nil && rn == nil {
		return
	}
	if ln == nil {
		n.val = rn.val
		n.L = rn.L
		n.R = rn.R
		n.F = rn.F
		return
	}
	if rn == nil {
		n.val = ln.val
		n.L = ln.L
		n.R = ln.R
		n.F = ln.F
		return
	}
	//  (F)    (F)
	// (L R), (L R)
	n.L = max(ln.L, ln.F+rn.L, ln.F+rn.F)
	n.R = max(rn.R, rn.F+ln.R, ln.F+rn.F)
	n.F = ln.F + rn.F
	n.val = max(n.L, n.R, ln.val, rn.val, ln.R+rn.L)
}

func print(t *SegmentTree) {
	fmt.Println("----st-----")
	for i := t.size; i < 2*t.size; i++ {
		if t.nodes[i] == nil {
			break
		}
		fmt.Print(t.nodes[i].val)
		fmt.Print(" ")
	}
	fmt.Println()
}

func max(a ...int) int {
	r := a[0]
	for _, v := range a {
		if v > r {
			r = v
		}
	}
	return r
}
