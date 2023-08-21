package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)

	N, K int
	arr  []int
	dp   [1001]int
)

// 난이도: Silver 1
// 메모리: 908KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	arr = make([]int, N)

	for i := 0; i < N; i++ {
		arr[i] = scanInt()
	}

	K = scanInt()

	sort.Ints(arr) // 오름차순 정렬 -> 탐색 횟수 줄이기 위함
}

func Solve() {
	dp[0] = K // 0에서 시작할 때, K개의 수를 가져갈 수 있음

	for i := 1; i <= 1000; i++ {
		dp[i] = -1 // -1: 불가능, 0 이상: 가능

		for j := 0; j < N; j++ {
			if i-arr[j] < 0 {
				break
			}

			dp[i] = max(dp[i], dp[i-arr[j]]-1) // i를 만들기 위해 필요한 최소 개수를 뺀 남은 개수
		}

		// i를 만들 수 없는 경우
		if dp[i] == -1 {
			// i가 짝수면 홀순 승리, 홀수면 짝순 승리
			if i%2 == 0 {
				fmt.Fprintln(writer, "holsoon win at", i)
			} else {
				fmt.Fprintln(writer, "jjaksoon win at", i)
			}
			return
		}
	}

	// 1에서 1000까지 모두 만들 수 있는 경우는 없음
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
