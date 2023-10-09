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
	tree    [5]int           // tree[i]: 옷걸이 트리의 깊이(높이)가 i일 때 필요한 옷걸이의 개수, 높이 5이상은 불가능
	dp      [5001][10001]int // dp[i][j]: i개의 위치에 j개의 옷걸이를 사용하여 걸 수 있는 옷의 최대 개수
)

const INF = -987654321

// 난이도: Gold 3
// 메모리: 574984KB
// 시간: 724ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N, M = scanInt(), scanInt()

	// 옷걸이 트리의 높이가 i일 때 필요한 옷걸이의 개수 구하기
	for i := 1; i <= 4; i++ {
		tree[i] = tree[i-1] + 1<<(i-1)
	}
}

func Solve() {
	ans := rec(N, M)
	if ans < 0 {
		fmt.Fprintln(writer, -1) // 옷걸이를 모두 사용할 수 없는 경우
	} else {
		fmt.Fprintln(writer, ans) // 옷걸이를 모두 사용한 경우
	}
}

func rec(space, hanger int) int {
	// 기저 사례1: 옷걸이를 걸 수 있는 위치가 없는데 옷걸이가 남은 경우
	if space == 0 && hanger != 0 {
		return INF
	}

	// 기저 사례2: 옷걸이를 모두 사용한 경우
	if hanger == 0 {
		return 0
	}

	ret := &dp[space][hanger]
	if *ret != 0 {
		return *ret
	}

	*ret = INF // 최대값을 구하기 위해 초기값을 아주 작은 음수로 설정

	for i := 1; i <= 4; i++ {
		// 높이가 i인 옷걸이 트리를 걸 수 있는 경우
		if hanger-tree[i] >= 0 {
			// 높이가 4인 옷걸이는 만들 수는 있지만 옷장의 높이때문에 옷을 걸 수는 없다
			if i == 4 {
				*ret = max(*ret, rec(space-1, hanger-tree[i]))
			} else {
				*ret = max(*ret, rec(space-1, hanger-tree[i])+1<<(i-1))
			}
		}
	}

	return *ret
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
