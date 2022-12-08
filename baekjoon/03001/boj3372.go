package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N       int
	board   [101][101]int
	dp      [101][101]*big.Int // 경로의 개수가 2^63 - 1 보다 클 수 있으므로 go 언어의 큰 수(big int) 패키지를 사용
)

// 난이도: Silver 1
// 메모리: 8676KB
// 시간: 48ms
// 분류: 다이나믹 프로그래밍, 큰 수 연산
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Input()
	Solve()
}

func Input() {
	N = scanInt()
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			board[i][j] = scanInt()
			dp[i][j] = big.NewInt(0)
		}
	}
}

func Solve() {
	ans := rec(1, 1)
	fmt.Fprintln(writer, ans)
}

func rec(y, x int) *big.Int {
	// (N, N)에 도달할 수 있는경우
	if y == N && x == N {
		return big.NewInt(1) // 1을 반환
	}

	// *이부분 중요*
	// (y, x) 위치에서 더 이상 이동할 수 없는 경우
	if board[y][x] == 0 {
		return big.NewInt(0) // 0을 반환
	}

	ret := &dp[y][x]
	if (*ret).Cmp(big.NewInt(0)) != 0 {
		return *ret
	}

	move := board[y][x]

	// 아래쪽으로 이동할 수 있는 경우
	if y+move <= N {
		(*ret).Add(*ret, rec(y+move, x))
	}
	// 오른쪽으로 이동할 수 있는 경우
	if x+move <= N {
		(*ret).Add(*ret, rec(y, x+move))
	}
	return *ret
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
