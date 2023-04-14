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

	R, C    int
	forest  [51][51]byte
	visited [51][51]bool

	beaver   move
	hedgehog []move
	water    []move

	dy = []int{-1, +0, +1, +0}
	dx = []int{+0, +1, +0, -1}
)

type move struct {
	y, x int
	cnt  int
}

// 난이도: Gold 4
// 메모리: 1212KB
// 시간: 4ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	R, C = scanInt(), scanInt()
	for i := 1; i <= R; i++ {
		row := scanBytes()
		for j, b := range row {
			forest[i][j+1] = b

			switch b {
			case 'D':
				beaver = move{y: i, x: j + 1}
			case 'S':
				visited[i][j+1] = true
				hedgehog = append(hedgehog, move{y: i, x: j + 1})
			case '*':
				visited[i][j+1] = true
				water = append(water, move{y: i, x: j + 1})
			}
		}
	}
}

func Solve() {
	for len(hedgehog) > 0 {
		candidates := []move{} // 물이 찰 예정인 곳

		// 고슴도치가 물이 찰 예정인 곳으로 이동할 수 없으므로 물이 이동하는 과정을 먼저 수행한다.
		for len(water) > 0 {
			w := water[0]
			water = water[1:]

			for i := 0; i < 4; i++ {
				ny, nx := w.y+dy[i], w.x+dx[i]
				if valid(ny, nx) && !visited[ny][nx] {
					// 비어있는 곳이거나 고슴도치가 있던 곳이라면 물이 이동할 수 있다.
					if forest[ny][nx] == '.' || forest[ny][nx] == 'S' {
						visited[ny][nx] = true
						forest[ny][nx] = '*' // 물이 찰 예정인 곳을 표시한다.
						candidates = append(candidates, move{y: ny, x: nx})
					}
				}
			}
		}

		water = append(water, candidates...)

		candidates = []move{} // 고슴도치가 이동할 수 있는 곳

		for len(hedgehog) > 0 {
			h := hedgehog[0]
			hedgehog = hedgehog[1:]

			// 고슴도치가 비버의 굴에 도착했다면 시간을 출력하고 종료한다.
			if h.y == beaver.y && h.x == beaver.x {
				fmt.Fprintln(writer, h.cnt)
				return
			}

			for i := 0; i < 4; i++ {
				ny, nx := h.y+dy[i], h.x+dx[i]
				if valid(ny, nx) && !visited[ny][nx] {
					// 비어있는 곳이거나 비버의 굴이라면 고슴도치가 이동할 수 있다.
					if forest[ny][nx] == '.' || forest[ny][nx] == 'D' {
						visited[ny][nx] = true
						candidates = append(candidates, move{y: ny, x: nx, cnt: h.cnt + 1})
					}
				}
			}
		}

		hedgehog = append(hedgehog, candidates...)
	}

	fmt.Fprintln(writer, "KAKTUS") // 고슴도치가 비버의 굴에 도착하지 못했다면 KAKTUS를 출력한다.
}

func valid(y, x int) bool {
	return y <= R && y >= 1 && x <= C && x >= 1
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}

func scanBytes() []byte {
	scanner.Scan()
	return scanner.Bytes()
}
