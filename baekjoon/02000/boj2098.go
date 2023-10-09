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
	N       int
	cost    [17][17]int
	dp      [17][1 << 17]int
)

const INF = 987654321

// 난이도: Gold 1
// 메모리: 9160KB
// 시간: 40ms
// 분류: 다이나믹 프로그래밍, 비트마스킹, 외판원 순회 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			cost[i][j] = scanInt()
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < 1<<N; j++ {
			dp[i][j] = -1
		}
	}
}

func Solve() {
	ans := rec(0, 1) // 0번 도시가 방문처리된 상태에서 시작
	fmt.Fprintln(writer, ans)
}

// curr: 현재 위치(도시 번호 0~N-1)
// visited: 방문한 도시들을 번호에 따라 비트마스킹한 값
// rec: 현재 위치가 curr이고 일부 도시들은 여행을 마친 상태(visited)에서
// 한 번도 방문하지 않은 도시들을 모두 여행하기 위한 최소 비용을 구한다
func rec(curr, visited int) int {
	// 기저 사례: 모든 도시를 한 번씩 방문한 경우
	if visited == (1<<N)-1 {
		if cost[curr][0] > 0 {
			return cost[curr][0]
		}
		// curr에서 시작 위치로 돌아가는 길이 없는 경우
		return INF
	}

	ret := &dp[curr][visited]
	// 기저 사례2: 이미 curr에서 남은 도시들을 방문하기 위한 최소 비용을 구해 놓은 경우
	if *ret != -1 {
		return *ret
	}

	*ret = INF

	for next := 0; next < N; next++ {
		// next를 아직 방문하지 않았고 curr에서 next로 가는 길이 있는 경우
		// *연산자 우선순위 주의*
		if visited&(1<<next) == 0 && cost[curr][next] != 0 {
			*ret = min(
				*ret,
				rec(next, visited+(1<<next))+cost[curr][next],
			) // 최소 비용 갱신
		}
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
