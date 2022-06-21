package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	scanner  = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	n        int
	graph    [][]int
	taken    []int
	inDegree []int
)

// 메모리: 11452KB
// 시간: 84ms
func main() {
	defer writer.Flush()
	scanner.Split(bufio.ScanWords)

	n = scanInt()
	graph = make([][]int, n+1)
	taken = make([]int, n+1)
	inDegree = make([]int, n+1)

	for i := 1; i <= n; i++ {
		taken[i] = scanInt()
		count := scanInt()

		for j := 1; j <= count; j++ {
			pre := scanInt()
			graph[pre] = append(graph[pre], i)
			inDegree[i] += 1
		}
	}

	TopologicalSort()
}

func TopologicalSort() {
	dp := make([]int, n+1)
	q := make([]int, 0, n+1)

	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			dp[i] = taken[i]
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]

		for i := 0; i < len(graph[x]); i++ {
			y := graph[x][i]
			dp[y] = max(dp[y], dp[x]+taken[y])
			inDegree[y] -= 1
			if inDegree[y] == 0 {
				q = append(q, y)
			}
		}
	}
	sort.Ints(dp)
	fmt.Fprintln(writer, dp[n])
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
