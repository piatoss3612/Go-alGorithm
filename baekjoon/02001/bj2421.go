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
	primes  map[int]bool    // 999999이하의 소수를 저장하는 맵
	dp      [1000][1000]int // 최대 입력값 999에 대해 999999까지 연산 과정을 저장하는 배열
)

// 메모리: 31540KB
// 시간: 92ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()
	primes = getPrimes() // 소수 구하기

	ans := rec(n, n) // 입력값 n과 n을 이어붙인 수를 구하는 과정에서 찾을 수 있는 소수의 최대 갯수

	fmt.Fprintln(writer, ans)
}

func rec(x, y int) int {
	// x 또는 y가 1보다 작은 경우
	if x < 1 || y < 1 {
		return 0
	}

	// x와 y가 모두 1인 경우
	if x == 1 && y == 1 {
		return 0
	}

	// dp[x][y]의 값이 이미 구해진 경우
	if dp[x][y] != 0 {
		return dp[x][y]
	}

	ret := &dp[x][y]
	// dp[x][y]의 값은 (x-1, y)까지 연산 과정에서 구할 수 있는 소수의 최대 갯수 또는
	// (x, y-1)까지 연산 과정에서 구할 수 있는 소수의 최대 갯수를 비교한 최댓값 + x, y가 소수인 경우
	*ret = max(rec(x-1, y), rec(x, y-1)) + calc(x, y)
	return *ret
}

// 에라토스테네스의 체
func getPrimes() map[int]bool {
	tmp := make([]int, 1000000)

	primes := make(map[int]bool)
	for i := 2; i <= 999999; i++ {
		if tmp[i] == 0 {
			primes[i] = true
			for j := i; j <= 999999; j += i {
				tmp[j] = 1
			}
		}
	}
	return primes
}

// x와 y를 이어붙인 수가 소수인지 판별
func calc(x, y int) int {
	tmp := y
	for tmp > 0 {
		tmp /= 10
		x *= 10
	}

	if primes[x+y] {
		return 1
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
