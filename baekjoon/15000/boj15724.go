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

	N, M int
	sum  [1025][1025]int
	K    int
)

// 난이도: Silver 1
// 메모리: 13600KB
// 시간: 136ms
// 분류: 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + scanInt()
		}
	}
}

func Solve() {
	K = scanInt()
	for i := 1; i <= K; i++ {
		x1, y1, x2, y2 := scanInt(), scanInt(), scanInt(), scanInt()
		fmt.Fprintln(writer, sum[x2][y2]-sum[x1-1][y2]-sum[x2][y1-1]+sum[x1-1][y1-1])
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
