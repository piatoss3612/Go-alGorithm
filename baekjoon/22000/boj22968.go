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

	dp   [50]int // 깊이가 i인 AVL Tree를 구성하는 데 필요한 노드의 최소 개수
	T, V int
)

// 난이도: Gold 5
// 메모리: 932KB
// 시간: 4ms
// 분류: 다이나믹 프로그래밍
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	Setup()
	Solve()
}

func Setup() {
	dp[0] = 0
	dp[1] = 1
	for i := 2; i < 50; i++ {
		dp[i] = dp[i-1] + dp[i-2] + 1 // 왼쪽 서브트리 + 오른쪽 서브트리 + 루트 노드
	}
	T = scanInt()
}

func Solve() {
	for i := 1; i <= T; i++ {
		V = scanInt()
		for j := 1; j < 50; j++ {
			if dp[j] > V {
				fmt.Fprintln(writer, j-1)
				break
			}
		}
	}
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}
