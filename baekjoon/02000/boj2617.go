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
	beads   [100][100]int
)

const INF = 987654321

// 난이도: Gold 4
// 메모리: 992KB
// 시간: 8ms
// 분류: 그래프 이론, 그래프 탐색, 플로이드-와샬
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
			beads[i][j] = INF
		}
	}

	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		beads[a][b] = 2 // a > b
		beads[b][a] = 1 // a < b
	}
}

func Solve() {
	for k := 1; k <= N; k++ {
		for i := 1; i <= N; i++ {
			for j := 1; j <= N; j++ {
				// i와 j가 비교가 안되는 경우
				if beads[i][j] == INF {
					// i와 k가 비교가 되고, k와 j가 비교가 되는 경우
					if beads[i][k] == 1 && beads[k][j] == 1 {
						beads[i][j] = 1 // i < j
					} else if beads[i][k] == 2 && beads[k][j] == 2 {
						beads[i][j] = 2 // i > j
					}
				}
			}
		}
	}

	answer := 0

	for i := 1; i <= N; i++ {
		small := 0 // i보다 작은 구슬의 개수
		big := 0  // i보다 큰 구슬의 개수

		for j := 1; j <= N; j++ {
			// i가 j보다 큰 경우
			if beads[i][j] == 2 {
				big++
			}

			// i가 j보다 작은 경우
			if beads[i][j] == 1 {
				small++
			}
		}

		/* 
			<i가 중간 구슬이 절대로 될 수 없는 경우>
		
			1. i보다 큰 구슬의 개수가 (N+1)/2보다 많은 경우
			2. i보다 작은 구슬의 개수가 (N+1)/2보다 많은 경우
		*/
		if small >= (N+1)/2 || big >= (N+1)/2 {
			answer++
		}
	}

	fmt.Fprintln(writer, answer)
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
