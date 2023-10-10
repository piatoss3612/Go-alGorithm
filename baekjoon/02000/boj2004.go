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
)

// 2004번: 조합 0의 개수
// https://www.acmicpc.net/problem/2004
// 난이도: 실버 2
// 메모리: 856KB
// 시간: 4ms
// 분류: 수학, 정수론
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
	if N == M || M == 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	// 조합 nCm = n! / (m! * (n-m)!)
	// fives = n!의 5의 개수 - m!의 5의 개수 - (n-m)!의 5의 개수
	// twos = n!의 2의 개수 - m!의 2의 개수 - (n-m)!의 2의 개수
	// 2의 개수와 5의 개수 중 작은 것이 0의 개수

	K := N - M

	fives := 0
	i := 5

	for i <= N {
		fives += N/i - M/i - K/i
		i *= 5
	}

	j := 2
	twos := 0

	for j <= N {
		twos += N/j - M/j - K/j
		j *= 2
	}

	fmt.Fprintln(writer, min(fives, twos))
}

func min(a, b int) int {
	if a < b {
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
