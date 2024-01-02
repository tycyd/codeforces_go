package main

import (
	"bufio"
	"flag"
	. "fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

/*
p, 1, 1, 1, 1, 1, 1....1, 2
2*p > p + (n - 2) + 2
2*p > p + n

p must > n
*/
var MAX_P int64 = 1 << 30

func solve(n int, a []int64) []int {

	var p int64 = 1
	for _, v := range a {
		p *= v
		if p > MAX_P {
			l := 0
			r := n - 1
			for a[l] == 1 {
				l++
			}
			for a[r] == 1 {
				r--
			}

			return []int{l + 1, r + 1}
		}
	}

	prefix_a := make([][]int64, n+1)
	prefix_a[0] = []int64{0, 1}
	idx_a := make([]int, 0)
	for i, v := range a {
		prefix_a[i+1] = []int64{prefix_a[i][0] + v, prefix_a[i][1] * v}
		if v > 1 {
			idx_a = append(idx_a, i+1)
		}
	}

	lr := []int{1, 1}
	var max int64 = 0

	for i := 0; i < len(idx_a); i++ {
		for j := i; j < len(idx_a); j++ {
			l := idx_a[i]
			r := idx_a[j]
			var p int64 = prefix_a[r][1] / prefix_a[l-1][1]
			var sum int64 = prefix_a[l-1][0] + prefix_a[n][0] - prefix_a[r][0]

			if p+sum > max {
				max = p + sum
				lr[0] = l
				lr[1] = r
			}
		}
	}

	return lr
}

func CF1872G(_r io.Reader, out io.Writer) {
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

	// scan
	t := readInt(in)

	for t > 0 {
		n := readInt(in)
		a := readArrInt64(in)
		lr := solve(n, a)
		Fprintln(out, lr[0], lr[1])
		t--
	}
}

func main() { CF1872G(os.Stdin, os.Stdout) }

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
