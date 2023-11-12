package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N, M, R  int
	stake    []int
	flagpole []int
)

// 30459번: 현수막 걸기
// https://www.acmicpc.net/problem/30459
// 난이도: 골드 5
// 메모리: 1200 KB
// 시간: 152 ms
// 분류: 정렬, 이분 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M, R = scanInt(), scanInt(), scanInt()
	stake = make([]int, N)
	flagpole = make([]int, M)
	for i := 0; i < N; i++ {
		stake[i] = scanInt()
	}
	for i := 0; i < M; i++ {
		flagpole[i] = scanInt()
	}

	sort.Ints(stake)
	sort.Ints(flagpole)
}

func Solve() {
	var maxArea float64 = -1

	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			bottom := stake[j] - stake[i]

			l, r := 0, M-1
			for l <= r {
				mid := (l + r) / 2

				area := (float64(bottom) * float64(flagpole[mid])) / 2

				if area <= float64(R) {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}

			if r < 0 {
				continue
			}

			area := (float64(bottom) * float64(flagpole[r])) / 2
			maxArea = max(maxArea, area)
		}
	}

	if maxArea == -1 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintf(writer, "%.1f\n", maxArea)
	}
}

func max(a, b float64) float64 {
	if a > b {
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
