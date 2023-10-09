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
	N       int
	cost    [17][17]float64
	dp      [17][1 << 17]float64
)

const INF = 987654321

// 난이도: Gold 1
// 메모리: 9176KB
// 시간: 48ms
// 분류: 다이나믹 프로그래밍, 비트마스킹, 외판원 순회 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func dist(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

// *1번부터 N번까지 번호가 매겨져 있는 도시들이 있고, 모든 도시 사이에는 길이 있다.*
func Setup() {
	N = scanInt()
	coords := make([][2]int, N) // 각 도시의 좌표 정보

	// 도시의 좌표 정보를 이용해 도시 사이의 거리를 구한다
	for i := 0; i < N; i++ {
		coords[i][0], coords[i][1] = scanInt(), scanInt()
		for j := 0; j < i; j++ {
			iToJ := dist(coords[i][0], coords[i][1], coords[j][0], coords[j][1])
			cost[i][j] = iToJ
			cost[j][i] = iToJ
		}
	}

	// dp를 -1로 초기화
	for i := 0; i < N; i++ {
		for j := 0; j < 1<<N; j++ {
			dp[i][j] = -1
		}
	}
}

func Solve() {
	// 모든 도시가 연결되어 있으므로 어디에서 순회를 시작하던 상관없다

	ans := rec(0, 1) // 0번 도시에서 시작하여 모든 도시를 한 번씩 방문하고 0번 도시로 돌아오는 최솟값
	fmt.Fprintln(writer, ans)
}

// curr: 현재 위치
// visited: 각 도시를 방문했는지 여부를 비트마스킹한 값
// rec: 현재 위치 curr에서 일부 도시를 한 번씩 방문한 상태에서 나머지 도시를 모두 한 번씩 방문하고 시작 도시로 돌아가는 비용의 최솟값
func rec(curr, visited int) float64 {
	// 기저 사례: 모든 도시를 방문한 경우
	if visited == (1<<N)-1 {
		return cost[curr][0]
	}

	// 기저 사례2: 이미 최솟값을 구한 경우
	ret := &dp[curr][visited]
	if *ret != -1 {
		return *ret
	}

	*ret = INF

	for i := 0; i < N; i++ {
		// i번 도시를 아직 방문하지 않은 경우
		if (visited & (1 << i)) == 0 {
			// curr에서 i번 도시를 방문했을 경우의 비용의 최솟값과 현재 최솟값 비교
			*ret = min(
				*ret,
				rec(i, visited+(1<<i))+cost[curr][i],
			)
		}
	}

	return *ret
}

func min(a, b float64) float64 {
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
