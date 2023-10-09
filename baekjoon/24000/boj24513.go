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
	N, M     int
	vills    [1001][1001]int  // 마을의 상태
	perfect  [1001][1001]bool // 마을이 이미 완전히 감염되었는지 여부
	infected []Infected       // 감염된 마을들의 정보
	dy       = []int{-1, +0, +1, +0}
	dx       = []int{+0, +1, +0, -1}
)

// 감염된 마을의 정보
type Infected struct {
	y, x, n int
}

// 난이도: Gold 3
// 메모리: 50644KB
// 시간: 236ms
// 분류: 너비 우선 탐색, 시뮬레이션
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			vills[i][j] = scanInt()

			// 바이러스 1 또는 2에 감염된 마을인 경우
			if vills[i][j] > 0 {
				infected = append(infected, Infected{i, j, vills[i][j]}) // 감염된 마을들의 정보를 저장
				perfect[i][j] = true                                     // 해당 마을들은 완전히 감염된 상태로 시작
			}
		}
	}
}

func Solve() {
	var A, B, C int // 각 1, 2, 3번 바이러스에 감염된 마을의 수

	// 1시간동안 바이러스가 퍼지는 과정을 반복
	for len(infected) > 0 {
		candidates := []Infected{} // 1시간동안 바이러스가 퍼진 후 감염된 마을들의 정보

		// 1시간동안 바이러스가 퍼지는 과정
		for len(infected) > 0 {
			here := infected[0] // 감염된 마을 정보
			infected = infected[1:]

			// 감염된 마을의 바이러스 종류 카운팅
			switch here.n {
			case 1:
				A++
			case 2:
				B++
			case 3:
				C++
			}

			// 3번 바이러스는 전염 x
			if here.n == 3 {
				continue
			}

			// 4방향으로 바이러스 전염
			for i := 0; i < 4; i++ {
				ny, nx := here.y+dy[i], here.x+dx[i]
				if !validCoord(ny, nx) {
					continue
				}

				switch {
				case vills[ny][nx] == -1: // 백신이 설치된 마을
				case vills[ny][nx] == 0: // 감염되지 않은 마을
					vills[ny][nx] = here.n
					candidates = append(candidates, Infected{ny, nx, here.n})
				case vills[ny][nx]+here.n == 3 && !perfect[ny][nx]: // 다른 바이러스에 감염되었으며 완전히 감염되지 않은 마을
					vills[ny][nx] = 3
					candidates = append(candidates, Infected{ny, nx, 3})
				}
			}
		}

		for _, c := range candidates {
			// 바이러스 정보가 일치하지 않은 경우
			if c.n != vills[c.y][c.x] {
				continue
			}
			perfect[c.y][c.x] = true       // 완전히 감염된 마을로 분류
			infected = append(infected, c) // 다음 1시간동안 바이러스가 퍼지는 과정에 포함
		}
	}

	fmt.Fprintln(writer, A, B, C) // 각 1, 2, 3번 바이러스에 감염된 마을의 수 출력
}

func validCoord(y, x int) bool {
	return y >= 1 && y <= N && x >= 1 && x <= M
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
