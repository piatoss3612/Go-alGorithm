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
	N, K    int
	tube    [201][201]int
	dx      = []int{-1, +0, +1, +0}
	dy      = []int{+0, +1, +0, -1}
	queue   []Virus
	S, X, Y int
)

type Virus struct {
	number int
	x, y   int
}

// 난이도: Gold 5
// 메모리: 6208KB
// 시간: 20ms
// 분류: 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, K = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			n := scanInt()
			tube[i][j] = n
			if n != 0 {
				queue = append(queue, Virus{n, i, j})
			}
		}
	}
	// 번호가 빠른 바이러스가 먼저 증식되므로 번호순으로 오름차순 정렬
	sort.Slice(queue, func(i, j int) bool {
		return queue[i].number < queue[j].number
	})
	S, X, Y = scanInt(), scanInt(), scanInt()
}

func Solve() {
	// 매초마다 큐에 들어있는 바이러스가 번호 순서대로 증식되고
	// i초에 증식되어 이동한 바이러스의 정보는 next 슬라이스에 저장한다
	// 큐에 들어있던 바이러스가 모두 증식을 마치면
	// next 슬라이스에 들어있는 증식되어 이동한 바이러스들의 정보를 큐에 담고
	// i+1초에 해당하는 증식 과정을 진행한다
	for i := 1; i <= S; i++ {
		next := []Virus{}

		for len(queue) > 0 {
			front := queue[0]
			queue = queue[1:]

			for j := 0; j < 4; j++ {
				nx, ny := front.x+dx[j], front.y+dy[j]
				if valid(nx, ny) && tube[nx][ny] == 0 {
					tube[nx][ny] = front.number
					next = append(next, Virus{front.number, nx, ny})
				}
			}
		}

		queue = append(queue, next...)
	}

	fmt.Fprintln(writer, tube[X][Y])
}

func valid(x, y int) bool {
	return x >= 1 && x <= N && y >= 1 && y <= N
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
