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

	isPrime [10001]bool
	N       int
	arr     [10001]int
)

// 난이도: Silver 2
// 메모리: 1072KB
// 시간: 8ms
// 분류: 소수 판정, 에라토스테네스의 체, 정수론, 수학
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	for i := 2; i <= 10000; i++ {
		isPrime[i] = true
	}
	
	// 에라토스테네스의 체
	for i := 2; i*i <= 10000; i++ {
		if isPrime[i] {
			for j := i * i; j <= 10000; j += i {
				isPrime[j] = false
			}
		}
	}

	N = scanInt()

	for i := 1; i <= N; i++ {
		arr[i] = scanInt()
	}
}

func Solve() {
	for i := 1; i <= N; i++ {
		a, b := Goldbach(arr[i]) // 골드바흐의 추측
		fmt.Fprintf(writer, "%d %d\n", a, b)
	}
}

func Goldbach(n int) (int, int) {
	// n/2부터 시작해서 내림차순으로 탐색하면 두 수의 차이가 가장 작은 것을 우선적으로 찾을 수 있다.
	for i := n / 2; i >= 2; i-- {
		if isPrime[i] && isPrime[n-i] {
			return swap(i, n-i) // 작은 수부터 출력해야 하므로 swap
		}
	}
	return 0, 0
}

func swap(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
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
