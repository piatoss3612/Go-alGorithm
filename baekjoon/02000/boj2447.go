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
	arr     [10000][10000]byte
)

// 2447번: 별 찍기 - 10
// hhttps://www.acmicpc.net/problem/2447
// 난이도: 골드 5
// 메모리: 14252 KB
// 시간: 364 ms
// 분류: 재귀, 분할 정복
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
	star(0, 0, N) // 0, 0 부터 너비와 높이가 N인 패턴을 그린다.

	// 출력
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Fprintf(writer, "%c", arr[i][j])
		}
		fmt.Fprintln(writer)
	}
}

func star(x, y, n int) {
	// base case: n이 1이면 더 이상 분할할 수 없으므로 별을 그린다.
	if n == 1 {
		arr[x][y] = '*'
		return
	}

	n /= 3

	// 3x3, 9개의 구역으로 나누어서 패턴을 그린다.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// 가운데(5번째) 구역은 공백으로 채운다.
			// 나머지 구역은 다시 3x3으로 나누어서 패턴을 그린다.
			if i == 1 && j == 1 {
				blank(x+i*n, y+j*n, n)
			} else {
				star(x+i*n, y+j*n, n)
			}
		}
	}
}

func blank(x, y, n int) {
	for i := x; i < x+n; i++ {
		for j := y; j < y+n; j++ {
			arr[i][j] = ' '
		}
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
