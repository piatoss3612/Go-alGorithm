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

	N, Q  int
	gap   []int
	birds []int
)

// 26152번: 플래피 버드 스코어링
// hhttps://www.acmicpc.net/problem/26152
// 난이도: 실버 1
// 메모리: 7212 KB
// 시간: 172 ms
// 분류: 이분 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	gap = make([]int, N)
	for i := 0; i < N; i++ {
		gap[i] = scanInt()
	}

	for i := 0; i < N; i++ {
		gap[i] -= scanInt()
	}

	for i := 1; i < N; i++ {
		gap[i] = min(gap[i], gap[i-1])
	}

	Q = scanInt()
	birds = make([]int, Q)
	for i := 0; i < Q; i++ {
		birds[i] = scanInt()
	}
}

func Solve() {
	for _, bird := range birds {
		l, r := 0, N-1
		for l <= r {
			mid := (l + r) / 2
			if gap[mid] >= bird {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		fmt.Fprintln(writer, l)
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
