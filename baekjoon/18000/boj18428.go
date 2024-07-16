package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	N        int
	corridor [7][7]byte
	blanks   [][2]int
	teachers [][2]int
	dx       = [4]int{-1, 0, 1, 0}
	dy       = [4]int{0, 1, 0, -1}
)

// 18428번: 감시 피하기
// hhttps://www.acmicpc.net/problem/18428
// 난이도: 골드 5
// 메모리: 876 KB
// 시간: 4 ms
// 분류: 브루트포스 알고리즘, 백트래킹
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			corridor[i][j] = scanBytes()[0]
			// 빈칸 위치 저장
			if corridor[i][j] == 'X' {
				blanks = append(blanks, [2]int{i, j})
			}

			// 선생님 위치 저장
			if corridor[i][j] == 'T' {
				teachers = append(teachers, [2]int{i, j})
			}
		}
	}
}

func Solve() {
	result := installObstacle(0, -1)
	if result {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}

func installObstacle(count, idx int) bool {
	// 3개의 장애물을 설치한 경우
	if count == 3 {
		// 모든 선생님에 대해 실행
		for _, t := range teachers {
			// 선생님의 현재 위치
			x, y := t[0], t[1]

			// 상, 우, 하, 좌 각 방향으로 진행
			for i := 0; i < 4; i++ {
				nx, ny := x+dx[i], y+dy[i]

				// 복도 범위 내에서만 진행
				for inRange(nx, ny) {
					// 장애물을 만난 경우
					if corridor[nx][ny] == 'O' {
						break
					}

					// 학생을 찾은 경우
					if corridor[nx][ny] == 'S' {
						return false
					}

					// 한 방향으로 계속 이동
					nx, ny = nx+dx[i], ny+dy[i]
				}
			}
		}

		return true // 모든 선생님이 상하좌우 어떤 방향에서도 학생을 발견할 수 없는 경우
	}

	// 이전에 idx를 방문했으므로, idx+1부터 진행
	for i := idx + 1; i < len(blanks); i++ {
		// i번째 빈칸의 위치
		x, y := blanks[i][0], blanks[i][1]

		if corridor[x][y] == 'X' {
			// 장애물 설치
			corridor[x][y] = 'O'

			// 모든 학생들이 선생님의 감시로부터 피할 수 있는 경우
			if installObstacle(count+1, i) {
				return true
			}

			// 설치한 장애물 제거
			corridor[x][y] = 'X'
		}
	}

	return false
}

func inRange(x, y int) bool {
	return x >= 1 && x <= N && y >= 1 && y <= N
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
