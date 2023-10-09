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

	N, M   int
	videos []int
)

// 난이도: Silver 1
// 메모리: 2132KB
// 시간: 20ms
// 분류: 이분 탐색, 매개 변수 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	videos = make([]int, N+1)
	for i := 1; i <= N; i++ {
		videos[i] = scanInt()
	}
}

func Solve() {
	l, r := 1, 10000*100000
	for l <= r {
		m := (l + r) / 2
		res := isPossible(m)
		if res {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	fmt.Fprintln(writer, l)
}

func isPossible(m int) bool {
	cnt, sum := 1, 0
	for i := 1; i <= N; i++ {
		if sum+videos[i] > m {
			if m < videos[i] {
				return false
			}
			cnt++
			sum = videos[i]
		} else {
			sum += videos[i]
		}
	}
	return cnt <= M
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
