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

	N       int
	checked [9]bool
	res     [9]int
)

// 난이도: Silver 3
// 메모리: 904KB
// 시간: 48ms
// 분류: 브루트포스 알고리즘, 백트래킹
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
}

func Solve() {
	backtracking(1)
}

func backtracking(n int) {
	if n == N+1 {
		for i := 1; i <= N; i++ {
			fmt.Fprintf(writer, "%d ", res[i])
		}
		fmt.Fprintln(writer)
		return
	}

	for i := 1; i <= N; i++ {
		if !checked[i] {
			checked[i] = true
			res[n] = i
			backtracking(n + 1)
			checked[i] = false
		}
	}
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	return mustParseInt(scanString())
}
