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
	Q       int
	a, d, x int
)

// 28127번: 숫자탑과 쿼리
// https://www.acmicpc.net/problem/28127
// 난이도: 골드 5
// 메모리: 5888 KB
// 시간: 424 ms
// 분류: 이분 탐색, 사칙연산
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	Q = scanInt()
}

func Solve() {
	for i := 0; i < Q; i++ {
		a, d, x = scanInt(), scanInt(), scanInt()

		// an = a + (n-1)d
		// sn = na + 0 * d + 1 * d + 2 * d + ... + (n-1) * d = na + d * (n-1) * n / 2
		l, r := 1, x
		for l <= r {
			mid := (l + r) / 2

			s := a*mid + d*(mid-1)*mid/2

			if s < x {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		rowStart := a*(l-1) + d*(l-2)*(l-1)/2 + 1 // l번째 행의 시작 숫자
		columnX := x - rowStart + 1               // l번째 행에서 x번째 숫자의 위치

		fmt.Fprintln(writer, l, columnX)
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
