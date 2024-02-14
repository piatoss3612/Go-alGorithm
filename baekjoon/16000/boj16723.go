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
)

// 16723번: 원영이는 ZOAC과 영원하고 싶다
// https://www.acmicpc.net/problem/16723
// 난이도: 실버 1
// 메모리: 856 KB
// 시간: 2572 ms
// 분류: 수학
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
	sum := 0

	for i := 1; i <= N; i++ {
		if i%2 != 0 {
			sum += 1
			continue
		}

		exp := 1
		for exp&i == 0 {
			exp <<= 1
		}

		sum += exp
	}

	fmt.Fprintln(writer, sum*2)
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
