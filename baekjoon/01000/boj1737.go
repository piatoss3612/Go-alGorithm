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
	N       float64
	dp      map[float64]int // 실수는 배열의 인덱스로 사용할 수 없으므로 map을 사용하여 메모이제이션
)

const MOD = 1000000000000000000

// 메모리: 159064KB
// 시간: 684ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanFloat64()
	dp = make(map[float64]int)

	ans := Pibonacci(N)
	fmt.Fprintln(writer, ans)
}

func Pibonacci(n float64) int {
	// 기저 사례: n이 파이보다 작거나 같은 경우
	if n <= math.Pi {
		return 1
	}

	// dp에 저장된 값인지 확인
	_, ok := dp[n]
	if ok {
		return dp[n]
	}

	// 새로운 dp값 갱신
	dp[n] = (Pibonacci(n-1) + Pibonacci(n-math.Pi)) % MOD
	return dp[n]
}

func scanFloat64() float64 {
	scanner.Scan()
	n, _ := strconv.ParseFloat(scanner.Text(), 64)
	return n
}
