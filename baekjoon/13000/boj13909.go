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

	N int
)

// 13909번: 약수의 개수
// https://www.acmicpc.net/problem/13909
// 난이도: 실버 5
// 메모리: 856 KB
// 시간: 4 ms
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
	// 1~N 중에 약수의 개수가 홀수인 수의 개수를 구해야 한다.
	// 약수의 개수가 홀수인 수는 제곱수이다.
	// 제곱수의 개수를 구하면 된다.
	ans := 0
	for i := 1; i*i <= N; i++ {
		ans++
	}

	fmt.Fprintln(writer, ans)
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
