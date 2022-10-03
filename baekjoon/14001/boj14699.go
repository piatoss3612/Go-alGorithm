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
	height  [5001]int       // 쉼터의 높이
	conn    [5001][]int     // 쉼터 연결 정보
	dp      [5001][5001]int // 쉼터 i에서 쉼터 j로 이동할 수 있을 때 최대로 방문할 수 있는 쉼터의 개수
)

// 메모리: 304800KB
// 시간: 252ms
// 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	N, M = scanInt(), scanInt()

	// 높이 정보 입력
	for i := 1; i <= N; i++ {
		height[i] = scanInt()
	}

	// 연결 정보 입력
	for i := 1; i <= M; i++ {
		a, b := scanInt(), scanInt()
		conn[a] = append(conn[a], b)
		conn[b] = append(conn[b], a)
	}

	// 각 쉼터에서 최대로 방문할 수 있는 쉼터의 개수 출력
	for i := 1; i <= N; i++ {
		fmt.Fprintln(writer, solve(0, i))
	}
}

func solve(x, y int) int {
	ret := &dp[x][y]

	if *ret != 0 {
		return *ret
	}

	*ret++ // y번 쉼터 방문

	for _, v := range conn[y] {
		// y와 연결된 쉼터 v가 y보다 더 높은 곳에 있을 때
		if height[v] > height[y] {
			*ret = max(*ret, solve(y, v)+1) // y에서 v를 방문하는 경우와 최댓값 비교
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
