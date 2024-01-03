package main

import (
	"bufio"
	"flag"
	"fmt"
	. "fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

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

var INF int = -1000000000000000000
var INVALID int = -1000000000000000001

func solve_dp(n int, k int, x int, a []int) int {
	memo := make([][][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			memo[i][j] = make([]int, 3)
			memo[i][j][0] = INF // before result
			memo[i][j][1] = INF // in result
			memo[i][j][2] = INF // after result
		}
	}

	dfs(0, k, 0, x, a, memo)
	if memo[0][k][0] < 0 {
		return 0
	}
	return memo[0][k][0]
}

func dfs(i int, k int, t int, x int, a []int, memo [][][]int) int {
	if i == len(a) {
		if k == 0 {
			if t == 0 {
				return INVALID
			} else if t == 1 || t == 2 {
				return 0
			}
		} else {
			return INVALID
		}
	}
	if k < 0 {
		return INVALID
	}

	if memo[i][k][t] != INF {
		return memo[i][k][t]
	}

	var r1 int = INVALID
	var r2 int = INVALID
	var r3 int = INVALID
	var r4 int = INVALID
	var r5 int = INVALID
	var r6 int = INVALID
	var r7 int = INVALID
	var r8 int = INVALID
	var r9 int = INVALID
	var r10 int = INVALID
	if t == 0 {
		// before -> before
		r1 = dfs(i+1, k, 0, x, a, memo)
		r2 = dfs(i+1, k-1, 0, x, a, memo)

		// before -> in
		r3 = dfs(i+1, k, 1, x, a, memo)
		if r3 != INVALID {
			r3 += a[i] - x
		}
		r4 = dfs(i+1, k-1, 1, x, a, memo)
		if r4 != INVALID {
			r4 += a[i] + x
		}
	} else if t == 1 {
		// in -> in
		r5 = dfs(i+1, k, 1, x, a, memo)
		if r5 != INVALID {
			r5 += a[i] - x
		}
		r6 = dfs(i+1, k-1, 1, x, a, memo)
		if r6 != INVALID {
			r6 += a[i] + x
		}
		// in -> after
		r7 = dfs(i+1, k, 2, x, a, memo)
		r8 = dfs(i+1, k-1, 2, x, a, memo)

	} else {
		// after -> after
		r9 = dfs(i+1, k, 2, x, a, memo)
		r10 = dfs(i+1, k-1, 2, x, a, memo)
	}
	memo[i][k][t] = max(r1, r2, r3, r4, r5, r6, r7, r8, r9, r10)
	return memo[i][k][t]
}

func solve(n int, k int, x int, a []int) int {

	for i := 0; i < n; i++ {
		if x > 0 {
			a[i] -= x
		} else {
			a[i] += x
		}
	}

	st := SegmentTree{}
	st.build(n, a)

	ans := 0
	if x > 0 {
		for i := 0; i <= n-k; i++ {
			for j := i; j < i+k; j++ {
				st.modify(j, a[j]+2*x)
			}
			ans = max(ans, st.nodes[1].val)
			for j := i; j < i+k; j++ {
				st.modify(j, a[j])
			}
		}
	} else {
		for i := 0; i <= k; i++ {
			for j := i; j < i+n-k; j++ {
				st.modify(j, a[j]-2*x)
			}
			ans = max(ans, st.nodes[1].val)
			for j := i; j < i+n-k; j++ {
				st.modify(j, a[j])
			}
		}
	}

	return ans
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func CF1796D(_r io.Reader, out io.Writer) {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	in := bufio.NewReader(_r)

	t := readInt(in)

	for t > 0 {
		nkx := readArrInt(in)
		a := readArrInt(in)
		ans := solve_dp(nkx[0], nkx[1], nkx[2], a)
		Fprintln(out, ans)
		t--
	}
}

func main() { CF1796D(os.Stdin, os.Stdout) }

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func readLine(in *bufio.Reader) string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	return line
}

func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

func readArrInt(in *bufio.Reader) []int {
	numbs := readLineNumbs(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
}

func readArrInt64(in *bufio.Reader) []int64 {
	numbs := readLineNumbs(in)
	arr := make([]int64, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.ParseInt(n, 10, 64)
		arr[i] = val
	}
	return arr
}
