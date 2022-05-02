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
	n, m    int
	graph   [501][501]int
	dp      [501][501]int
	dx      = []int{0, 0, 1, -1}
	dy      = []int{1, -1, 0, 0}
)

// 일반적인 dfs만으로 풀면 시간 초과 발생
// 메모이제이션을 활용해야 한다!
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			graph[i][j] = scanInt()
			dp[i][j] = -1 // 방문 여부를 확인하기 위해 -1로 초기화
		}
	}

	/*
		예제 입력:
		4 5
		50 45 37 32 30
		35 50 40 20 25
		30 30 25 17 28
		27 24 22 15 10

		dp:
		3  2  2  2  1
		1 -1 -1  1  1
		1 -1 -1  1 -1
		1  1  1  1 -1

		예제 출력:
		3
	*/

	res := rec(1, 1)
	fmt.Fprintln(writer, res)
}

// 재귀 함수 정의
func rec(x, y int) int {
	// 기저 사례: x = n, y = m
	if x == n && y == m {
		return 1
	}

	// dp[x][y]를 한 번이라도 지나간 경우
	if dp[x][y] != -1 {
		return dp[x][y]
	}

	ret := &dp[x][y]
	*ret = 0

	// 상하좌우 및 내리막길인지 검증
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if validCoord(nx, ny) && graph[nx][ny] < graph[x][y] {
			// 내리막길인 경우, 해당 좌표를 재귀 호출한 값을 dp[x][y]에 더해준다
			*ret += rec(nx, ny)
		}
	}

	return *ret
}

func validCoord(x, y int) bool {
	if x >= 1 && y >= 1 && x <= n && y <= m {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
