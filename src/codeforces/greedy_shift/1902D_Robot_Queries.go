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

func solve(n int, x int, y int, l int, r int, mp map[int]map[int][]int, r_mp map[int]map[int][]int, ps [][]int, r_ps [][]int) int {
	if x == 0 && y == 0 {
		return 1
	}

	// 0 - L-1
	if l > 1 {
		if _, x_ok := mp[x]; x_ok {
			if _, y_ok := mp[x][y]; y_ok {
				if mp[x][y][0] < l-1 {
					return 1
				}
			}
		}
	}

	// L - R => - (0 - L-1) + (R+1 - N)
	x2 := x - ps[l-1][0] + r_ps[r][0]
	y2 := y - ps[l-1][1] + r_ps[r][1]

	if _, x_ok := r_mp[x2]; x_ok {
		if _, y_ok := r_mp[x2][y2]; y_ok {
			if bs2(l-1, r-1, r_mp[x2][y2]) {
				return 1
			}
		}
	}

	// R+1 - N => (L, R) - (R, L)
	x3 := x + (ps[r][0] - ps[l-1][0]) - (r_ps[l-1][0] - r_ps[r][0])
	y3 := y + (ps[r][1] - ps[l-1][1]) - (r_ps[l-1][1] - r_ps[r][1])

	if r < n {
		if _, x_ok := mp[x3]; x_ok {
			if _, y_ok := mp[x3][y3]; y_ok {
				if r <= mp[x3][y3][len(mp[x3][y3])-1] {
					return 1
				}
			}
		}
	}

	return 0
}

// 8  4(3) 2  1
func bs2(i int, j int, r_a []int) bool {

	l := 0
	r := len(r_a) - 1

	for l < r {
		m := (l + r + 1) / 2
		if i > r_a[m] {
			r = m - 1
		} else {
			l = m
		}
	}

	return i <= r_a[l] && r_a[l] <= j
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func CF1902D(_r io.Reader, out io.Writer) {
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

	nq := readArrInt(in)
	s := readLine(in)
	// x, y, indexes
	mp := map[int]map[int][]int{}
	r_mp := map[int]map[int][]int{}

	ps := make([][]int, nq[0]+1)
	r_ps := make([][]int, nq[0]+1)
	ps[0] = make([]int, 2)
	r_ps[nq[0]] = make([]int, 2)
	for i := 0; i < len(s); i++ {
		x := 0
		y := 0
		if s[i] == 85 { // U
			y++
		} else if s[i] == 68 { // D
			y--
		} else if s[i] == 76 { // L
			x--
		} else { // R
			x++
		}

		r_x := 0
		r_y := 0
		if s[nq[0]-i-1] == 85 { // U
			r_y++
		} else if s[nq[0]-i-1] == 68 { // D
			r_y--
		} else if s[nq[0]-i-1] == 76 { // L
			r_x--
		} else { // R
			r_x++
		}

		ps[i+1] = make([]int, 2)
		ps[i+1][0] = ps[i][0] + x
		ps[i+1][1] = ps[i][1] + y

		r_ps[nq[0]-i-1] = make([]int, 2)
		r_ps[nq[0]-i-1][0] = r_ps[nq[0]-i][0] + r_x
		r_ps[nq[0]-i-1][1] = r_ps[nq[0]-i][1] + r_y

		if _, ok := mp[ps[i+1][0]]; !ok {
			mp[ps[i+1][0]] = map[int][]int{}
		}
		if _, ok := mp[ps[i+1][0]][ps[i+1][1]]; !ok {
			mp[ps[i+1][0]][ps[i+1][1]] = make([]int, 0)
		}
		mp[ps[i+1][0]][ps[i+1][1]] = append(mp[ps[i+1][0]][ps[i+1][1]], i)

		if _, ok := r_mp[r_ps[nq[0]-i-1][0]]; !ok {
			r_mp[r_ps[nq[0]-i-1][0]] = map[int][]int{}
		}
		if _, ok := r_mp[r_ps[nq[0]-i-1][0]][r_ps[nq[0]-i-1][1]]; !ok {
			r_mp[r_ps[nq[0]-i-1][0]][r_ps[nq[0]-i-1][1]] = make([]int, 0)
		}
		r_mp[r_ps[nq[0]-i-1][0]][r_ps[nq[0]-i-1][1]] = append(r_mp[r_ps[nq[0]-i-1][0]][r_ps[nq[0]-i-1][1]], nq[0]-i-1)
	}

	for i := 0; i < nq[1]; i++ {
		xylr := readArrInt(in)
		ans := solve(nq[0], xylr[0], xylr[1], xylr[2], xylr[3], mp, r_mp, ps, r_ps)
		if ans == 1 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}

	}
}

func main() { CF1902D(os.Stdin, os.Stdout) }

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
