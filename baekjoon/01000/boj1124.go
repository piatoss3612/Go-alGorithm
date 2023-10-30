package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner      = bufio.NewScanner(os.Stdin)
	writer       = bufio.NewWriter(os.Stdout)
	A, B         int
	isPrime      [100001]bool
	minFactor    [100001]int
	primeFactors [100001]int
)

// 1124번: 언더프라임
// https://www.acmicpc.net/problem/1124
// 난이도: 실버 1
// 메모리: 2512 KB
// 시간: 8 ms
// 분류: 수학, 정수론, 소수 판정, 에라토스테네스의 체
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	A, B = scanInt(), scanInt()

	for i := 2; i <= B; i++ {
		isPrime[i] = true
	}

	for i := 2; i <= B; i++ {
		if isPrime[i] {
			minFactor[i] = i
			for j := i * i; j <= B; j += i {
				isPrime[j] = false
				minFactor[j] = i
			}
		}
	}

	for i := 2; i <= B; i++ {
		primeFactors[i] = primeFactors[i/minFactor[i]] + 1
	}
}

func Solve() {
	count := 0
	for i := A; i <= B; i++ {
		if isPrime[primeFactors[i]] {
			count++
		}
	}

	fmt.Fprintln(writer, count)
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