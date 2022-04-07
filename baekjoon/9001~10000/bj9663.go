package bj9663

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	n       int
	board   []int
	ans     = 0
)

func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	n = scanInt()

	/*
		입력: 4

		백트래킹 프로세스:
		[0 0 0 0 0]
		[0 1 0 0 0]
		[0 1 3 0 0]
		[0 1 4 4 0]
		[0 1 4 2 0]
		[0 2 4 4 4]
		[0 2 4 4 4]
		[0 2 4 1 4]
		[0 2 4 1 3]
		[0 3 4 4 4]
		[0 3 1 4 4]
		[0 3 1 4 4]
		[0 3 1 4 2]
		[0 4 4 4 4]
		[0 4 1 4 4]
		[0 4 1 3 4]
		[0 4 2 4 4]

		출력: 2
	*/

	board = make([]int, n+1)

	nQueen(1)
	fmt.Fprintln(writer, ans)
}

func canPlace(x int) bool {
	// x번째 행 이전의 행들에 놓인 퀸과 같은 열 또는 대각선에 위치하고 있는지 검사
	for i := 1; i < x; i++ {
		// 같은 열인지 판별 || 대각선 판별
		if board[x] == board[i] || x-i == abs(board[x]-board[i]) {
			return false
		}
	}
	return true
}

func nQueen(x int) {
	if x == n+1 {
		ans += 1
		return
	}

	for i := 1; i <= n; i++ {
		board[x] = i
		// x행의 i열에 퀸을 놓을 수 있는지 판별하고
		// 가능한 경우: 다음 행으로 넘어간다
		// 불가능한 경우: 퀸을 다음 열로 움직인다
		if canPlace(x) {
			nQueen(x + 1)
		}
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
