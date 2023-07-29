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

	isPrime [246913]bool
	sum     [246913]int
	N       int
)

// 난이도: Silver 2
// 메모리: 3080KB
// 시간: 8ms
// 분류: 수학, 정수론, 소수 판정, 에라토스테네스의 체, 누적 합
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	for i := 2; i <= 246912; i++ {
		isPrime[i] = true
	}

	// 에라토스테네스의 체
	for i := 2; i*i <= 246912; i++ {
		if isPrime[i] {
			for j := i * i; j <= 246912; j += i {
				isPrime[j] = false
			}
		}
	}

	// 누적 합: sum[i] = 1부터 i까지의 소수의 개수
	for i := 2; i <= 246912; i++ {
		if isPrime[i] {
			sum[i] = sum[i-1] + 1
		} else {
			sum[i] = sum[i-1]
		}
	}
}

func Solve() {
	for {
		N = scanInt()
		if N == 0 {
			return
		}

		fmt.Fprintln(writer, sum[2*N]-sum[N]) // N < x <= 2N인 x 중 소수의 개수
	}
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
