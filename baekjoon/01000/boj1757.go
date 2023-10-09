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
	run     [10001]int      // i분에 달리는 거리
	dp      [10001][501]int // [달리는 시간][지침 지수]
)

const INF = -987654321

// 메모리: 79764KB
// 시간: 132ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()

	// 달리기를 시작하는 시간을 0으로 두고 0번 인덱스부터 입력을 받는다
	for i := 0; i < N; i++ {
		run[i] = scanInt()
	}

	// 시간 0에서 지침 지수 0으로 시작하여 달릴 수 있는 최대 거리 출력
	fmt.Fprintln(writer, solve(0, 0))
}

func solve(i, tired int) int {
	// 기저 사례: N분동안 달리기를 마친 경우
	if i >= N {
		// 지침 지수가 0인 경우
		if tired == 0 {
			// 공부를 할 수 있는 상태이므로 0 반환
			return 0
		} else {
			// 지침 지수가 0이 아닌 경우
			// 공부를 할 수 없는 상태이므로 INF 반환
			return INF
		}
	}

	ret := &dp[i][tired]

	if *ret != 0 {
		return *ret
	}

	*ret = INF // 최댓값 비교를 위해 -987654321로 초기화

	// 1. i분에 1분 동안 달리는 경우
	// 달리더라도 지침 지수가 M을 초과해서는 안된다
	if tired+1 <= M {
		*ret = max(*ret, solve(i+1, tired+1)+run[i])
	}

	// 2. i분에 쉬기 시작하는 경우
	// 지침 지수가 0보다 큰 경우
	if tired > 0 {
		// 한 번 쉴 때 반드시 지침 지수가 0이될 때까지 쉬어야 한다
		// 따라서 현재 시간 i와 지침 지수를 더한 값이 N보다 작거나 같은 경우에만 쉰다
		if i+tired <= N {
			*ret = max(*ret, solve(i+tired, 0))
		}
	} else {
		// 지침 지수가 0이라면 1분만 쉰다
		*ret = max(*ret, solve(i+1, 0))
	}

	return *ret
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
