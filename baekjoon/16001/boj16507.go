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

	R, C, Q int
	sum     [1001][1001]int
)

// 난이도: Silver 1
// 메모리: 12560KB
// 시간: 124ms
// 분류: 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	R, C, Q = scanInt(), scanInt(), scanInt()
	for i := 1; i <= R; i++ {
		for j := 1; j <= C; j++ {
			sum[i][j] += scanInt() + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] // 누적 합
		}
	}
}

func Solve() {
	for i := 1; i <= Q; i++ {
		r1, c1, r2, c2 := scanInt(), scanInt(), scanInt(), scanInt()

		psum := sum[r2][c2] - sum[r1-1][c2] - sum[r2][c1-1] + sum[r1-1][c1-1] // 부분 합
		cnt := (r2 - r1 + 1) * (c2 - c1 + 1) // 픽셀의 개수

		fmt.Fprintln(writer, psum/cnt) // 평균(소수점 이하는 버림)
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
