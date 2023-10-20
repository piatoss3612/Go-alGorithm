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
	N, X    int
	arr     []int
)

// 21920번: 서로소 평균
// https://www.acmicpc.net/problem/21920
// 난이도: 실버 4
// 메모리: 5184 KB
// 시간: 124 ms
// 분류: 수학, 정수론, 유클리드 호제법
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	arr = make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = scanInt()
	}
	X = scanInt()
}

func Solve() {
	sum := 0
	cnt := 0

	for i := 0; i < N; i++ {
		if gcd(X, arr[i]) == 1 {
			sum += arr[i]
			cnt++
		}
	}

	if cnt == 0 {
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintf(writer, "%.6f\n", float64(sum)/float64(cnt))
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
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
