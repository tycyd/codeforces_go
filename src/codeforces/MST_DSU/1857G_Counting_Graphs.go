package main

import (
	"bufio"
	"flag"
	. "fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"
)

var MOD int = 998244353

func solve(n int, S int, uvw_a [][]int) int {

	sort.Slice(uvw_a, func(i, j int) bool {
		return uvw_a[i][2] < uvw_a[j][2]
	})

	ua := make([][]int, 1)
	for i := 1; i <= n; i++ {
		ua = append(ua, []int{i, 1})
	}

	ans := 1
	for _, uvw := range uvw_a {
		u_r := ufind(uvw[0], ua)
		v_r := ufind(uvw[1], ua)

		if u_r == v_r {
			continue
		}

		ans *= mod_power(S-uvw[2]+1, ua[u_r][1]*ua[v_r][1]-1, MOD)
		ans %= MOD

		union(u_r, v_r, ua)
	}

	return ans
}

func ufind(a int, ua [][]int) int {
	if a != ua[a][0] {
		ua[a][0] = ufind(ua[a][0], ua)
		ua[a][1] = ua[ua[a][0]][1]
	}
	return ua[a][0]
}

// a -> b
func union(a int, b int, ua [][]int) {
	a_r := ufind(a, ua)
	b_r := ufind(b, ua)

	ua[b_r][1] += ua[a_r][1]
	ua[a_r][0] = b_r
}

func mod_power(x int, y int, mod int) int {
	if x <= 0 || y == 0 {
		return 1
	}
	p := mod_power(x, y/2, mod) % mod
	p = (p * p) % mod
	if y%2 == 1 {
		p = (p * x) % mod
	}
	return p
}

func dif(c int, S int) int {
	if S-c > 0 {
		return S - c + 1
	}
	return 1
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func CF1857G(_r io.Reader, out io.Writer) {
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
		nS := readArrInt(in)
		uvw_a := make([][]int, 0)
		for i := 1; i < nS[0]; i++ {
			uvw_a = append(uvw_a, readArrInt(in))
		}
		ans := solve(nS[0], nS[1], uvw_a)
		Fprintln(out, ans)
		t--
	}
}

func main() { CF1857G(os.Stdin, os.Stdout) }

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
