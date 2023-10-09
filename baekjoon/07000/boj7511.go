package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	T       int
	n, k, m int
	rels    []int
)

// 메모리: 5572KB
// 시간: 180ms
// 분리 집합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	T = scanInt()
	for i := 1; i <= T; i++ {
		n, k = scanInt(), scanInt()
		rels = make([]int, n)
		for j := 0; j < n; j++ {
			rels[j] = j
		}

		for j := 0; j < k; j++ {
			a, b := scanInt(), scanInt()
			union(a, b)
		}

		m = scanInt()

		fmt.Fprintf(writer, "Scenario %d:\n", i)

		for j := 0; j < m; j++ {
			u, v := scanInt(), scanInt()
			if find(u) == find(v) {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		}

		if i != T {
			fmt.Fprintln(writer)
		}
	}
}

func find(x int) int {
	if rels[x] == x {
		return x
	}
	rels[x] = find(rels[x])
	return rels[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x != y {
		rels[y] = x
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
