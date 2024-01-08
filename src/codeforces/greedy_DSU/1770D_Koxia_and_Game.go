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

var MOD int = 998244353

/*
invalid:
1 --- 2
|     |
3 --- 4
|     |
5 --- 6

invalid:
1 --- 2
|     |
3 --- 4
|
5 --- 6
|     |
7 --- 8

valid:
1 --- 2
|     |
3 --- 4
|
5
*/
func ufind(a int, ua []int) int {
	if ua[a] != a {
		ua[a] = ufind(ua[a], ua)
	}
	return ua[a]
}

func union(a int, b int, ua []int, vcnt []int, ecnt []int, self_loop []int) {
	ra := ufind(a, ua)
	rb := ufind(b, ua)

	ecnt[rb] += 1
	self_loop[rb] |= self_loop[ra]

	if ra == rb {
		return
	}

	ua[ra] = rb
	vcnt[rb] += vcnt[ra]
	ecnt[rb] += ecnt[ra]
}

func solve(n int, a []int, b []int) int {
	ua := make([]int, n+1)
	vcnt := make([]int, n+1)
	ecnt := make([]int, n+1)
	self_loop := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ua[i] = i
		vcnt[i] = 1
		if a[i-1] == b[i-1] {
			self_loop[a[i-1]] = 1
		}
	}

	for i := 0; i < n; i++ {
		union(a[i], b[i], ua, vcnt, ecnt, self_loop)
	}

	ans := 1
	for i := 1; i <= n; i++ {
		if ua[i] == i {
			if vcnt[i] != ecnt[i] {
				return 0
			}
			if self_loop[i] == 1 {
				ans *= n
			} else {
				ans *= 2
			}
			ans %= MOD
		}
	}
	return ans
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func CF1770D(_r io.Reader, out io.Writer) {
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
		n := readInt(in)
		a := readArrInt(in)
		b := readArrInt(in)
		ans := solve(n, a, b)
		Fprintln(out, ans)
		t--
	}

}

func main() { CF1770D(os.Stdin, os.Stdout) }

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
