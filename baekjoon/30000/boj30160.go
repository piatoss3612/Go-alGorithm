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
	arr     [100001]int
	sum     [100001]int
)

// 30160번: 제곱 가중치
// https://www.acmicpc.net/problem/30160
// 난이도: 골드 5
// 메모리: 2740 KB
// 시간: 44 ms
// 분류: 수학, 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	sum[1] = arr[1]
	x := 3 * arr[1]
	y := 2 * arr[1]

	// sum[i] - sum[i-1] = (i*2 - 1) * arr[1] + (i*2 - 3) * arr[2] + ... + 1 * arr[i]
	for i := 2; i <= N; i++ {
		x += arr[i]
		sum[i] = sum[i-1] + x
		y += 2 * arr[i]
		x += y
	}

	for i := 1; i <= N; i++ {
		fmt.Fprintf(writer, "%d ", sum[i])
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
