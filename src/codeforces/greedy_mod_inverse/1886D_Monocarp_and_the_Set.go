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
6 4
<?>?>
1 ?
4 <
5 <
1 >

??>?>
??><>
??><<
>?><<
*/
func CF1886D(_r io.Reader, out io.Writer) {
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

	nm := readArrInt(in)

	sb := []byte(readLine(in))
	invalid := false
	var ans int64 = 1
	var MOD int64 = 998244353
	// build first result
	for i := 0; i < len(sb); i++ {
		symbol := string(sb[i])
		if symbol == "?" {
			if i == 0 {
				invalid = true
			} else {
				ans *= int64(i)
				ans %= MOD
			}
		}
	}
	Fprintln(out, get_ans(ans, invalid))

	for nm[1] > 0 {
		ic := readLineNumbs(in)

		i, _ := strconv.Atoi(string(ic[0]))
		i--
		org_c := string(sb[i])
		c := ic[1]

		if org_c != "?" && c == "?" {
			if i == 0 {
				invalid = true
			} else {
				ans *= int64(i)
				ans %= MOD
			}
		} else if org_c == "?" && c != "?" {
			if i == 0 {
				invalid = false
			} else {
				ans *= mod_inverse(int64(i), MOD)
				ans %= MOD
			}
		}

		sb[i] = []byte(ic[1])[0]

		Fprintln(out, get_ans(ans, invalid))
		nm[1]--
	}
}

func get_ans(ans int64, invalid bool) int64 {
	if invalid {
		return 0
	}
	return ans
}

/** modular multiplicative inverse **/
/*
https://en.wikipedia.org/wiki/Fermat%27s_little_theorem
*/
func mod_inverse(x int64, mod int64) int64 {
	return mod_power(x, mod-2, mod)
}

func mod_power(x int64, y int64, mod int64) int64 {
	if y == 0 {
		return 1
	}
	p := mod_power(x, y/2, mod) % mod
	p = (p * p) % mod
	if y%2 == 1 {
		p = (p * x) % mod
	}
	return p
}

func main() { CF1886D(os.Stdin, os.Stdout) }

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
