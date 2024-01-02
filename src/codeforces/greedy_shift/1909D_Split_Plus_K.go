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
shifted problem:
A	 B
x    x - k

A:
y + z = x + k
y - k + z - k = x - k

B:
(y - k) + (z - k) = (x - k)
y' + z' = x'

assume result is
m', m', m'....m'
combine any get N*m'
m'=gcd(all numbers)

if all x' are positive, m'=gcd(all x') => means m must be bigger than k. ex x=12, k=10, if m exists, it must > k
if all x' are 0, m'=0
if all x' are negative, m'=-gcd(all x') => means m must be less than k. ex x=8, k=10, if m exists, it must < k
else -1
*/
func solve(n int, k int, a []int) int {

	pcnt := 0
	ncnt := 0
	zcnt := 0

	g := 0
	for i := 0; i < n; i++ {
		a[i] -= k

		if a[i] > 0 {
			g = gcd(a[i], g)
			pcnt++
		} else if a[i] < 0 {
			g = gcd(-a[i], g)
			ncnt++
		} else {
			zcnt++
		}
	}

	ans := -1
	if pcnt == n || ncnt == n {
		ans = 0
		for _, v := range a {
			ans += abs(v/g) - 1
		}
	} else if zcnt == n {
		ans = 0
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func CF1909D(_r io.Reader, out io.Writer) {
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
		nk := readArrInt(in)
		a := readArrInt(in)
		ans := solve(nk[0], nk[1], a)
		Fprintln(out, ans)
		t--
	}
}

func main() { CF1909D(os.Stdin, os.Stdout) }

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
