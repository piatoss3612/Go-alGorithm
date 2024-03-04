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
	T, N    int
)

// 28706번: 럭키 세븐
// hhttps://www.acmicpc.net/problem/28706
// 난이도: 골드 5
// 메모리: 2556 KB
// 시간: 64 ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
}

func Setup() {
	T = scanInt()
	for t := 0; t < T; t++ {
		Solve()
	}
}

func Solve() {
	N = scanInt()

	dp := make([][7]bool, N+1)
	dp[0][1] = true

	for i := 1; i <= N; i++ {
		op1, v1, op2, v2 := scanString(), scanInt(), scanString(), scanInt()

		for j := 0; j < 7; j++ {
			if !dp[i-1][j] {
				continue
			}

			dp[i][calc(op1, j, v1)] = true
			dp[i][calc(op2, j, v2)] = true
		}
	}

	if dp[N][0] {
		fmt.Fprintln(writer, "LUCKY")
	} else {
		fmt.Fprintln(writer, "UNLUCKY")
	}
}

func calc(op string, k, v int) int {
	switch op {
	case "+":
		return (k + v) % 7
	case "*":
		return (k * v) % 7
	}

	return 0
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
