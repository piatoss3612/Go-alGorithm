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
	galaxy  [501][501]byte
	PR, PC  int
	dy      = [4]int{-1, 0, 1, 0} // U, R, D, L
	dx      = [4]int{0, 1, 0, -1}
)

const INF = 987654321

// 3987번: 보이저 1호
// hhttps://www.acmicpc.net/problem/3987
// 난이도: 골드 5
// 메모리: 1108 KB
// 시간: 16 ms
// 분류: 그래프 탐색, 시뮬레이션
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		b := scanBytes()
		for j := 1; j <= M; j++ {
			galaxy[i][j] = b[j-1]
		}
	}
	PR, PC = scanInt(), scanInt()
}

func Solve() {
	ans := 0
	maxDir := 0

	for i := 0; i < 4; i++ {
		cnt := sendSignal(i)
		if cnt > ans {
			ans = cnt
			maxDir = i
		}
	}

	switch maxDir {
	case 0:
		fmt.Fprintln(writer, "U")
	case 1:
		fmt.Fprintln(writer, "R")
	case 2:
		fmt.Fprintln(writer, "D")
	case 3:
		fmt.Fprintln(writer, "L")
	}

	if ans == INF {
		fmt.Fprintln(writer, "Voyager")
	} else {
		fmt.Fprintln(writer, ans)
	}
}

func sendSignal(dir int) int {
	cnt := 0
	r, c := PR, PC

	for {
		// cnt가 N*M*4보다 크면 무한 루프로 판단 (방향의 가짓수 * 항성계 크기를 모두 돌아도 블랙홀을 만나지 못한 경우)
		if cnt > N*M*4 {
			return INF
		}

		r, c = r+dy[dir], c+dx[dir]

		// 신호를 보낸 횟수 증가
		cnt++

		// 항성계 밖으로 나가면 종료
		if !inRange(r, c) {
			break
		}

		// 블랙홀을 만나면 종료
		if galaxy[r][c] == 'C' {
			break
		}

		// 행성을 만난 경우
		if galaxy[r][c] == '/' {
			// 오른쪽 이동 중이면 위로 방향 전환
			// 왼쪽 이동 중이면 아래로 방향 전환
			// 위로 이동 중이면 오른쪽으로 방향 전환
			// 아래로 이동 중이면 왼쪽으로 방향 전환
			dir = dir ^ 1
		} else if galaxy[r][c] == '\\' {
			// 오른쪽 이동 중이면 아래로 방향 전환
			// 왼쪽 이동 중이면 위로 방향 전환
			// 위로 이동 중이면 왼쪽으로 방향 전환
			// 아래로 이동 중이면 오른쪽으로 방향 전환
			dir = dir ^ 3
		}
	}

	return cnt
}

func inRange(r, c int) bool {
	return 1 <= r && r <= N && 1 <= c && c <= M
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
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
