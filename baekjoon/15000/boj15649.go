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

	N, M  int
	arr   [9]int
	taken [9]bool
)

// 난이도: Silver 3
// 메모리: 904KB
// 시간: 44ms
// 분류: 백트래킹
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
}

func Solve() {
	backTracking(0)
}

func backTracking(turn int) {
	if turn == M {
		for i := 1; i <= M; i++ {
			fmt.Fprintf(writer, "%d ", arr[i])
		}
		fmt.Fprintln(writer)
		return
	}

	for i := 1; i <= N; i++ {
		if !taken[i] {
			taken[i] = true
			arr[turn+1] = i
			backTracking(turn + 1)
			taken[i] = false
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
