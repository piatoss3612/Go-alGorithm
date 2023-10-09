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

	N, M     int
	A, B     int
	relation [101][101]int
	visited  [101]bool
	depth    [101]int
)

// 난이도: Silver 2
// 메모리: 968KB
// 시간: 8ms
// 분류: 그래프 이론, 그래프 탐색, 너비 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	A, B = scanInt(), scanInt()
	M = scanInt()

	for i := 1; i <= M; i++ {
		x, y := scanInt(), scanInt()
		relation[x][y] = 1
		relation[y][x] = 1
	}
}

func Solve() {
	q := []int{A}
	visited[A] = true

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur == B {
			break
		}

		for i := 1; i <= N; i++ {
			if relation[cur][i] == 1 && !visited[i] {
				visited[i] = true
				q = append(q, i)
				depth[i] = depth[cur] + 1
			}
		}
	}

	if !visited[B] {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, depth[B])
	}
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
