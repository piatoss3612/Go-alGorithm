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

	N     int
	arr   [9]int
	taken [9]bool
	temp  [9]int
	ans   int
)

// 난이도: Silver 2
// 메모리: 912KB
// 시간: 12ms
// 분류: 백트래킹, 브루트포스
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	backTracking(1)
	fmt.Fprintln(writer, ans)
}

func backTracking(turn int) {
	if turn == N+1 {
		sum := 0
		for i := 1; i <= N-1; i++ {
			sum += abs(temp[i] - temp[i+1])
		}
		ans = max(ans, sum)
		return
	}

	for i := 1; i <= N; i++ {
		if !taken[i] {
			taken[i] = true
			temp[turn] = arr[i]
			backTracking(turn + 1)
			temp[turn] = 0
			taken[i] = false
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
