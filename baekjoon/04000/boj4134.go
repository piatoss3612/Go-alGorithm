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
	T       int
)

// 4134번: 다음 소수
// https://www.acmicpc.net/problem/4134
// 난이도: 실버 4
// 메모리: 820 KB
// 시간: 264 ms
// 분류: 수학, 정수론, 소수 판정, 브루트포스
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	T = scanInt()
}

func Solve() {
	for i := 1; i <= T; i++ {
		n := scanInt()

		for {
			if isPrime(n) {
				fmt.Fprintln(writer, n)
				break
			}

			n++
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
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
