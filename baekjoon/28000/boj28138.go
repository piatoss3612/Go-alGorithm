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
	N, R    int
)

// 28138번: 재밌는 나머지 연산
// https://www.acmicpc.net/problem/28138
// 난이도: 실버 3
// 메모리: 892 KB
// 시간: 16 ms
// 분류: 수학, 정수론, 소수 판정
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, R = scanInt(), scanInt()
}

func Solve() {
	n := N - R
	if n <= R {
		fmt.Fprintln(writer, 0)
		return
	}

	sum := 0

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			a, b := i, n/i

			// 제곱근을 중복해서 더해주는 경우 제거
			if a == b && a > R {
				sum += a
				continue
			}

			if a > R {
				sum += a
			}

			if b > R {
				sum += b
			}
		}
	}

	fmt.Fprintln(writer, sum)
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
