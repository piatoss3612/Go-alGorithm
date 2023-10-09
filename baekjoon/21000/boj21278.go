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
	cities  [101][101]int
)

const INF = 987654321

// 난이도: Gold 5
// 메모리: 1020KB
// 시간: 16ms
// 분류: 플로이드-와샬, 브루트포스 알고리즘
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			cities[i][j] = INF
		}
	}

	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		// 양방향
		cities[a][b] = 1
		cities[b][a] = 1
	}
}

func Solve() {
	// 플로이드-와샬 알고리즘을 이용하여 모든 도시간의 최단거리를 구한다.
	for k := 1; k <= N; k++ {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				if i == j {
					continue
				}

				cities[i][j] = min(cities[i][j], cities[i][k]+cities[k][j])
			}
		}
	}

	minCity1, minCity2, minSum := 0, 0, INF

	// 브루트포스 알고리즘을 이용하여 다른 모든 도시로부터 도시 i, j로 가는 최단거리의 합이 최소가 되는 도시 i, j를 구한다.
	for i := 1; i < N; i++ {
		for j := i + 1; j <= N; j++ {
			sum := 0

			for k := 1; k <= N; k++ {
				if k == i || k == j {
					continue
				}
				sum += min(cities[i][k], cities[j][k])
			}

			if sum < minSum {
				minCity1, minCity2, minSum = i, j, sum
			}
		}
	}

	fmt.Fprintln(writer, minCity1, minCity2, minSum*2) // 왕복이므로 2를 곱한다.
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
