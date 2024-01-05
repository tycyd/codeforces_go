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

func solve(n int, a []int) int {

	memo := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		memo[i] = map[int]int{}
	}

	return dfs(1, 0, a, memo)
}

// top -> bottom dp
func solve2(n int, a []int) int {

	OFFSET := 90000

	dp := make([][]int, n)
	dp[1] = make([]int, OFFSET*2+1)
	dp[1][a[1]+OFFSET] = 1

	for i := 2; i < n; i++ {

		dp[i] = make([]int, OFFSET*2+1)

		for j := -OFFSET + a[i]; j <= OFFSET-a[i]; j++ {
			// a[i]+ pj = j or a[i]-pj = j
			// pj = j-a[i] or pj = a[i]-j
			r := dp[i-1][j-a[i]+OFFSET]
			if j-a[i] != 0 {
				r += dp[i-1][a[i]-j+OFFSET]
				r %= MOD
			}
			dp[i][j+OFFSET] = max(dp[i][j+OFFSET], r)
		}
	}

	res := 0
	for _, v := range dp[n-1] {
		res += v
		res %= MOD
	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// bottom -> top dp
func dfs(i int, pv int, a []int, memo []map[int]int) int {
	if i == len(a)-1 {
		return 1
	}

	if _, ok := memo[i][pv]; ok {
		return memo[i][pv]
	}

	res := dfs(i+1, a[i]+pv, a, memo)
	if a[i]+pv != 0 {
		res += dfs(i+1, -a[i]-pv, a, memo)
		res %= MOD
	}

	memo[i][pv] = res
	return res
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func CF1783D(_r io.Reader, out io.Writer) {
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
	n := readInt(in)
	a := readArrInt(in)

	ans := solve2(n, a)

	// print
	Fprint(out, ans)
}

func main() { CF1783D(os.Stdin, os.Stdout) }

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
