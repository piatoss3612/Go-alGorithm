package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N, K     int
	notPrime [1000001]bool
	prime    []int
)

// 31827번: 소수 수열
// hhttps://www.acmicpc.net/problem/31827
// 난이도: 골드 5
// 메모리: 3108 KB
// 시간: 56 ms
// 분류: 수학, 정수론, 소수 판정, 에라토스테네스의 체
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()

	prime = make([]int, 0, 1000000)

	for i := 2; i <= 1000000; i++ {
		if notPrime[i] {
			continue
		}

		prime = append(prime, i)

		for j := i * i; j <= 1000000; j += i {
			notPrime[j] = true
		}
	}

}

func Solve() {
	cnt := 0

	for i := 0; i < len(prime); i++ {
		// K로 나눈 나머지가 1인 소수를 선택하면
		// K개의 연속된 소수가 모두 나머지가 1이 되므로, 나머지의 합이 K가 된다.
		// 따라서 K개의 연속된 소수의 합은 K의 배수가 된다.
		if prime[i]%K == 1 {
			fmt.Fprintf(writer, "%d ", prime[i])
			cnt++
			if cnt == N {
				break
			}
		}
	}

	fmt.Fprintln(writer)
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
