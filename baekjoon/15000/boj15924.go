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
	mat     [3001][3001]byte
	dp      [3001][3001]int
)

// 15924번: 욱제는 사과팬이야!!
// hhttps://www.acmicpc.net/problem/15924
// 난이도: 골드 5
// 메모리: 84764 KB
// 시간: 584 ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		s := scanString()
		for j := 1; j <= M; j++ {
			mat[i][j] = s[j-1] // 인덱스 주의
			dp[i][j] = -1      // 초기화
		}
	}
}

func Solve() {
	fmt.Fprintln(writer, rec(N, M)) // (N, M)에서 시작하여 역으로 이동 가능한 경로의 수를 구한다.
}

func rec(x, y int) int {
	ret := &dp[x][y]

	// 이미 계산된 값이 있으면 그 값을 반환
	if *ret != -1 {
		return *ret
	}

	*ret = 1 // (x, y)에서 시작하는 경로의 수는 1로 초기화

	// 1. 왼쪽에서 (x, y)로 이동 가능한 경우
	nx, ny := x, y-1
	if inRange(nx, ny) && (mat[nx][ny] == 'E' || mat[nx][ny] == 'B') {
		*ret += rec(nx, ny)
		*ret %= 1000000009
	}

	// 2. 위쪽에서 (x, y)로 이동 가능한 경우
	nx, ny = x-1, y
	if inRange(nx, ny) && (mat[nx][ny] == 'S' || mat[nx][ny] == 'B') {
		*ret += rec(nx, ny)
		*ret %= 1000000009
	}

	return *ret // (x, y)에서 시작하는 경로의 수를 반환
}

func inRange(x, y int) bool {
	return 1 <= x && x <= N && 1 <= y && y <= M
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
