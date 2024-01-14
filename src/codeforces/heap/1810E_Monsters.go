package main

import (
	"bufio"
	"container/heap"
	"flag"
	. "fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

// heap
// An IntHeap is a min-heap of ints.
type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func solve(n int, a []int, mp map[int][]int) bool {

	visited := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if visited[i] {
			continue
		}
		if a[i-1] == 0 {
			if bfs(i, n, a, mp, visited) {
				return true
			}
		}
	}

	return false
}

func bfs(i int, n int, a []int, mp map[int][]int, visited []bool) bool {
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, []int{a[i-1], i})

	v2 := map[int]bool{}

	d := 0
	for h.Len() > 0 {
		cur := heap.Pop(h).([]int)
		if _, ok := v2[cur[1]]; ok {
			continue
		}
		if d < cur[0] {
			return false
		}
		d++
		if d == n {
			return true
		}

		visited[cur[1]] = true
		v2[cur[1]] = true
		for _, nxt := range mp[cur[1]] {
			if _, ok := v2[nxt]; ok {
				continue
			}
			heap.Push(h, []int{a[nxt-1], nxt})
		}
	}

	return false
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func CF1810E(_r io.Reader, out io.Writer) {
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
		nm := readArrInt(in)
		a := readArrInt(in)
		mp := map[int][]int{}
		for i := 0; i < nm[1]; i++ {
			uv := readArrInt(in)
			if _, ok := mp[uv[0]]; !ok {
				mp[uv[0]] = make([]int, 0)
			}
			if _, ok := mp[uv[1]]; !ok {
				mp[uv[1]] = make([]int, 0)
			}
			mp[uv[0]] = append(mp[uv[0]], uv[1])
			mp[uv[1]] = append(mp[uv[1]], uv[0])
		}
		if solve(nm[0], a, mp) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
		t--
	}
}

func main() { CF1810E(os.Stdin, os.Stdout) }

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
