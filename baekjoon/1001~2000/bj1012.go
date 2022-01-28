package bj1012

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	field   [][]int // 배추밭
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 0; i < t; i++ {
		testCase()
	}

}

func testCase() {
	m, n, k := scanInt(), scanInt(), scanInt()
	field = make([][]int, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, m)
	}
	for i := 0; i < k; i++ {
		x, y := scanInt(), scanInt()
		field[y][x] = 1
	}
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// field[i][j]에 배추가 있는 경우에
			// 상하좌우를 탐색하는 것을 반복하면서
			// 인접한 1들을 지워준다
			if field[i][j] == 1 {
				cnt += 1
				DFS(i, j)
			}
		}
	}
	fmt.Fprintln(writer, cnt)
}

// field[i][j]가 1인 경우, 0으로 바꿔주고 상하좌우 탐색
func DFS(i, j int) {
	field[i][j] = 0
	if i-1 >= 0 { // 위쪽 탐색 조건
		if field[i-1][j] == 1 {
			DFS(i-1, j)
		}
	}
	if j-1 >= 0 { // 왼쪽 탐색 조건
		if field[i][j-1] == 1 {
			DFS(i, j-1)
		}
	}
	if i+1 < len(field) { // 아래쪽 탐색 조건
		if field[i+1][j] == 1 {
			DFS(i+1, j)
		}
	}
	if j+1 < len(field[i]) { // 오른쪽 탐색 조건
		if field[i][j+1] == 1 {
			DFS(i, j+1)
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
