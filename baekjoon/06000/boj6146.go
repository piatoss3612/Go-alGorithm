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

	X, Y, N int
	board   [1001][1001]int
	visited [1001][1001]bool
	dy      = [4]int{0, 0, 1, -1}
	dx      = [4]int{1, -1, 0, 0}
)

// 난이도: Silver 1
// 메모리: 15920KB
// 시간: 60ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	X, Y, N = scanInt(), scanInt(), scanInt()
	for i := 0; i < N; i++ {
		a, b := scanInt(), scanInt()
		board[a+500][b+500] = 1 // 배열의 인덱스가 음수가 되지 않도록 500씩 더해서 저장
	}

	X, Y = X+500, Y+500 // 마찬가지로 신아의 위치도 500씩 더해서 저장
}

func Solve() {
	q := [][3]int{{500, 500, 0}} // 시작점은 0,0이므로 500,500에서 시작, 이동 횟수는 0
	visited[500][500] = true

	for len(q) > 0 {
		x, y, cnt := q[0][0], q[0][1], q[0][2]
		q = q[1:]

		if x == X && y == Y {
			fmt.Fprintln(writer, cnt)
			return
		}

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx < 0 || nx > 1000 || ny < 0 || ny > 1000 {
				continue
			}
			if visited[nx][ny] || board[nx][ny] == 1 {
				continue
			}
			visited[nx][ny] = true
			q = append(q, [3]int{nx, ny, cnt + 1}) // nx, ny로 이동했으므로 이동횟수 1 증가
		}
	}
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
