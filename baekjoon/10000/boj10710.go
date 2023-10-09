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
	dist    [1001]int // dist[i]: i-1 도시에서 i 도시까지의 거리
	weather [1001]int // weather[j]: j일에 날씨의 나쁨 정도

	// dp[i][j]: 0번 도시에서 출발한지 i일이 지난 시점에서 j도시에서 출발할 때,
	// M일 이내로 N번 도시에 도착하는 피로도의 총합의 최솟값을 메모이제이션
	dp [1001][1001]int
)

const INF = 987654321 // 최솟값 비교를 위한 상수

// 난이도: Gold 5
// 메모리: 14768KB
// 시간: 16ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		dist[i] = scanInt()
	}
	for i := 1; i <= M; i++ {
		weather[i] = scanInt()
	}

	ans := solve(1, 0)
	fmt.Fprintln(writer, ans)
}

func solve(day, pos int) int {
	// 기저 사례: M일이 지났음에도 N번 도시에 도착하지 못한 경우
	if day > M && pos != N {
		return INF
	}

	// 기저 사례: M일 이내로 N번 도시에 도착한 경우
	if day <= M+1 && pos == N {
		return 0
	}

	ret := &dp[day][pos]
	if *ret != 0 {
		return *ret
	}

	*ret = INF
	*ret = min(*ret, solve(day+1, pos))                            // pos번 도시에서 대기
	*ret = min(*ret, solve(day+1, pos+1)+weather[day]*dist[pos+1]) // pos+1번 도시로 이동

	return *ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
