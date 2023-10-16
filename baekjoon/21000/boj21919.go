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
	isPrime [1000001]bool
)

// 21919번: 소수 최소공배수
// https://www.acmicpc.net/problem/21919
// 난이도: 실버 3
// 메모리: 1840 KB
// 시간: 8 ms
// 분류: 수학, 정수론, 소수 판정, 에라토스테네스의 체
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= 1000000; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= 1000000; i++ {
		if isPrime[i] {
			for j := i * i; j <= 1000000; j += i {
				isPrime[j] = false
			}
		}
	}
}

func Solve() {
	hasPrime := false
	ans := 0

	for i := 1; i <= N; i++ {
		x := scanInt()
		if isPrime[x] {
			if !hasPrime {
				hasPrime = true
				ans = x
			} else {
				ans = lcm(ans, x)
			}
		}
	}

	if hasPrime {
		fmt.Fprintln(writer, ans)
	} else {
		fmt.Fprintln(writer, -1)
	}
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
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
