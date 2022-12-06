package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner     = bufio.NewScanner(os.Stdin)
	writer      = bufio.NewWriter(os.Stdout)
	N           int
	computers   []Computer // 컴퓨터 정보
	rankToIndex [][]int    // 각 계급에 해당하는 컴퓨터의 인데스들
	dp          []int      // 임무 수행 시간 메모이제이션
	maxRank     int        // 최고 계급
	ans         int        // 임무 수행 시간
)

type Computer struct {
	rank  int // 계급
	speed int // 동작 속도
}

// 난이도: Gold 4
// 메모리: 924KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍, 위상 정렬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
	computers = make([]Computer, N+1)
	rankToIndex = make([][]int, N+1)
	dp = make([]int, N+1)

	for i := 1; i <= N; i++ {
		r, t := scanInt(), scanInt()
		computers[i] = Computer{r, t}
		rankToIndex[r] = append(rankToIndex[r], i)
		maxRank = max(maxRank, r)
	}
}

func Solve() {
	for i := 1; i <= maxRank; i++ {
		// i 계급에 해당하는 컴퓨터들 임무 수행
		for len(rankToIndex[i]) > 0 {
			curr := rankToIndex[i][0] // i 계급에 해당하는 컴퓨터의 인덱스
			rankToIndex[i] = rankToIndex[i][1:]

			// 1. i가 최대 계급이면 동작만하고 정보 전달 x
			if i == maxRank {
				dp[curr] = max(dp[curr], dp[curr]+computers[curr].speed)
				ans = max(ans, dp[curr]) // 전체 임무 수행 시간 갱신
				continue
			}

			// 2. i가 최대 계급이 아닌 경우는 다음 계급의 컴퓨터로 정보 전달
			for _, next := range rankToIndex[i+1] {
				// next: i+1 계급에 해당하는 컴퓨터의 인덱스
				// next 컴퓨터는 모든 i 계급의 컴퓨터로부터 정보를 전달받아야 하므로
				// next 컴퓨터가 동작을 시작하는 시간은 i 계급의 컴퓨터가 동작 및 i+1 계급의 컴퓨터로 정보 전달을 마치는 시간의 최댓값
				dp[next] = max(dp[next], dp[curr]+(next-curr)*(next-curr)+computers[curr].speed)
			}
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
