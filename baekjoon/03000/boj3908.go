package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	primes  []int
	dp      [1121][15]int // dp[i][j] = j개의 서로다른 소수를 더해 i를 만들 수 있는 경우의 수
)

// dp 전처리
func init() {
	primes = *eratosthenes(1120) // 1120까지의 소수들 구하기

	dp[0][0] = 1 // 소수 1개로 i를 만들 수 있는 경우의 수를 따지기 위해 dp[0][0]을 1로 설정

	// 가장 작은 소수 2부터 오름차순으로 진행
	for _, prime := range primes {
		/*
			i의 값이 1120부터 prime까지 역순으로 감소하면서
			현재 prime보다 작은 소수들의 합으로만 이루어진 수에 prime을 더함으로써
			중복된 연산을 제거할 수 있다

			prime=2:
			dp[2][1] = 1

			prime=3:
			dp[2][1] = 1, dp[3][1] = 1, dp[5][2] = 1

			prime=5:
			dp[2][1] = 1, dp[3][1] = 1, dp[5][1] = 1, dp[5][2] = 1, dp[7][2] = 1, dp[8][2] = 1, dp[10][3] = 1

			...
		*/
		for i := 1120; i >= prime; i-- {
			for j := 14; j >= 1; j-- {
				dp[i][j] += dp[i-prime][j-1]
			}
		}
	}

}

// 메모리: 1196KB
// 시간: 8ms
// 서로 다른 k개의 소수로 n을 만드는 방법의 수를 구하는 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()

	for i := 1; i <= t; i++ {
		fmt.Fprintln(writer, dp[scanInt()][scanInt()])
	}
}

// 에라토스테네스의 체
func eratosthenes(n int) *[]int {
	primes := []int{}
	check := make([]bool, n+1)

	sqrtn := int(math.Sqrt(float64(n)))

	// 최적화1: n의 제곱근 이하의 수까지만 체크
	for i := 2; i <= sqrtn; i++ {
		if !check[i] {
			check[i] = true
			primes = append(primes, i)
			// 최적화2: i가 소수인 경우, i의 배수를 i*i부터 제거
			for j := i * i; j <= n; j += i {
				check[j] = true
			}
		}
	}

	// 체로 걸러지고 남은 소수들 주워담기
	for i := sqrtn + 1; i <= n; i++ {
		if !check[i] {
			primes = append(primes, i)
		}
	}

	return &primes // 포인터 반환
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
