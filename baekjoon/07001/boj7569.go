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
	M, N, H int
	box     [101][101][101]int  // 토마토가 들어있는 박스
	visited [101][101][101]bool // 너비 우선 탐색 방문 여부
	unripe  int                 // 익지 않은 토마토의 개수
	ripen   []Tomato            // 익은 토마토들을 저장하는 큐
	dy      = []int{-1, +0, +1, +0}
	dx      = []int{+0, +1, +0, -1}
	dz      = []int{-1, +1}
)

// 토마토의 좌표 정보를 저장하는 구조체
type Tomato struct {
	z, y, x int
}

// 난이도: Gold 5
// 메모리: 53328KB
// 시간: 264ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	M, N, H = scanInt(), scanInt(), scanInt()
	for z := 1; z <= H; z++ {
		for y := 1; y <= N; y++ {
			for x := 1; x <= M; x++ {
				box[z][y][x] = scanInt()
				if box[z][y][x] == 1 {
					visited[z][y][x] = true
					ripen = append(ripen, Tomato{z, y, x})
				} else if box[z][y][x] == 0 {
					unripe++
				}
			}
		}
	}
}

func Solve() {
	// 이미 모든 토마토가 익은 상태인 경우
	if unripe == 0 {
		fmt.Fprintln(writer, 0)
		return
	}

	ans := 0

	// 모든 토마토가 익을 때까지 며칠이 걸리는지 탐색
	for len(ripen) > 0 && unripe > 0 {
		OneDay()
		ans++
	}

	if unripe > 0 {
		fmt.Fprintln(writer, -1) // 모든 토마토가 익지 않은 경우
	} else {
		fmt.Fprintln(writer, ans) // 모든 토마토가 익은 경우
	}
}

func OneDay() {
	candidates := []Tomato{} // 새롭게 익은 토마토들의 좌표 정보를 저장할 슬라이스

	// 하루동안 큐에 담겨있는 모든 익은 토마토들이 인접한 토마토들에게 주는 영향에 대해 탐색
	for len(ripen) > 0 {
		next := ripen[0]
		ripen = ripen[1:]

		// 박스의 앞, 오른쪽, 뒤, 왼쪽 탐색
		for i := 0; i < 4; i++ {
			ny, nx := next.y+dy[i], next.x+dx[i]
			if valid(ny, nx) && !visited[next.z][ny][nx] {
				if box[next.z][ny][nx] == -1 {
					visited[next.z][ny][nx] = true
				} else if box[next.z][ny][nx] == 0 {
					box[next.z][ny][nx] = 1
					visited[next.z][ny][nx] = true
					candidates = append(candidates, Tomato{next.z, ny, nx})
					unripe--
				}
			}
		}

		// 박스의 위, 아래 탐색
		for i := 0; i < 2; i++ {
			nz := next.z + dz[i]

			if nz >= 1 && nz <= H && !visited[nz][next.y][next.x] {
				if box[nz][next.y][next.x] == -1 {
					visited[nz][next.y][next.x] = true
				} else if box[nz][next.y][next.x] == 0 {
					box[nz][next.y][next.x] = 1
					visited[nz][next.y][next.x] = true
					candidates = append(candidates, Tomato{nz, next.y, next.x})
					unripe--
				}
			}
		}

		// 토마토가 모두 익은 경우
		if unripe == 0 {
			return
		}
	}

	ripen = append(ripen, candidates...) // 새롭게 익은 토마토들은 다음날 인접한 토마토들에게 영향을 미친다
}

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
