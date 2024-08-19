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

// 26015번: Enjoyable Entree
// hhttps://www.acmicpc.net/problem/26015
// 난이도: 실버 1
// 메모리: 856 KB
// 시간: 4 ms
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
	if N == 1 {
		fmt.Fprintln(writer, 100, 0)
		return
	}

	if N == 2 {
		fmt.Fprintln(writer, 0, 100)
		return
	}

	// 계속 더하다보면 1/3, 2/3로 수렴한다.
	if N >= 30 {
		fmt.Fprintln(writer, "33.333333 66.666667")
		return
	}

	a := [2]float64{100, 0}
	b := [2]float64{0, 100}

	for i := 3; i <= N; i++ {
		c0, c1 := (a[0]+b[0])/2, (a[1]+b[1])/2
		a[0], a[1] = b[0], b[1]
		b[0], b[1] = c0, c1
	}

	fmt.Fprintf(writer, "%0.6f %0.6f\n", b[0], b[1])
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
