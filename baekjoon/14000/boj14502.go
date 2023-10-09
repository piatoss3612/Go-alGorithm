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
	lab     [9][9]int  // 연구소의 지도
	virus   []Infected // 바이러스가 위치한 지도 상의 좌표
	dy      = []int{-1, +0, +1, +0}
	dx      = []int{+0, +1, +0, -1}
	ans     int // 안전 영역의 최대 크기
)

type Infected struct {
	y, x int
}

// 난이도: Gold 4
// 메모리: 17696KB
// 시간: 296ms
// 분류: 브루트포스, 너비 우선 탐색, 구현
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			lab[i][j] = scanInt()
			// 바이러스의 최초 위치는 항상 동일하므로 따로 저장해 놓았다
			if lab[i][j] == 2 {
				virus = append(virus, Infected{i, j})
			}
		}
	}
}

func Solve() {
	/*
		1. 벽을 3개 설치한다
		2. 바이러스 확산 시뮬레이션을 돌려본다
		3. 안전 영역의 크기를 비교한다
		4. 가능한 모든 경우에 대해 1~3반복
	*/
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if lab[i][j] == 0 {
				lab[i][j] = 1
				installWalls(i, j, 1)
				lab[i][j] = 0
			}
		}
	}
	fmt.Fprintln(writer, ans)
}

// y: 연구소 지도 상의 y좌표
// x: 연구소 지도 상의 x좌표
// walls: 설치한 벽의 개수
// 연구소 지도 상의 좌표 (y,x)에 walls 번째 벽이 설치되어 있고
// 오른쪽 또는 아래로 내려가면서 새로운 벽을 설치한다
func installWalls(y, x, walls int) {
	// 벽을 3개 설치한 경우
	// 바이러스 확산 시뮬레이션 실행
	if walls == 3 {
		defenseVirus()
		return
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			// 중복된 형태로 벽을 설치하는 경우는 배제한다
			if i <= y && j < x {
				continue
			}

			// 벽을 설치할 수 있는 경우
			if lab[i][j] == 0 {
				lab[i][j] = 1
				installWalls(i, j, walls+1)
				lab[i][j] = 0
			}
		}
	}
}

// 깊이 우선 탐색을 실행하여 바이러스가 확산되는 상황을 시뮬레이션한다
func defenseVirus() {
	virtualLab := lab // 현재 연구소의 지도를 복사

	attack := []Infected{}
	attack = append(attack, virus...) // 바이러스의 좌표 복사

	// 깊이 우선 탐색 실행
	for len(attack) > 0 {
		next := attack[0]
		attack = attack[1:]

		// 감연된 위치로부터 상하좌우 탐색
		for i := 0; i < 4; i++ {
			ny, nx := next.y+dy[i], next.x+dx[i]
			// (ny, nx)가 범위 내에 있고 안전 영역인 경우
			if inRange(ny, nx) && virtualLab[ny][nx] == 0 {
				virtualLab[ny][nx] = 2                    // 안전 영역이 감염된다
				attack = append(attack, Infected{ny, nx}) // 감염된 안전영역은 다시 바이러스를 전파한다
			}
		}
	}

	// 안전영역의 크기 구하기
	cnt := 0
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if virtualLab[i][j] == 0 {
				cnt++
			}
		}
	}

	ans = max(ans, cnt) // 안전영역의 크기의 최댓값 갱신
}

func inRange(y, x int) bool {
	if y >= 1 && y <= N && x >= 1 && x <= M {
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
