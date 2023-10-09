package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, k    int
	inp     [201][3]int // inp[i][1]: 왼쪽 방의 가치, inp[i][2]: 오른쪽 방의 가치, inp[i][0]: 양쪽 방의 가치의 합
	// dp[i][j][k]: i번째 세로줄에서 닫아야 하는 방이 j개 남았을 때
	// k = 1: 왼쪽 방을 닫았을 때 공개된 방들의 가치의 합의 최댓값
	// k = 2: 오른쪽 방을 닫았을 때 공개된 방들의 가치의 합의 최댓값
	// k = 0: 양쪽 방을 닫지 않았을 때 공개된 방들의 가치의 합의 최댓값
	dp [201][201][3]int
)

// 메모리: 1888KB
// 시간: 4ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanLines)
	N, k = scanLine()

	// 1. 각 방의 가치를 입력
	for i := 1; i <= N; i++ {
		inp[i][1], inp[i][2] = scanLine()
		inp[i][0] = inp[i][1] + inp[i][2] // 양쪽 방의 가치의 합을 i행 0번 인덱스에 저장
	}

	_, _ = scanLine() // 의미없는 마지막 입력 0, 0

	// 2. dp 초기화, dp값은 항상 0 이상이므로 0이 아닌 값으로 초기화해야 한다
	for i := 0; i <= N; i++ {
		for j := 0; j <= k; j++ {
			dp[i][j] = [3]int{-1, -1, -1}
		}
	}

	// 3. 재귀 함수로 공개된 방들의 가치의 합의 최댓값 찾기
	ans := solve(0, k, 0)
	fmt.Fprintln(writer, ans)
}

// row: 미술관의 row (0 <= row <= N) 번째 세로줄을 방문했을 때
// remain: 비용 절감을 위해 닫아야 하는 방의 개수가 remain (0 <= remain <= k) 만큼 남아있고
// now: 현재 row 번째 세로줄의 now (0 <= now <= 2) 번 방이 닫혀있을 때의 최댓값을 구한다
func solve(row, remain, now int) int {
	// 1. 기저 사례1: 남아있는 세로줄의 개수가 닫아야 하는 방의 개수보다 적은 경우
	// 절대 remain 개의 방을 닫을 수 없으므로 해답이 될 수 없는 아주 작은 값을 반환
	if remain > N-row {
		return -987654321
	}

	// 2. 기저 사례2: 모든 세로줄을 탐색하고 닫아야 하는 방의 개수가 남아있지 않을 때
	// 0을 반환하고 종료
	if row == N && remain == 0 {
		return 0
	}

	ret := &dp[row][remain][now]

	if *ret != -1 {
		return *ret
	}

	// 3. row 번째 세로줄에서 방이 닫혀있는 상태 now에 따라 분기처리
	switch now {

	// 3-1. 모든 방이 열려 있을 때
	case 0:
		// 3-1-1. 다음 세로줄에 있는 방을 닫는 경우
		if remain > 0 {
			// 모든 방이 열려 있는 상태이므로 다음 세로줄의 어떤 방을 닫아도 상관이 없다
			*ret = max(*ret, solve(row+1, remain-1, 1)+inp[row+1][2])
			*ret = max(*ret, solve(row+1, remain-1, 2)+inp[row+1][1])
		}
		// 3-1-2. 다음 세로줄에 있는 방을 닫지 않는 경우
		*ret = max(*ret, solve(row+1, remain, 0)+inp[row+1][0])

	// 3-2. 왼쪽 방이 닫혀 있을 때
	case 1:
		// 3-2-1. 다음 세로줄에 있는 방을 닫는 경우
		if remain > 0 {
			// 같은 열에 있는 방만 닫을 수 있으므로 다음 세로줄의 왼쪽 방을 닫는 경우만 탐색
			*ret = max(*ret, solve(row+1, remain-1, 1)+inp[row+1][2])
		}
		// 3-2-2. 다음 세로줄에 있는 방을 닫지 않는 경우
		*ret = max(*ret, solve(row+1, remain, 0)+inp[row+1][0])

	// 3-3. 오른쪽 방이 닫혀 있을 때
	case 2:
		// 3-3-1. 다음 세로줄에 있는 방을 닫는 경우
		if remain > 0 {
			// 같은 열에 있는 방만 닫을 수 있으므로 다음 세로줄의 오른쪽 방을 닫는 경우만 탐색
			*ret = max(*ret, solve(row+1, remain-1, 2)+inp[row+1][1])
		}
		// 3-3-2. 다음 세로줄에 있는 방을 닫지 않는 경우
		*ret = max(*ret, solve(row+1, remain, 0)+inp[row+1][0])
	}

	return *ret // 모든 경우의 최댓값을 반환
}

func scanLine() (int, int) {
	scanner.Scan()
	line := strings.Fields(scanner.Text())
	x, _ := strconv.Atoi(line[0])
	y, _ := strconv.Atoi(line[1])
	return x, y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
