// https://codeforces.com/problemset/problem/1900/D
/*
Let a, b, and c be integers. We define function f(a,b,c) as follows:
Order the numbers a, b, c in such a way that a≤b≤c. Then return gcd(a,b), where gcd(a,b) denotes the greatest common divisor (GCD) of integers a and b.\
So basically, we take the gcd of the 2 smaller values and ignore the biggest one.
You are given an array a of n elements. Compute the sum of f(ai,aj,ak) for each i, j, k, such that 1≤i<j<k≤n.

2 3 6 12 17
2 3 * 3
2 6 * 2

17 3 6 12 2
2 3 * 1
3 2 * 2

1. sort array: a,b,c,d,e
2. gcd(a,b)*3 + gcd(a,c)*2 + gcd(a,d)
   gcd(b,c)*2 + gcd(b,d)*1
   gcd(c,d)*1
*/
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

// var divisors [][]int = make([][]int, 100001)

var MAX int = 100000

// O(N*logN)
func solve_optimize(n int, a []int) int64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	var cnt_a []int64 = make([]int64, MAX+1)
	var up_a []int64 = make([]int64, MAX+1)
	var dp_a []int64 = make([]int64, MAX+1)

	for _, v := range a {
		cnt_a[v]++
	}
	pre := 0
	for i := MAX; i > 0; i-- {
		up_a[i] = int64(pre)
		pre += int(cnt_a[i])
	}

	var ans int64 = 0
	for i := MAX; i > 0; i-- {
		var pre1 int64 = 0
		var pre2 int64 = 0
		for j := i; j <= MAX; j += i {
			// c[j] * (c[j] - 1) / 2 => combination formula j happens multiple times. ex  1,1,1 => (cur, cur) + after
			// pre1 * c[j] => pre + (cur) + after
			pre2 += (pre1*cnt_a[j] + cnt_a[j]*(cnt_a[j]-1)/2) * up_a[j]
			// (cur, cur, cur)
			pre2 += cnt_a[j] * (cnt_a[j] - 1) * (cnt_a[j] - 2) / 6
			// pre + (cur,cur)
			pre2 += pre1 * cnt_a[j] * (cnt_a[j] - 1) / 2
			dp_a[i] -= dp_a[j]
			pre1 += cnt_a[j]
		}
		dp_a[i] += pre2
		ans += dp_a[i] * int64(i)
	}

	return ans
}

// O(N*logN*logN)
/*
func solve(n int, a []int) int64 {
	var ans int64 = 0
	mp := map[int]int{}

	sort.Ints(a)
	f_count(a[0], mp)

	for i := 1; i < n-1; i++ {
		fa := divisors[a[i]]
		mp_d := map[int]int{}

		for j := len(fa) - 1; j >= 0; j-- {

			cnt := mp_get(mp, fa[j]) - mp_get(mp_d, fa[j])
			ans += int64(cnt * fa[j] * (n - 1 - i))

			for _, f := range divisors[fa[j]] {
				mp_inc(mp_d, f, cnt)
			}
		}

		f_count(a[i], mp)
	}

	return ans
}

func f_count(v int, mp map[int]int) {
	for _, f := range divisors[v] {
		mp_inc(mp, f, 1)
	}
}

func mp_get(mp map[int]int, k int) int {
	if _, found := mp[k]; found {
		return mp[k]
	}
	return 0
}

func mp_inc(mp map[int]int, k int, cnt int) {
	if _, found := mp[k]; found {
		mp[k] += cnt
	} else {
		mp[k] = cnt
	}
}
*/

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	/*
		for i := 1; i < 100001; i++ {
			divisors[i] = make([]int, 0)
		}

		// precalculate divisors Nlog(N)
		for i := 1; i < 100001; i++ {
			for j := i; j < 100001; j += i {
				divisors[j] = append(divisors[j], i)
			}
		}
	*/

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
	t := readInt(in)
	for t > 0 {
		n := readInt(in)
		a := readArrInt(in)

		// fmt.Println(solve(n, a))
		fmt.Println(solve_optimize(n, a))

		t--
	}

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
