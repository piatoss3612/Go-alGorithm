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
	sea       [21][21]int // 상어는 바다 생물이다
	dx        = []int{-1, 0, 1, 0}
	dy        = []int{0, -1, 0, 1}
	babyShark Shark // 애기 상어
	N         int   // 바다의 크기 N*N
)

// 애기 상어 정보
type Shark struct {
	size int // 크기
	y, x int // 위치
}

// 메모리: 5956KB
// 시간: 12ms
// 너비 우선 탐색, 시뮬레이션
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N = scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			sea[i][j] = scanInt()
			if sea[i][j] == 9 {
				babyShark.y, babyShark.x = i, j // 애기 상어의 위치 설정
			}
		}
	}

	babyShark.size = 2 // 애기 상어 시작 크기
	totalDistance := 0 // 총 이동한 거리 = 시간
	eaten := 0         // 먹은 물고기 수

	// 전체 배열의 크기의 최댓값이 20*20밖에 안되므로 매번 너비 우선 탐색으로
	// 현재 애기 상어의 위치에서 먹을 수 있는 물고기들을 탐색하고
	// 먹을 수 있는 물고기가 없다면 엄크가 떠서 반복문이 종료된다
	for {
		distance, edible := BFS()

		// 먹을 수 있는 물고기가 없다면
		if !edible {
			break
		}
		totalDistance += distance
		eaten++

		// 먹은 물고기의 수가 애기 상어의 크기와 같아지면
		if babyShark.size == eaten {
			babyShark.size++ // 애기 상어 사이즈업
			eaten = 0        // 먹은 물고기 수 초기화
		}
	}

	fmt.Fprintln(writer, totalDistance)
}

// 애기 상어의 이동 정보
type Hunting struct {
	y, x   int  // 위치
	dist   int  // 이동 거리
	edible bool // y,x에서 물고기를 먹을 수 있는지 여부
}

func BFS() (int, bool) {
	var visited [21][21]bool // 깊이 우선 탐색을 진행할 때마다 방문 여부 초기화

	// 큐에 애기 상어의 현재 위치 추가 및 방문 처리
	q := []Hunting{{babyShark.y, babyShark.x, 0, false}}
	visited[babyShark.y][babyShark.x] = true

	// 만약 애기 상어가 있던 위치의 값이 9라면 다음 번 탐색을 위해 0으로 변경해준다
	sea[babyShark.y][babyShark.x] = 0

	// 엄크 뜨냐?
	mamaShark := true

	// 먹을 수 있는 물고기의 위치와 최단 거리
	y, x, minDistance := 0, 0, 987654321

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		// 애기 상어가 현재 위치에서 물고기를 먹을 수 있다면
		if curr.edible {
			// 엄마 상어의 도움은 필요없다
			mamaShark = false

			// 애기 상어의 원래 위치에서 현재 위치까지의 이동거리가
			// 물고기를 먹을 수 있는 다른 위치보다 짧다면
			if curr.dist < minDistance {
				minDistance = curr.dist
				y, x = curr.y, curr.x
				// 거리가 가까운 물고기가 여러 마리라면
			} else if curr.dist == minDistance {
				// 가장 위에 있는 물고기
				if curr.y < y {
					y, x = curr.y, curr.x
					// 가장 위에 있는 물고기가 여러 마리라면
					// 가장 왼쪽에 있는 물고기
				} else if curr.y == y && curr.x < x {
					y, x = curr.y, curr.x
				}
			}
			// *** 조건에 맞는 물고기의 정확한 위치를 찾는 것이 중요하다
			continue
		}

		// 현재 위치에서 먹을 수 있는 물고기가 없다면
		// 상하좌우 이동가능한 위치 탐색
		for i := 0; i < 4; i++ {
			ny, nx := curr.y+dy[i], curr.x+dx[i]
			if valid(ny, nx) && !visited[ny][nx] && sea[ny][nx] <= babyShark.size {
				visited[ny][nx] = true
				next := Hunting{
					ny, nx, curr.dist + 1, false,
				}
				// 이동할 위치의 값이 0이 아니고 애기 상어의 크기보다 작은 값이라면
				if sea[ny][nx] != 0 && sea[ny][nx] < babyShark.size {
					next.edible = true // 먹을 수 있는 물고기가 있다는 의미
				}
				q = append(q, next)
			}
		}
	}

	// 애기가 엄마를 호출해야 되는 경우
	if mamaShark {
		return 0, false
	}

	// 물고기를 먹었다면
	// 물고기가 있던 위치의 값을 0으로 변경하고
	// 애기 상어의 위치를 갱신해준다
	sea[y][x] = 0
	babyShark.y, babyShark.x = y, x
	return minDistance, true
}

func valid(y, x int) bool {
	if y >= 1 && y <= N && x >= 1 && x <= N {
		return true
	}
	return false
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
