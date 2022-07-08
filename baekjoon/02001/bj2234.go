package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, M    int
	castle  [51][51]int // 동서남북 어디에 벽이 있는지 비트마스크로 표현된 값 저장
	visited [51][51]int // 방문여부를 표시하면서 방의 번호를 구분
)

// 전처리: 방문 여부를 확인하기 위해 visited 슬라이스를 -1로 초기화
func init() {
	for i := 1; i <= 50; i++ {
		for j := 1; j <= 50; j++ {
			visited[i][j] = -1
		}
	}
}

// 메모리: 1076KB
// 시간: 4ms
// 너비 우선 탐색, 비트마스크 문제
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()
	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			castle[i][j] = scanInt()
		}
	}

	// 현재 방번호와 sizes 슬라이스의 인덱스를 동기화해준다
	rooms := 0
	sizes := []int{}

	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			// i,j를 아직 방문하지 않은 경우
			if visited[i][j] == -1 {
				// i,j에서 깊이 우선 탐색, 방 번호를 매기고 방 넓이를 반환받아 sizes 슬라이스에 추가
				visited[i][j] = rooms
				sizes = append(sizes, BFS(i, j, rooms))
				rooms++
			}
		}
	}

	// 벽을 하나 부순 경우 2개의 방 넓이의 합의 최댓값
	twoRooms := 0

	// 1,1에서 시작해서 M,N으로 내려가면서 동서남북 다 따질 필요 없이
	// 동쪽, 남쪽만 탐색
	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			// 동쪽으로 1칸 이동하면 방번호가 달라지는 경우
			if Valid(i+1, j) && visited[i][j] != visited[i+1][j] {
				twoRooms = max(twoRooms, sizes[visited[i][j]]+sizes[visited[i+1][j]])
			}

			// 남쪽으로 1칸 이동하면 방번호가 달라지는 경우
			if Valid(i, j+1) && visited[i][j] != visited[i][j+1] {
				twoRooms = max(twoRooms, sizes[visited[i][j]]+sizes[visited[i][j+1]])
			}
		}
	}

	sort.Ints(sizes)

	fmt.Fprintln(writer, rooms)               // 방의 개수
	fmt.Fprintln(writer, sizes[len(sizes)-1]) // 방 크기 최댓값
	fmt.Fprintln(writer, twoRooms)            // 벽 하나를 부순 경우의 가장 넓은 방의 크기
}

type pos struct {
	y, x int
}

func BFS(y, x, roomNo int) int {
	q := []pos{{y, x}}
	cnt := 1

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		// 비트마스크로 벽의 정보가 저장된 값을 15(동서남북 모두 막혀있는 경우)와 xor 연산
		// compass는 이동할 수 있는 방향 정보의 비트마스크
		compass := castle[p.y][p.x] ^ 15

		for compass > 0 {
			direction := compass & -compass // 최소 비트값 추출

			// 1: 서쪽, 2: 북쪽, 4: 동쪽, 8: 남쪽
			// 최소 비트값에 따른 분기 처리
			switch direction {
			case 1:
				ny, nx := p.y, p.x-1
				if Valid(ny, nx) && visited[ny][nx] == -1 {
					visited[ny][nx] = roomNo
					q = append(q, pos{ny, nx})
					cnt++
				}
			case 2:
				ny, nx := p.y-1, p.x
				if Valid(ny, nx) && visited[ny][nx] == -1 {
					visited[ny][nx] = roomNo
					q = append(q, pos{ny, nx})
					cnt++
				}
			case 4:
				ny, nx := p.y, p.x+1
				if Valid(ny, nx) && visited[ny][nx] == -1 {
					visited[ny][nx] = roomNo
					q = append(q, pos{ny, nx})
					cnt++
				}
			case 8:
				ny, nx := p.y+1, p.x
				if Valid(ny, nx) && visited[ny][nx] == -1 {
					visited[ny][nx] = roomNo
					q = append(q, pos{ny, nx})
					cnt++
				}
			}
			compass &= compass - 1 // compass에서 최소 비트 제거
		}
	}
	return cnt
}

func Valid(y, x int) bool {
	if y >= 1 && y <= M && x >= 1 && x <= N {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
