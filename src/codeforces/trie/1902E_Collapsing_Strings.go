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

type Node struct {
	nodes map[int]*Node
	cnt   int
}

func (p *Node) AddNext(i int) *Node {
	if _, ok := p.nodes[i]; !ok {
		p.nodes[i] = NewNode(0)
	}
	p.nodes[i].cnt++

	return p.nodes[i]
}

func NewNode(c int) *Node {
	node := Node{}
	node.nodes = map[int]*Node{}
	node.cnt = c
	return &node
}

// func solve(n int, sa []string, root *Node, r_root *Node) int {
func solve(n int, sa []string, root *Node) int {

	res := 0

	for i := 0; i < n; i++ {

		/*
			r_c := r_root
			for j := 0; j < len(sa[i]); j++ {
				v := int(sa[i][j] - 97)

				if _, ok := r_c.nodes[v]; !ok {
					break
				}

				r_c = r_c.nodes[v]
				res += r_c.cnt
			}
		*/

		c := root
		for j := len(sa[i]) - 1; j >= 0; j-- {
			r_v := int(sa[i][j] - 97)

			if _, ok := c.nodes[r_v]; !ok {
				break
			}

			c = c.nodes[r_v]
			res += c.cnt
		}
	}

	return 2 * res
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func CF1902E(_r io.Reader, out io.Writer) {
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

	n := readInt(in)
	sa := make([]string, n)

	root := NewNode(0)
	// r_root := NewNode(0)
	tt := 0

	for i := 0; i < n; i++ {
		c := root
		// r_c := r_root

		sa[i] = readLine(in)
		tt += len(sa[i])

		for j := 0; j < len(sa[i]); j++ {
			c = c.AddNext(int(sa[i][j] - 97))
			// r_c = r_c.AddNext(int(sa[i][len(sa[i])-j-1] - 97))
		}
	}

	// ans := 2*n*tt - solve(n, sa, root, r_root)
	ans := 2*n*tt - solve(n, sa, root)
	Fprintln(out, ans)
}

func main() { CF1902E(os.Stdin, os.Stdout) }

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
