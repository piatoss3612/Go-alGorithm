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
	N, M, K int
	road    [101][101]int // (0,0)에서 (N,M)으로 가는 경로

	// (i,j)에서 (i, j+1)로 가는 도로가 공사중인 경우: blocked[i][j][0] = 1
	// (i,j)에서 (i+1, j)로 가는 도로가 공사중인 경우: blocked[i][j][1] = 1
	blocked [101][101][2]int
)

// 메모리: 1180KB
// 시간: 4ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	K = scanInt()

	// 1. 공사중인 도로의 정보 입력
	for i := 1; i <= K; i++ {
		a, b, c, d := scanInt(), scanInt(), scanInt(), scanInt()
		// 1-1. 좌표 정렬
		if a > c {
			a, c = c, a
		} else if b > d {
			b, d = d, b
		}

		// 1-2. 공사중인 도로 정보 저장
		if a == c {
			blocked[a][b][0] = 1 // x축으로 이동하는 경로가 막힌 경우
		} else if b == d {
			blocked[a][b][1] = 1 // y축으로 이동하는 경로가 막힌 경우
		}

	}

	road[0][0] = 1 // 시작 경로

	// 2. 경로 탐색
	for i := 0; i <= N; i++ {
		for j := 0; j <= M; j++ {
			// 2-1. (i, j-1)이 유효한 좌표이며, (i, j-1)에서 (i, j) 경로가 공사중이 아닌 경우
			if j >= 1 && !(blocked[i][j-1][0] == 1) {
				road[i][j] += road[i][j-1]
			}

			// 2-2. (i-1, j)이 유효한 좌표이며, (i-1, j)에서 (i, j) 경로가 공사중이 아닌 경우
			if i >= 1 && !(blocked[i-1][j][1] == 1) {
				road[i][j] += road[i-1][j]
			}
		}
	}

	fmt.Fprintln(writer, road[N][M]) // 가능한 경로의 수 출력
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
