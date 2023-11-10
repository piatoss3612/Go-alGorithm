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
	N, Q    int
	sum     [1000002]int
)

// 28018번: 시간이 겹칠까?
// https://www.acmicpc.net/problem/28018
// 난이도: 골드 5
// 메모리: 16484 KB
// 시간: 72 ms
// 분류: 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		s, e := scanInt(), scanInt()
		sum[s], sum[e+1] = sum[s]+1, sum[e+1]-1
	}

	for i := 1; i <= 1000000; i++ {
		sum[i] += sum[i-1]
	}
}

func Solve() {
	Q = scanInt()
	for i := 1; i <= Q; i++ {
		x := scanInt()
		fmt.Fprintln(writer, sum[x])
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
