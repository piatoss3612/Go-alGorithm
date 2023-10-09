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
	caffeine [101]int // 각 커피의 카페인 수치

	// dp[i][j][k]: i번째 커피를 마시거나(j=1) 안마셨을 때(j=0), 총 k만큼의 카페인을 섭취한 상태
	dp [101][2][100001]int
)

const INF = 987654321

// 난이도: Gold 5
// 메모리: 83800KB
// 시간: 124ms
// 분류: 다이나믹 프로그래밍

func main() {
	defer writer.Flush()
	Input()
	Solve()
}

func Input() {
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		caffeine[i] = scanInt()
	}
}

func Solve() {
	ans := rec(0, 0, 0)
	if ans == INF {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func rec(idx, drink, total int) int {
	// 기저 사례: 정확히 K만큼의 카페인을 섭취한 경우
	if total == K {
		return 0
	}

	// 기저 사례2: 조건을 만족하지 못한 경우
	if idx == N {
		return INF
	}

	ret := &dp[idx][drink][total]
	if *ret != 0 {
		return *ret
	}

	*ret = INF // 최솟값 비교를 위해 INF로 초기화

	// idx+1번째 커피를 마시지 않는 경우
	*ret = min(*ret, rec(idx+1, 0, total))

	// idx+1번째 커피를 마시는 경우: idx+1번째 커피를 마셨을 때 총 카페인 양이 K보다 작거나 같아야 한다
	if total+caffeine[idx+1] <= K {
		*ret = min(*ret, rec(idx+1, 1, total+caffeine[idx+1])+1)
	}
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
