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
	graph   [101][101]byte // 미로
	dist    [101][101]int  // 부숴야 하는 벽의 최소 개수
	dy      = []int{-1, 0, 1, 0}
	dx      = []int{0, -1, 0, 1}
	M, N    int
	INF     = 987654321
)

// 다익스트라 탐색을 위한 그래프 초기화
func init() {
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			dist[i][j] = 987654321
		}
	}
}

// 메모리: 3852KB
// 시간: 12ms
// 다익스트라 + BFS
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	M, N = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		line := scanBytes()
		for j := 1; j <= M; j++ {
			graph[i][j] = line[j-1]
		}
	}

	Dijkstra() // 다익스트라 탐색
	fmt.Fprintln(writer, dist[N][M])
}

type move struct {
	y, x int
}

func Dijkstra() {
	dist[1][1] = 0      // 시작 위치 1,1에서 부숴야 하는 벽의 수를 1로 초기화
	q := []move{{1, 1}} // 우선수위 큐를 사용하지 않고 BFS처럼 일반 큐를 사용

	for len(q) > 0 {
		// 1. 큐에서 현재 위치값 꺼내오기
		f := q[0]
		q = q[1:]

		for i := 0; i < 4; i++ {
			// 2. 이동할 위치값 구하기
			ny, nx := f.y+dy[i], f.x+dx[i]
			if !valid(ny, nx) {
				continue
			}

			// 3. 이동할 위치가 벽이라면
			if graph[ny][nx] == '1' {
				// 지금까지 이동하면서 부순 벽의 수와 이동할 위치의 벽을 부순 경우를 더하여
				// 이동할 위치의 dist값과 비교하여 더 작은 값으로 갱신
				if dist[ny][nx] > dist[f.y][f.x]+1 {
					dist[ny][nx] = dist[f.y][f.x] + 1
					q = append(q, move{ny, nx})
				}
				// 4. 이동할 위치가 벽이 아니라면
			} else {
				// 지금까지 이동하면서 부순 벽의 수와 이동할 위치의 dist값을 비교하여 더 작은 값으로 갱신
				if dist[ny][nx] > dist[f.y][f.x] {
					dist[ny][nx] = dist[f.y][f.x]
					q = append(q, move{ny, nx})
				}
			}
		}
	}
}

func valid(y, x int) bool {
	if y >= 1 && y <= N && x >= 1 && x <= M {
		return true
	}
	return false
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
