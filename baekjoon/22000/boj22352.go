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
	before  [31][31]int  // 백신을 투약하기 전의 상태
	after   [31][31]int  // 백신을 투약한 후의 상태
	visited [31][31]bool // 방문 여부

	dy = []int{-1, +0, +1, +0} // 상, 우, 하, 좌
	dx = []int{+0, +1, +0, -1} // 상, 우, 하, 좌
)

// 난이도: Gold 5
// 메모리: 916KB
// 시간: 4ms
// 시간복잡도: O(NM)
// 공간복잡도: O(NM)
// 분류: 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()

	// 백신을 투약하기 전의 상태
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			before[i][j] = scanInt()
		}
	}

	// 백신을 투약한 후의 상태
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			after[i][j] = scanInt()
		}
	}
}

func Solve() {
	// 백신을 투여하기 전의 상태와 백신을 투여한 후의 상태를 비교하고
	// 다른 부분을 찾아서 깊이 우선 탐색을 통해 백신을 투여한 후의 상태로 갱신한다.
	// 백신은 한 번만 투여하므로, 한 번의 깊이 우선 탐색이 끝나면 반복문을 종료한다.
Loop:
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if before[i][j] != after[i][j] && !visited[i][j] {
				dfs(i, j, before[i][j], after[i][j])
				break Loop
			}
		}
	}

	// 백신을 투여했으므로, before의 상태가 백신을 투여한 후의 상태 after와 같은지 확인한다.
	// 만약 다르다면, NO를 출력하고 종료한다.
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if before[i][j] != after[i][j] {
				fmt.Fprintln(writer, "NO")
				return
			}
		}
	}
	// 모든 부분이 같다면, CPCU-1202 백신을 투약한 것이므로 YES를 출력한다.
	fmt.Fprintln(writer, "YES")
}

func dfs(y, x, a, b int) {
	visited[y][x] = true // 방문 처리
	before[y][x] = b     // 백신을 투여한 후의 상태로 갱신

	for i := 0; i < 4; i++ {
		ny, nx := y+dy[i], x+dx[i]
		// 범위 내에 있고, 방문하지 않은 경우
		if valid(ny, nx) && !visited[ny][nx] {
			// 상태값이 a인 경우에만
			if before[ny][nx] == a {
				dfs(ny, nx, a, b)
			}
		}
	}
}

func valid(y, x int) bool {
	return y >= 1 && y <= N && x >= 1 && x <= M
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func scanInt() int {
	scanner.Scan()
	return parseInt(scanner.Text())
}
