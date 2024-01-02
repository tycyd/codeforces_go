package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

/*
	1 (0)

2 1 3

	1 (3)

1 1

		   1 (3)
	    	 2 (0)

2 2 1

		   1 (3)
	         2 (1)

1 1

	   1 (3)
	2 (1) 3 (0)

2 3 2

		   1 (3)
	    2 (1) 3 (2)

1 3

		   1 (3)
	    2 (1) 3 (2)
	            4 (0)

2 1 4

		   1 (7)
	    2 (5) 3 (6)
	            4 (4)

1 3

		   1 (7)
	    2 (5)   3 (6)
	          5(0)  4 (4)

2 3 2

		   1 (7)
	    2 (5)   3 (8)
	          5(2)  4 (6)
*/
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	in := bufio.NewReader(os.Stdin)
	out := os.Stdout

	T := readInt(in)
	for T > 0 {
		mp := map[int][]int{}
		q := readInt(in)

		size := 1
		tvxa := make([][]int, q)
		for i := 0; i < q; i++ {
			tvxa[i] = readArrInt(in)

			if tvxa[i][0] == 1 {
				size++
				_, ok := mp[tvxa[i][1]]
				if !ok {
					mp[tvxa[i][1]] = make([]int, 0)
				}
				mp[tvxa[i][1]] = append(mp[tvxa[i][1]], size)
			}
		}

		et := make([]Node, size+1)
		Time = 0
		dfs(1, mp, et)

		br := BIT_Range{}
		br.build(size + 1)

		node := 2
		for _, tvx := range tvxa {
			if tvx[0] == 1 {
				sum := br.range_sum(et[node].start, et[node].start)
				br.range_add(et[node].start, et[node].start, -sum)
				node++
			} else {
				br.range_add(et[tvx[1]].start, et[tvx[1]].end, int64(tvx[2]))
			}
		}

		for i := 1; i <= size; i++ {
			v := br.range_sum(et[i].start, et[i].start)
			fmt.Fprint(out, v, " ")
		}
		fmt.Fprintln(out)

		T--
	}
}

// Euler Tour
var Time int

func dfs(i int, mp map[int][]int, et []Node) {
	Time++
	et[i].start = Time
	for _, next := range mp[i] {
		dfs(next, mp, et)
	}
	et[i].end = Time
}

type Node struct {
	start int
	end   int
}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
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

/** binary indexed tree with range update and query **/
/*
https://cp-algorithms.com/data_structures/fenwick.html
sum[0,i] =
i<l			0
l<=i<=r	 	x*(i-(l-1))
i>r			x*(r-l+1)

sum[0,i] = sum(B1,i)*i - sum(B2,i)
i<l			0*i - 0
l<=i<=r	 	x*i - x*(l-1)
i>r			0*i - (x*(l-1) - x*r)

def range_add(l, r, x):
    add(B1, l, x)
    add(B1, r+1, -x)
    add(B2, l, x*(l-1))
    add(B2, r+1, -x*r))
*/

type BIT_Range struct {
	B1   []int64
	B2   []int64
	size int
}

func (t *BIT_Range) build(sz int) {
	t.size = sz + 2
	t.B1 = make([]int64, t.size)
	t.B2 = make([]int64, t.size)
}

func (t *BIT_Range) add(b []int64, i int, x int64) {
	i++
	for i < t.size {
		b[i] += x
		i += i & -i
	}
}

func (t *BIT_Range) range_add(l int, r int, x int64) {
	t.add(t.B1, l, x)
	t.add(t.B1, r+1, -x)
	t.add(t.B2, l, x*int64(l-1))
	t.add(t.B2, r+1, -x*int64(r))
}

func (t *BIT_Range) sum(b []int64, i int) int64 {
	i++
	var res int64 = 0
	for i > 0 {
		res += b[i]
		i -= i & -i
	}
	return res
}

func (t *BIT_Range) prefix_sum(i int) int64 {
	return t.sum(t.B1, i)*int64(i) - t.sum(t.B2, i)
}

func (t *BIT_Range) range_sum(l int, r int) int64 {
	return t.prefix_sum(r) - t.prefix_sum(l-1)
}
