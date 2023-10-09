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
	N, K    int
	point   [501]*Point   // 체크포인트 좌표
	dp      [501][501]int // dp[i][j]: j개의 좌표를 건너뛴 상태에서 i부터 N까지 달릴 수 있는 최소 거리
)

type Point struct {
	x, y int
}

const INF = 987654321

// 메모리: 4908KB
// 시간: 84ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, K = scanInt(), scanInt()
	// 체크포인트 좌표 입력
	for i := 1; i <= N; i++ {
		point[i] = &Point{scanInt(), scanInt()}
	}

	// 재귀 호출: 어떤 체크포인트도 건너뛰지 않은 상태에서 1번 체크포인트부터 시작
	ans := solve(1, 0)
	fmt.Fprintln(writer, ans)
}

func solve(pos, skip int) int {
	// 기저 사례: N번 체크포인트에 도착한 경우
	if pos == N {
		return 0
	}

	// 메모이제이션
	ret := &dp[pos][skip]
	if *ret != 0 {
		return *ret
	}

	*ret = INF // 최솟값 비교를 위해 INF값으로 초기화

	// (0 <= i <= K-skip: 이미 skip개의 체크포인트를 건너뛰고 남은, 건너뛸 수 있는 체크포인트의 수)
	for i := 0; i <= K-skip; i++ {
		next := pos + i + 1 // 현재 위치 pos에서 i개의 체크포인트를 건너뛰고 도달한 체크포인트 번호

		// 건너뛰고 이동한 체크포인트 번호가 N을 벗어난 경우
		if next > N {
			break
		}
		// next에서부터의 최솟값 + pos에서 next까지의 택시 거리의 최솟값 비교
		*ret = min(*ret, solve(next, skip+i)+dist(point[pos], point[next]))
	}
	return *ret
}

func dist(p1, p2 *Point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
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
