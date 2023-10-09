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
	N, M    int
	dp      [1000001]int // dp[i]: 가방에 i만큼의 무게를 담을 경우, 물건의 가치의 최댓값
)

// 난이도: Gold 4
// 메모리: 8728KB
// 시간: 180ms
// 분류: 다이나믹 프로그래밍, 배낭 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= 1000000; i++ {
		dp[i] = -1 // 최댓값 비교를 위해 -1로 초기화
	}
	// dp[0] = 0
}

func Solve() {
	// 배낭 문제
	var w, v int
	for i := 1; i <= N; i++ {
		w, v = scanInt(), scanInt()
		for j := 1000000; j >= 0; j-- {
			if dp[j] != -1 && j+w <= 1000000 {
				dp[j+w] = max(dp[j+w], dp[j]+v)
			}
		}
	}

	ans := 0
	maxValue, maxBackpack := 0, 0

	for i := 1; i <= M; i++ {
		backpack := scanInt()
		value := 0
		// 가방에 담을 수 있는 물건들의 가치의 합의 최댓값 찾기
		for j := 1; j <= backpack; j++ {
			value = max(value, dp[j])
		}

		if i == 1 {
			maxValue = value
			maxBackpack = backpack
			ans = 1
			continue
		}

		// 효율성 검사:
		// 효율성 비교식 maxValue / maxBackpack < value / backpack을
		// 소수 연산의 오차로 인해 발생하는 문제를 방지하기 위해 정수값으로 비교할 수 있도록 변형하였다
		// [참고] https://dong-gas.tistory.com/29
		if maxValue*backpack < value*maxBackpack {
			maxValue = value
			maxBackpack = backpack
			ans = i
		}
	}

	fmt.Fprintln(writer, ans)
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
