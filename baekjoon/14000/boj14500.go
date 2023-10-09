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
	paper    [501][501]int
	visited  [501][501]int
	dy       = []int{-1, +0, +1, +0}
	dx       = []int{+0, +1, +0, -1}
	sum, ans int
)

// 난이도: Gold 4
// 메모리: 7268KB
// 시간: 168ms
// 분류: 브루트포스 알고리즘, 깊이 우선 탐색, 백트래킹

// 총 19개의 형태의 테트로미노를 각각 종이 위의 가능한 모든 위치에 배치해 봄으로써
// 테트로미노가 놓인 칸에 쓰인 수들의 합의 최댓값을 구한다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N, M = scanInt(), scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			paper[i][j] = scanInt()
		}
	}
}

func Solve() {
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			// (i, j)를 시작점으로 하는 모든 형태의 테트로미노를 만들어 본다
			// 탐색 중에 변경된 값들은 탐색이 종료될 때 모두 초기화된다
			visited[i][j] = 1
			sum += paper[i][j]
			DFS(i, j, 1)
			sum -= paper[i][j]
			visited[i][j] = 0
			sum = 0
		}
	}

	fmt.Fprintln(writer, ans)
}

func DFS(y, x, turn int) {
	// 4개의 폴리오미노를 연속으로 선택한 경우 == 테트로미노를 구성한 경우
	if turn == 4 {
		ans = max(ans, sum) // 최댓값을 갱신하고 탐색 종료
		return
	}

	// (y, x)에서 상하좌우로 연속된 칸에 폴리오미노를 배치할 수 있는지 탐색
	for i := 0; i < 4; i++ {
		ny, nx := y+dy[i], x+dx[i]
		// 연속된 칸에 폴리오미노를 배치할 수 있는 경우
		if valid(ny, nx) && visited[ny][nx] == 0 {
			/*
				ㅏ,ㅓ,ㅗ,ㅜ 형태의 테트로미노를 구성하기 위한 조건 처리

				이미 2개의 연속된 폴리오미노를 배치한 상황에서
				(#: 이전에 선택된 칸, @: 현재 위치)

				#@

				ny, nx 칸에 폴리오미노를 배치하고 깊이 우선 탐색을 실행하면

				 @
				##, ##@, ##
				          @

				ㅏ,ㅓ,ㅗ,ㅜ 형태의 테트로미노를 구성할 수 없게 된다

				그러므로 ny, nx 칸에 폴리오미노를 배치한 뒤에
				y, x 칸으로 돌아와서

				 #
				#@, #@#, #@
				          #

				깊이 우선 탐색을 실행을 하면

				 #   #    #
				#@, #@#, #@, #@#
				 #        #   #

				이와 같이 ㅏ,ㅓ,ㅗ,ㅜ 형태의 테트로미노를 구성할 수 있게 된다
			*/
			if turn == 2 {
				visited[ny][nx] = 1
				sum += paper[ny][nx]
				DFS(y, x, turn+1)
				sum -= paper[ny][nx]
				visited[ny][nx] = 0
			}

			// 나머지 15개의 테트로미노는 깊이 우선 탐색을 통해 자연적(?)으로 구성된다
			visited[ny][nx] = 1
			sum += paper[ny][nx]
			DFS(ny, nx, turn+1)
			sum -= paper[ny][nx]
			visited[ny][nx] = 0
		}
	}
}

func valid(y, x int) bool {
	if y >= 1 && y <= N && x >= 1 && x <= M {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
