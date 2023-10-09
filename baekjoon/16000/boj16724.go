package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner   = bufio.NewScanner(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	graph     [1001][1001]byte
	visited   [1001][1001]int
	direction = map[byte]int{'U': 0, 'D': 1, 'L': 2, 'R': 3}
	dy        = []int{-1, 1, 0, 0}
	dx        = []int{0, 0, -1, 1}
	N, M      int
	ans       = 0
)

// 전처리: 방문 여부를 확인하기 위해 visited 배열을 -1로 초기화
func init() {
	for i := 0; i <= 1000; i++ {
		for j := 0; j <= 1000; j++ {
			visited[i][j] = -1
		}
	}
}

// 메모리: 9704KB
// 시간: 68ms
// 깊이 우선 탐색, 그래프 내에 형성되는 사이클의 개수 = SAFE ZONE의 최소 개수
// 문제 조건에 '지도 밖으로 나가는 방향의 입력은 주어지지 않는다.'라고 분명 적혀 있는데
// 왜 IndexOutOfRange 에러가 나냐 염병할 놈들아??????????????
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()

	// 런타임 에러가 여기서 나는 줄 알고 한참 헤맸다 :)
	for i := 1; i <= N; i++ {
		temp := scanBytes()
		for j := 1; j <= M; j++ {
			graph[i][j] = temp[j-1] // 1번 인덱스부터 탐색할 수 있도록 값 복사
		}
	}

	cycle := 0 // 사이클 번호

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			// 아직 방문하지 않았다면
			if visited[i][j] == -1 {
				// i,j에서 깊이 우선 탐색 시작
				DFS(i, j, cycle)
				cycle++
			}
		}
	}

	fmt.Fprintln(writer, ans)
}

func DFS(y, x, cycle int) {
	visited[y][x] = cycle                                      // y,x 방문처리
	currrentDirection := direction[graph[y][x]]                // y,x의 방향에 해당하는 dy, dx의 인덱스 번호
	ny, nx := y+dy[currrentDirection], x+dx[currrentDirection] // 이동할 좌표

	// 좌표가 유효하지 않은 경우
	if !valid(ny, nx) {
		return
	}

	// 이동할 좌표의 사이클 번호가 현재 진행중인 사이클 번호와 동일하다면
	if visited[ny][nx] == cycle {
		ans++ // 하나의 사이클이 형성됬으므로 ans를 1증가
		return
	}

	// 아직 방문하지 않은 좌표라면
	if visited[ny][nx] == -1 {
		DFS(ny, nx, cycle)
	}

	// 이동할 좌표의 visited 값이 -1이나 동일한 사이클 번호가 아닌 경우는
	// 현재 진행중인 경로가 해당 좌표가 포함된 기존의 사이클과
	// 연결되어 있는 것이므로 탐색을 종료한다
}

// 	런타임 에러 (IndexOutOfRange)가 발생해서 추가했다 ^^
func valid(y, x int) bool {
	if y >= 1 && y <= N && x >= 1 && x <= M {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
