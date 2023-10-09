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
	N, T, M int
	cities  [1001]City
	dist    [1001][1001]int
)

type City struct {
	s, x, y int
}

const INF = 987654321

// 난이도: Gold 4
// 메모리: 8792KB
// 시간: 1856ms
// 분류: 플로이드 와샬
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, T = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			dist[i][j] = INF
		}
	}

	for i := 1; i <= N; i++ {
		s, x, y := scanInt(), scanInt(), scanInt()
		cities[i] = City{s, x, y}
	}
}

func Solve() {
	// 모든 도시에서 모든 도시로의 거리를 구한다.
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if i == j {
				continue
			}

			temp := abs(cities[i].x-cities[j].x) + abs(cities[i].y-cities[j].y)

			if cities[i].s == 1 && cities[j].s == 1 {
				temp = min(temp, T)
			}

			dist[i][j] = temp
		}
	}

	// 플로이드 와샬 알고리즘을 이용하여 모든 도시에서 모든 도시로의 최단 거리를 구한다.
	for k := 1; k <= N; k++ {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				if i == j {
					continue
				}

				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	M = scanInt()

	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		fmt.Fprintln(writer, dist[a][b])
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
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
