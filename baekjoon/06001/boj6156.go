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
	cows    [101][101]int
)

const INF = 987654321

// 난이도: Gold 4
// 메모리: 1016KB
// 시간: 8ms
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
			cows[i][j] = INF
		}
	}

	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		cows[a][b] = 1
	}
}

func Solve() {
	// 다른 소들과 연결되어 있는지 플로이드 와샬로 확인
	for k := 1; k <= N; k++ {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				if i == j {
					continue
				}

				cows[i][j] = min(cows[i][j], cows[i][k]+cows[k][j])
			}
		}
	}

	cnt := 0

	for i := 1; i <= N; i++ {
		flag := true

		for j := 1; j <= N; j++ {
			if i == j {
				continue
			}

			// i번째 소와 j번째 소가 서로 연결되어 있지 않다면 비교 불가
			if cows[i][j] == INF && cows[j][i] == INF {
				flag = false
				break
			}
		}

		// i번째 소와 다른 소들이 모두 비교 가능하다면 카운트를 증가
		if flag {
			cnt++
		}
	}

	fmt.Fprintln(writer, cnt)
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
