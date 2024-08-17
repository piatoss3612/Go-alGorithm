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
	minPrimeFact [1000001]int
)

// 4937번: It’s All About Three
// hhttps://www.acmicpc.net/problem/4937
// 난이도: 실버 2
// 메모리: 16464 KB
// 시간: 16 ms
// 분류: 수학, 정수론
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	for i := 2; i <= 1000000; i++ {
		if minPrimeFact[i] == 0 {
			minPrimeFact[i] = i
			for j := i * i; j <= 1000000; j += i {
				if minPrimeFact[j] == 0 {
					minPrimeFact[j] = i
				}
			}
		}
	}
}

func Solve() {
	for {
		n := scanInt()
		if n == -1 {
			return
		}

		if n == 0 || n == 1 {
			fmt.Fprintf(writer, "%d NO\n", n)
			continue
		}

		if hasLeastThree(n) {
			fmt.Fprintf(writer, "%d YES\n", n)
		} else {
			fmt.Fprintf(writer, "%d NO\n", n)
		}
	}
}

func hasLeastThree(n int) bool {
	if n == 3 {
		return true
	}
	for n > 1 {
		p := minPrimeFact[n] // n의 가장 작은 소인수
		// 나머지가 3이 아니면 false
		if p%10 != 3 {
			return false
		}
		// p로 나눌 수 있을 때까지 나눔
		for n%p == 0 {
			n /= p
		}
	}

	// n의 모든 소인수가 3으로 끝나면 true
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
