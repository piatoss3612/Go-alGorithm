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

	n       int
	isPrime [MAX + 1]bool
)

const (
	MAX = 100000000
	MOD = 1 << 32
)

// 난이도: Gold 4
// 메모리: 98568KB
// 시간: 1244ms
// 분류: 수학, 정수론, 에라토스테네스의 체
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	n = scanInt()
	eratosthenes()
}

func Solve() {
	ans := 1

	// 1부터 n까지의 모든 자연수의 최소 공배수 구하기

	// 1. 모든 수는 소인수분해를 통해 소수의 곱으로 나타낼 수 있다.
	// 2. 1~n 사이의 소수 각각 최대 거듭제곱 (<=n)을 모두 곱하면 1~n 사이의 모든 수의 최소 공배수를 구할 수 있다.
	for i := 2; i <= n; i++ {
		// i가 소수일 경우
		if isPrime[i] {
			// i의 최대 거듭제곱을 구한다.
			p := i
			for p*i <= n {
				p *= i
			}
			ans = (ans * p) % MOD // 누적해서 곱한다.
		}
	}

	fmt.Fprintln(writer, ans%MOD) // 1~n 사이의 모든 수의 최소 공배수를 출력한다.
}

func eratosthenes() {
	for i := 2; i <= MAX; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= MAX; i++ {
		if isPrime[i] {
			for j := i * i; j <= MAX; j += i {
				isPrime[j] = false
			}
		}
	}
}

func mustParseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return mustParseInt(scanner.Text())
}
