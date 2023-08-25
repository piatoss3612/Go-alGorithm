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

	N, M int
	path [501][501]int
)

// 난이도: Gold 5
// 메모리: 2668KB
// 시간: 80ms
// 분류: 플로이드 와샬
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
			path[i][j] = scanInt()
		}
	}

	// 모든 파티장이 직접적으로 연결되어 있고
	// i에서 j로 직접 연결되어 있는 통로보다 어딘가를 경유해서 가는게 더 빠를 수 있으므로
	// 플로이드 와샬 알고리즘을 사용하여 i에서 j로 가는 최단거리를 구한다.
	for k := 1; k <= N; k++ {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				if i == j {
					continue
				}
				path[i][j] = min(path[i][j], path[i][k]+path[k][j])
			}
		}
	}
}

func Solve() {
	for i := 1; i <= M; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		// a에서 b로 가는 최단거리가 c보다 작거나 같으면 시간 내에 갈 수 있다.
		// 그렇지 않으면 파티가 시작하기 전에 도착할 수 없다.
		if path[a][b] <= c {
			fmt.Fprintln(writer, "Enjoy other party")
		} else {
			fmt.Fprintln(writer, "Stay here")
		}
	}
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
