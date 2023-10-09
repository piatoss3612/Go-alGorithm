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
	N, M      int
	graph     [1001][1001]byte // 구사과의 이동 패턴이 쓰여진 지도
	visited   [1001][1001]int  // 방문 여부를 체크하면서 사이클 번호 부여
	direction = map[byte]int{'N': 0, 'S': 1, 'W': 2, 'E': 3}
	dy        = []int{-1, 1, 0, 0} // y축 이동: direction과 매핑
	dx        = []int{0, 0, -1, 1} // x축 이동: direction과 매핑
	cycle     int                  // 발견한 사이클의 번호
	ans       int                  // 발견한 사이클의 개수
)

// 난이도: Gold 2
// 메모리: 17532KB
// 시간: 72ms
// 분류: 깊이 우선 탐색, 분리 집합
// 구사과는 지도 범위 내에서만 움직이므로 사이클을 형성하면서 서로 분리되어 있는 이동 경로의 개수를 찾아야 한다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		bytes := scanBytes()
		for j := 1; j <= M; j++ {
			graph[i][j] = bytes[j-1]
		}
	}
}

func Solve() {
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			// (i,j)를 아직 방문하지 않은 경우
			if visited[i][j] == 0 {
				cycle++   // 사이클 번호 증가
				DFS(i, j) // 깊이 우선 탐색 실행
			}
		}
	}
	fmt.Fprintln(writer, ans)
}

func DFS(y, x int) {
	visited[y][x] = cycle            // 방문 확인 및 사이클 번호 부여
	next := direction[graph[y][x]]   // 다음 이동할 방향을 정수로 변환한 값
	ny, nx := y+dy[next], x+dx[next] // 다음 이동할 좌표

	// '지도에 쓰여 있는대로 이동했을 때, 지도를 벗어나는 경우는 없다.'
	// 라고 문제에 쓰여 있지만 불안하니 유효한 좌표인지 체크해준다
	if !valid(ny, nx) {
		return
	}

	// (ny, nx) 방문 여부 확인
	switch visited[ny][nx] {
	// 1. 아직 방문하지 않은 경우
	case 0:
		DFS(ny, nx) // 깊이 우선 탐색을 이어간다

	// 2. 이미 방문했는데 사이클 번호가 동일한 경우
	case cycle:
		ans++ // 완전한 사이클을 발견했으므로 사이클 개수를 증가시키고 깊이 우선 탐색 종료
	}
	// 3. 이미 방문했는데 사이클 번호가 다른 경우
	// 현재 탐색을 진행중이던 사이클과 기존에 있던 사이클이 합쳐지는 경우이므로
	// 깊이 우선 탐색을 종료한다
}

func valid(y, x int) bool {
	if y >= 1 && x >= 1 && y <= N && x <= M {
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
