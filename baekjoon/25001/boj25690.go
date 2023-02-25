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
	N        int
	children [100000][]int
	dp       [100000][2]int
)

// 난이도: Gold 3
// 메모리: 23436KB
// 시간: 96ms
// 분류: 다이나믹 프로그래밍, 트리, 깊이 우선 탐색
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)
	Setup()
	Solve()
}

func Setup() {
	N = scanInt()
	for i := 0; i < N-1; i++ {
		p, c := scanInt(), scanInt()
		children[p] = append(children[p], c)
	}
	for i := 0; i < N; i++ {
		dp[i][0], dp[i][1] = scanInt(), scanInt()
	}
}

func Solve() {
	fmt.Fprintln(writer, min(DFS(0)))
}

func DFS(x int) (int, int) {
	for _, child := range children[x] {
		w, b := DFS(child)
		dp[x][0] += min(w, b)
		dp[x][1] += w
	}
	return dp[x][0], dp[x][1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
