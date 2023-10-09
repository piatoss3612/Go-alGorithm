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

	A, B    int
	isPrime [MAX + 1]bool
)

const MAX = 10000000

// 난이도: Gold 5
// 메모리: 10680KB
// 시간: 88ms
// 분류: 수학, 정수론, 소수 판정, 에라토스테네스의 체
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	A, B = scanInt(), scanInt()
	eratosthenes()
}

func Solve() {
	cnt := 0

	for i := 2; i <= MAX; i++ {
		if isPrime[i] {
			p := i
			// p * i <= B overflow 발생
			for p <= B/i {
				p *= i
				if p >= A {
					cnt++
				}
			}
		}
	}

	fmt.Fprintln(writer, cnt)
}

func eratosthenes() {
	for i := 2; i <= MAX; i++ {
		isPrime[i] = true
	}

	for i := 2; i <= MAX; i++ {
		if isPrime[i] {
			for j := i * i; j <= MAX; j += i {
				isPrime[j] = false
			}
		}
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
