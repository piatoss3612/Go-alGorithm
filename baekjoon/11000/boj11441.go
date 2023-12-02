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
	N, M    int
	sum     []int
)

// 11441번: 합 구하기
// https://www.acmicpc.net/problem/11441
// 난이도: 실버 3
// 메모리: 2476 KB
// 시간: 64 ms
// 분류: 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	sum = make([]int, N+1)
	for i := 1; i <= N; i++ {
		sum[i] = sum[i-1] + scanInt()
	}

	M = scanInt()
}

func Solve() {
	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		fmt.Fprintln(writer, sum[b]-sum[a-1])
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
