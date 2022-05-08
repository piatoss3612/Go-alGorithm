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
	n, m     int
	costs    []int
	orders   [][]int
	inDegree []int
	dp       []int
)

// 메모리: 6416KB
// 시간: 204ms

/*
시간 초과가 발생한 이유:

건설 시간이 [0, 0, 0, 0]처럼 0이 많은 경우
dp 값도 [0, 0, 0, 0]이 되고

func rec(x int) int {
	ret := &dp[x]
	if *ret != 0 { <- 여기서 탈출하지 못하고 재귀함수 호출을 지수시간 단위로 늘리기 때문
		return *ret
	}

	...

	return *ret
}

따라서 dp를 -1로 초기화 하고 재귀 함수 종료 조건을 -1이 아닌 경우로 수정했다
*/

// 나머지 풀이는 1561번 문제 풀이를 참고하면 된다
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	t := scanInt()
	for i := 1; i <= t; i++ {
		testCase()
	}
}

func testCase() {
	n, m = scanInt(), scanInt()

	costs = make([]int, n+1)
	for i := 1; i <= n; i++ {
		costs[i] = scanInt()
	}

	orders = make([][]int, n+1)
	inDegree = make([]int, n+1)
	var a, b int
	for i := 1; i <= m; i++ {
		a, b = scanInt(), scanInt()
		orders[b] = append(orders[b], a)
		inDegree[a] += 1
	}

	dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = -1
	}
	topologicalSort()

	fmt.Fprintln(writer, dp[scanInt()])
}

func topologicalSort() {
	var queue []int

	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]

		_ = rec(x)

		var y int
		for i := 0; i < len(orders[x]); i++ {
			y = orders[x][i]
			inDegree[y] -= 1
			if inDegree[y] == 0 {
				queue = append(queue, y)
			}
		}
	}
}

func rec(x int) int {
	ret := &dp[x]
	if *ret != -1 {
		return *ret
	}

	*ret = costs[x]

	for i := 0; i < len(orders[x]); i++ {
		*ret = max(*ret, rec(orders[x][i])+costs[x])
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
