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
	T, N    int
)

// 13412번: 서로소 쌍
// https://www.acmicpc.net/problem/13412
// 메모리: 836 KB
// 시간: 12 ms
// 분류: 수학, 정수론, 유클리드 호제법
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
}

func Setup() {
	T = scanInt()
	for i := 0; i < T; i++ {
		Solve()
	}
}

func Solve() {
	N = scanInt()

	cnt := 0

	for i := 1; i*i <= N; i++ {
		// i가 N의 약수이고, gcd(N/i, i) == 1이면 N/i와 i는 N을 최소공배수로 하는 서로소 쌍이다.
		if N%i == 0 {
			if gcd(N/i, i) == 1 {
				cnt++
			}
		}
	}

	fmt.Fprintln(writer, cnt)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
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
