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

/*
L <= k < R
max{aL,...ak} < min{ak+1...aR}
ans = R - L - number of k

iterate i and assume min{ak+1...aR}=ai,find triple L,k,R

left range:
ak < ai and is the cloest one
if x < k and ax > ai and is the cloest one, then L = (x, k]

right range:
if y > i and ay < ai and also is the closest one, then R = [i, y)

(x+1 ... k) (k + 1... i ... y-1)
min(k + 1... i ... y-1) = ai, y-i ways
max(x+1 ... k) < ai, k-x ways

contribution: -(k-x)*(y-i)
*/
func solve(n int, a []int) int64 {
	var ans int64 = 0

	// all R-L
	for l := 0; l < n; l++ {
		for r := l + 1; r < n; r++ {
			ans += int64(r - l)
		}
	}

	for i := 1; i < n; i++ {
		// left range
		k := i - 1
		for k > 0 && a[k] > a[i] {
			k--
		}
		x := k
		for x >= 0 && a[x] < a[i] {
			x--
		}
		// right range
		y := i
		for y < n && a[y] >= a[i] {
			y++
		}

		if a[k] > a[i] {
			continue
		}

		ans -= int64((k - x) * (y - i))
	}

	return ans
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func CF1827B1(_r io.Reader, out io.Writer) {
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
		a := readArrInt(in)
		ans := solve(n, a)
		Fprintln(out, ans)
		t--
	}
}

func main() { CF1827B1(os.Stdin, os.Stdout) }

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
