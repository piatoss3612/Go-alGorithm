package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
)

// 11815번: 짝수? 홀수?
// https://www.acmicpc.net/problem/11815
// 메모리: 900 KB
// 시간: 4 ms
// 난이도: 실버 4
// 분류: 수학, 정수론
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
	for i := 0; i < N; i++ {
		s := scanString()
		x, _ := big.NewInt(0).SetString(s, 10)

		sqrtX := big.NewInt(0).Sqrt(x)

		if sqrtX.Mul(sqrtX, sqrtX).Cmp(x) == 0 {
			fmt.Fprintf(writer, "%d ", 1)
		} else {
			fmt.Fprintf(writer, "%d ", 0)
		}
	}
	fmt.Fprintln(writer)
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
